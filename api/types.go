package api

import (
	"errors"
	"time"

	"github.com/graphql-go/graphql"
	impact "github.com/impactasaurus/server"
	"github.com/impactasaurus/server/auth"
	"github.com/impactasaurus/server/logic"
)

func (v *v1) initSchemaTypes() {
	v.organisationType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Organisation",
		Description: "An organisation",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.ID),
				Description: "Unique identifier for the organisation",
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Organisation's name",
			},
		},
	})

	v.questionInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "QuestionInterface",
		Description: "The interface satisfied by all question types",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.ID),
				Description: "Unique ID for the question",
			},
			"question": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The question",
			},
			"archived": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "Whether the question has been archived",
			},
			"categoryID": &graphql.Field{
				Type:        graphql.String,
				Description: "The category the question belongs to",
			},
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			obj, ok := p.Value.(impact.Question)
			if !ok {
				return v.likertScale
			}
			switch obj.Type {
			case impact.LIKERT:
				return v.likertScale
			default:
				return v.likertScale
			}
		},
	})

	v.likertScale = graphql.NewObject(graphql.ObjectConfig{
		Name:        "LikertScale",
		Description: "Question gathering information using Likert Scales",
		Interfaces: []*graphql.Interface{
			v.questionInterface,
		},
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.ID),
				Description: "Unique ID for the question",
			},
			"question": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The question",
			},
			"archived": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "Whether the question has been archived",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					obj, ok := p.Source.(impact.Question)
					if !ok {
						return nil, errors.New("Expecting an impact.Question")
					}
					return obj.Deleted, nil
				},
			},
			"categoryID": &graphql.Field{
				Type:        graphql.String,
				Description: "The category the question belongs to",
			},
			"minValue": &graphql.Field{
				Type:        graphql.Int,
				Description: "The minimum value in the scale",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					obj, ok := p.Source.(impact.Question)
					if !ok {
						return nil, errors.New("Expecting an impact.Question")
					}
					minValue, ok := obj.Options["minValue"]
					if !ok {
						return nil, nil
					}
					minValueInt, ok := minValue.(int)
					if !ok {
						return nil, errors.New("Min likert value should be an int")
					}
					return minValueInt, nil
				},
			},
			"maxValue": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The maximum value in the scale",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					obj, ok := p.Source.(impact.Question)
					if !ok {
						return nil, errors.New("Expecting an impact.Question")
					}
					maxValue, ok := obj.Options["maxValue"]
					if !ok {
						return nil, nil
					}
					maxValueInt, ok := maxValue.(int)
					if !ok {
						return nil, errors.New("Max likert value should be an int")
					}
					return maxValueInt, nil
				},
			},
			"minLabel": &graphql.Field{
				Type:        graphql.String,
				Description: "The string labelling the minimum value in the scale",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					obj, ok := p.Source.(impact.Question)
					if !ok {
						return nil, errors.New("Expecting an impact.Question")
					}
					label, ok := obj.Options["minLabel"]
					if !ok {
						return nil, nil
					}
					labelStr, ok := label.(string)
					if !ok {
						return nil, errors.New("Min likert label should be an string")
					}
					return labelStr, nil
				},
			},
			"maxLabel": &graphql.Field{
				Type:        graphql.String,
				Description: "The string labelling the maximum value in the scale",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					obj, ok := p.Source.(impact.Question)
					if !ok {
						return nil, errors.New("Expecting an impact.Question")
					}
					label, ok := obj.Options["maxLabel"]
					if !ok {
						return nil, nil
					}
					labelStr, ok := label.(string)
					if !ok {
						return nil, errors.New("Max likert label should be an string")
					}
					return labelStr, nil
				},
			},
		},
	})

	v.aggregationEnum = graphql.NewEnum(graphql.EnumConfig{
		Name:        "Aggregation",
		Description: "Aggregation functions available",
		Values: graphql.EnumValueConfigMap{
			string(impact.MEAN): &graphql.EnumValueConfig{
				Value:       impact.MEAN,
				Description: "Mean",
			},
			string(impact.SUM): &graphql.EnumValueConfig{
				Value:       impact.SUM,
				Description: "Sum",
			},
		},
	})

	v.categoryType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Category",
		Description: "Categorises a set of questions. Used for aggregation",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.ID),
				Description: "Unique ID",
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Name of the category",
			},
			"description": &graphql.Field{
				Type:        graphql.String,
				Description: "Description of the category",
			},
			"aggregation": &graphql.Field{
				Type:        graphql.NewNonNull(v.aggregationEnum),
				Description: "The aggregation applied to the category",
			},
		},
	})

	v.outcomeSetType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "OutcomeSet",
		Description: "A set of questions to determine outcomes",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.ID),
				Description: "Unique ID",
			},
			"organisationID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Organisation's unique ID",
			},
			"organisation": &graphql.Field{
				Type:        graphql.NewNonNull(v.organisationType),
				Description: "The owning organisation of the outcome set",
				Resolve: userRestrictedResolver(func(p graphql.ResolveParams, u auth.User) (interface{}, error) {
					obj, ok := p.Source.(impact.OutcomeSet)
					if !ok {
						return nil, errors.New("Expecting an impact.Meeting")
					}
					return v.db.GetOrganisation(obj.OrganisationID, u)
				}),
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Name of the outcome set",
			},
			"description": &graphql.Field{
				Type:        graphql.String,
				Description: "Information about the outcome set",
			},
			"questions": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.NewList(v.questionInterface)),
				Description: "Questions associated with the outcome set",
			},
			"categories": &graphql.Field{
				Type:        graphql.NewList(v.categoryType),
				Description: "Questions associated with the outcome set",
			},
		},
	})

	v.answerInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "AnswerInterface",
		Description: "The interface satisfied by all answer types",
		Fields: graphql.Fields{
			"questionID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The ID of the question answered",
			},
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			obj, ok := p.Value.(impact.Answer)
			if !ok {
				return v.intAnswer
			}
			switch obj.Type {
			case impact.INT:
				return v.intAnswer
			default:
				return v.intAnswer
			}
		},
	})

	v.intAnswer = graphql.NewObject(graphql.ObjectConfig{
		Name:        "IntAnswer",
		Description: "Answer containing an integer value",
		Interfaces: []*graphql.Interface{
			v.answerInterface,
		},
		Fields: graphql.Fields{
			"questionID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The ID of the question answered",
			},
			"answer": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The provided int answer",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					obj, ok := p.Source.(impact.Answer)
					if !ok {
						return nil, errors.New("Expecting an impact.Answer")
					}
					num, ok := obj.Answer.(int)
					if !ok {
						return nil, errors.New("Expected an int value")
					}
					return num, nil
				},
			},
		},
	})

	v.categoryAggregate = graphql.NewObject(graphql.ObjectConfig{
		Name:        "CategoryAggregate",
		Description: "An aggregation of answers to the category level",
		Fields: graphql.Fields{
			"categoryID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The ID of the category being aggregated",
			},
			"aggregate": &graphql.Field{
				Type:        graphql.Float,
				Description: "The aggregated value",
			},
		},
	})

	v.aggregates = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Aggregates",
		Description: "Aggregations of the meeting",
		Fields: graphql.Fields{
			"category": &graphql.Field{
				Type:        graphql.NewList(v.categoryAggregate),
				Description: "Answers aggregated to the category level",
			},
		},
	})

	v.meetingType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Meeting",
		Description: "A set of answers for an outcome set",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.ID),
				Description: "Unique ID for the meeting",
			},
			"beneficiary": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The beneficiary providing the answers",
			},
			"user": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The user who collected the answers",
			},
			"outcomeSetID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The ID of the outcome set answered",
			},
			"outcomeSet": &graphql.Field{
				Type:        graphql.NewNonNull(v.outcomeSetType),
				Description: "The outcome set answered",
				Resolve: userRestrictedResolver(func(p graphql.ResolveParams, u auth.User) (interface{}, error) {
					obj, ok := p.Source.(impact.Meeting)
					if !ok {
						return nil, errors.New("Expecting an impact.Meeting")
					}
					return v.db.GetOutcomeSet(obj.OutcomeSetID, u)
				}),
			},
			"organisationID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Organisation's unique ID",
			},
			"organisation": &graphql.Field{
				Type:        graphql.NewNonNull(v.organisationType),
				Description: "The owning organisation of the outcome set",
				Resolve: userRestrictedResolver(func(p graphql.ResolveParams, u auth.User) (interface{}, error) {
					obj, ok := p.Source.(impact.Meeting)
					if !ok {
						return nil, errors.New("Expecting an impact.Meeting")
					}
					return v.db.GetOrganisation(obj.OrganisationID, u)
				}),
			},
			"answers": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.NewList(v.answerInterface)),
				Description: "The answers provided in the meeting",
			},
			"aggregates": &graphql.Field{
				Type:        v.aggregates,
				Description: "Aggregations of the meeting's answers",
				Resolve: userRestrictedResolver(func(p graphql.ResolveParams, u auth.User) (interface{}, error) {
					obj, ok := p.Source.(impact.Meeting)
					if !ok {
						return nil, errors.New("Expecting an impact.Meeting")
					}
					os, err := v.db.GetOutcomeSet(obj.OutcomeSetID, u)
					if err != nil {
						return nil, err
					}
					catAgs, err := logic.GetCategoryAggregates(obj, os)
					if err != nil {
						return nil, err
					}
					return impact.Aggregates{
						Category: catAgs,
					}, nil
				}),
			},
			"conducted": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "When the meeting was conducted",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					obj, ok := p.Source.(impact.Meeting)
					if !ok {
						return nil, errors.New("Expecting an impact.Meeting")
					}
					return obj.Conducted.Format(time.RFC3339), nil
				},
			},
			"created": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "When the meeting was entered into the system",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					obj, ok := p.Source.(impact.Meeting)
					if !ok {
						return nil, errors.New("Expecting an impact.Meeting")
					}
					return obj.Created.Format(time.RFC3339), nil
				},
			},
			"modified": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "When the meeting was last modified in the system",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					obj, ok := p.Source.(impact.Meeting)
					if !ok {
						return nil, errors.New("Expecting an impact.Meeting")
					}
					return obj.Modified.Format(time.RFC3339), nil
				},
			},
		},
	})
}
