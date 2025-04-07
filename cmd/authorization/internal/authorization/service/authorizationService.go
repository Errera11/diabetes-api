package service

import (
	"context"
	"fmt"
	repository "github.com/Errera11/authorization/internal/authorization/domain"
	authorizationProto "github.com/Errera11/authorization/internal/protogen/authorization"
	userProto "github.com/Errera11/authorization/internal/protogen/user"
	"github.com/google/uuid"
)

type AuthorizationService struct {
	authorizationRepo repository.AuthorizationRepo
	authorizationProto.UnimplementedAuthServiceServer
}

func New(authorizationRepo repository.AuthorizationRepo) *AuthorizationService {
	return &AuthorizationService{authorizationRepo: authorizationRepo}
}

func (s *AuthorizationService) SignIn(ctx context.Context, payload authorizationProto.SigninRequest) (authorizationProto.SigninResponse, error) {
	// TODO implement kafka
	user := kafka.userService.get(payload.Email)

	if user == nil {
		return authorizationProto.SigninResponse{}, fmt.Errorf("user with email %v was not found", payload.Email)
	}

	sessionId := uuid.New().String()

	token, err := s.authorizationRepo.CreateSession(ctx, sessionId, user.id)

	return authorizationProto.SigninResponse{
		Token: token,
	}, err
}
func (s *AuthorizationService) SignUp(ctx context.Context, user userProto.CreateUserRequest) (authorizationProto.SignupResponse, error) {
	// TODO implement kafka
	user := kafka.userService.get(payload.Email)

	if user == nil {
		return authorizationProto.SignupResponse{}, fmt.Errorf("user with email %v already exist", payload.Email)
	}

	newUser = kafka.userService.create(payload.Email)

	sessionId := uuid.New().String()

	token, err := s.authorizationRepo.CreateSession(ctx, sessionId, user.id)

	return authorizationProto.SignupResponse{
		userId: newUser.id,
		Token:  token,
	}, err
}
func (s *AuthorizationService) Logout(ctx context.Context, payload authorizationProto.LogoutRequest) (authorizationProto.LogoutResponse, error) {
	err := s.authorizationRepo.DeleteSession(ctx, payload.Token)

	if err != nil {
		return authorizationProto.LogoutResponse{
			Message: "Failed to delete session",
		}, err
	}

	return authorizationProto.LogoutResponse{
		Message: "Session deleted succesfully",
	}, nil
}
