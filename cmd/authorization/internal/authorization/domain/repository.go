package repository

import (
	"context"
)

type AuthorizationRepo interface {
	CreateSession(ctx context.Context, sessionId string, userId string) (string, error)
	DeleteSession(ctx context.Context, sessionId string) (string, error)
	GetSession(ctx context.Context, sessionId string) (string, error)
}
