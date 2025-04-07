package repository

import (
	"context"
	"fmt"
	repository "github.com/Errera11/authorization/internal/authorization/domain"
	"time"

	"github.com/redis/go-redis/v9"
)

type AuthorizationRepo struct {
	conn *redis.Client
}

const sessionExp = 24 * time.Hour

func NewAuthorizationRepo(conn *redis.Client) repository.AuthorizationRepo {
	return &AuthorizationRepo{
		conn: conn,
	}
}

func (a AuthorizationRepo) CreateSession(ctx context.Context, sessionId string, userId string) (string, error) {
	err := a.conn.Set(ctx, sessionId, userId, sessionExp).Err()
	if err != nil {
		return "", fmt.Errorf("Unable to set key: %v", sessionId)
	}

	return sessionId, nil
}

func (a AuthorizationRepo) DeleteSession(ctx context.Context, sessionId string) (string, error) {
	err := a.conn.Del(ctx, sessionId).Err()

	if err != nil {
		return "", fmt.Errorf("Unable to delete key: %v", sessionId)
	}

	return sessionId, nil
}

func (a AuthorizationRepo) GetSession(ctx context.Context, sessionId string) (string, error) {
	userId, err := a.conn.Get(ctx, sessionId).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("Session with key %v not found", sessionId)
	} else if err != nil {
		return "", fmt.Errorf("Error happened retrieving key: %v ", sessionId)
	} else {
		return userId, nil
	}
}
