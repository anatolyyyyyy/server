package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"

	jwtLib "github.com/dgrijalva/jwt-go"
)

type jwt struct {
	aud string
	iss string
	key *rsa.PublicKey
}

type jwtUser struct {
	AppMetadata  map[string]interface{} `json:"app_metadata"`
	UserMetadata map[string]interface{} `json:"user_metadata"`
	jwtLib.StandardClaims
}

// NewJWTAuthenticator returns an Authenticator which supports JWTs
func NewJWTAuthenticator(aud, iss string, key *rsa.PublicKey) Authenticator {
	return &jwt{
		aud: aud,
		iss: iss,
		key: key,
	}
}

func (j *jwt) AuthUser(jwt string) (User, error) {
	token, err := jwtLib.ParseWithClaims(jwt, &jwtUser{}, func(token *jwtLib.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtLib.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return j.key, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtUser)
	if !ok {
		return nil, errors.New("Failed to parse token to user object")
	}
	if err = claims.Valid(); err != nil {
		return nil, err
	}
	if !claims.VerifyAudience(j.aud, true) {
		return nil, errors.New("Token validation error: aud")
	}
	if !claims.VerifyIssuer(j.iss, true) {
		return nil, errors.New("Token validation error: iss")
	}
	return claims, nil
}

func (j *jwtUser) Organisation() (string, error) {
	org, ok := j.AppMetadata["organisation"].(string)
	if !ok {
		return "", errors.New("Failed to extract organisation")
	}
	return org, nil
}

func (j *jwtUser) UserID() string {
	return j.Subject
}

func (j *jwtUser) IsBeneficiary() bool {
	ben, ok := j.AppMetadata[beneficiaryKey].(bool)
	return ok && ben
}
func (j *jwtUser) GetAssessmentScope() (string, bool) {
	assessmentID, ok := j.AppMetadata[assessmentScopeKey].(string)
	return assessmentID, ok
}
