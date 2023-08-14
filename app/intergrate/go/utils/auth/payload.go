package authprovider

import (
	"errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	SessionId string    `json:"sid"`
	Issuer    string    `json:"iss"`
	UserName  string    `json:"un"`
	IssuedAt  time.Time `json:"iat"`
	ExpiredAt time.Time `json:"exp"`
	ClientID  string    `json:"aud"` //Audience
	UserID    string    `json:"sub"`
	ShardId   int32     `json:"sh"`
	Status    int32     `json:"us"`
	IsAdmin   bool      `json:"ad"`
}

func NewPayload(issuer string, clientId string, userID string, userName string, duration time.Duration, sessionID string, shardId int32) (*Payload, error) {

	if issuer == "" {
		issuer = "id.mimiland.com"
	}
	payload := &Payload{
		Issuer:    issuer,
		SessionId: sessionID,
		ClientID:  clientId,
		UserID:    userID,
		UserName:  userName,
		ShardId:   shardId,
		IssuedAt:  time.Now().UTC(),
		ExpiredAt: time.Now().UTC().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
