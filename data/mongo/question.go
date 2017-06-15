package mongo

import (
	"github.com/impactasaurus/server/auth"
	"github.com/impactasaurus/server/data"
	impact "github.com/impactasaurus/server"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

func (m *mongo) GetQuestion(outcomeSetID string, questionID string, u auth.User) (*impact.Question, error) {
	os, err := m.GetOutcomeSet(outcomeSetID, u)
	if err != nil {
		return nil, err
	}

	for _, q := range os.Questions {
		if q.ID == questionID {
			return &q, nil
		}
	}
	return nil, data.NewNotFoundError("Question")
}

func (m *mongo) NewQuestion(outcomeSetID, question, questionType string, options map[string]interface{}, u auth.User) (*impact.Question, error) {
	userOrg, err := u.Organisation()
	if err != nil {
		return nil, err
	}

	col, closer := m.getOutcomeCollection()
	defer closer()

	id := uuid.NewV4()

	newQuestion := &impact.Question{
		ID: id.String(),
		Question: question,
		Type: questionType,
		Options: options,
		Deleted: false,
	}

	if err := col.Update(bson.M{
		"_id": outcomeSetID,
		"organisationID": userOrg,
	}, bson.M{
		"$push": bson.M{
			"questions": newQuestion,
		},
	}); err != nil {
		return nil, err
	}

	return m.GetQuestion(outcomeSetID, id.String(), u)
}

func (m *mongo) DeleteQuestion(outcomeSetID, questionID string, u auth.User) error {
	userOrg, err := u.Organisation()
	if err != nil {
		return err
	}

	col, closer := m.getOutcomeCollection()
	defer closer()

	return col.Update(bson.M{
		"_id": outcomeSetID,
		"organisationID": userOrg,
		"questions.id": questionID,
	}, bson.M{
		"$set": bson.M{
			"questions.$.deleted": true,
		},
	})
}

func (m *mongo) EditQuestion(outcomeSetID, questionID, question, questionType string, options map[string]interface{}, u auth.User) (*impact.Question, error) {
	userOrg, err := u.Organisation()
	if err != nil {
		return nil, err
	}

	col, closer := m.getOutcomeCollection()
	defer closer()

	newQ := &impact.Question{
		ID: questionID,
		Question: question,
		Type: questionType,
		Options: options,
		Deleted: false,
	}

	if err := col.Update(bson.M{
		"_id": outcomeSetID,
		"organisationID": userOrg,
		"questions.id": questionID,
	}, bson.M{
		"$set": bson.M{
			"questions.$": newQ,
		},
	}); err != nil {
		return nil, err
	}
	return newQ, nil
}
