package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("invalid token")
)

type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	IssueAt  time.Time `json:"issued_at"`
	ExpireAt time.Time `json:"expired_at"`
}

func NewPayload(userName string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:       tokenID,
		Username: userName,
		IssueAt:  time.Now(),
		ExpireAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload Payload) Valid() error {
	if time.Now().After(payload.ExpireAt) {
		return ErrExpiredToken
	}
	return nil
}
