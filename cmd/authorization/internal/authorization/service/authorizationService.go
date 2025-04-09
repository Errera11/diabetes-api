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
	userService       userProto.UserServiceClient
	authorizationProto.UnimplementedAuthServiceServer
}

func New(authorizationRepo repository.AuthorizationRepo, userService userProto.UserServiceClient) *AuthorizationService {
	return &AuthorizationService{authorizationRepo: authorizationRepo, userService: userService}
}

func (s *AuthorizationService) SignIn(ctx context.Context, payload *authorizationProto.SigninRequest) (*authorizationProto.SigninResponse, error) {
	dbUser, err := s.userService.GetUserByEmail(ctx, &userProto.GetUserByEmailRequest{
		Email: payload.Email,
	})

	if err != nil {
		return &authorizationProto.SigninResponse{}, fmt.Errorf("Error signing in user with email %v: %e", payload.Email, err)
	}

	if dbUser == nil {
		return &authorizationProto.SigninResponse{}, fmt.Errorf("User with email %e was not found", payload.Email)
	}

	sessionId := uuid.New().String()

	token, err := s.authorizationRepo.CreateSession(ctx, sessionId, dbUser.Id)

	return &authorizationProto.SigninResponse{
		Token: token,
	}, err
}
func (s *AuthorizationService) SignUp(ctx context.Context, payload *authorizationProto.SignupRequest) (*authorizationProto.SignupResponse, error) {
	dbUser, err := s.userService.GetUserByEmail(ctx, &userProto.GetUserByEmailRequest{
		Email: payload.Email,
	})

	if dbUser != nil {
		return &authorizationProto.SignupResponse{}, fmt.Errorf("user with email %v already exist", payload.Email)
	}

	newUser, err := s.userService.CreateUser(ctx, &userProto.CreateUserRequest{
		Email:    payload.Email,
		Password: payload.Password,
		Username: payload.Username,
		Image:    payload.Image,
	})

	sessionId := uuid.New().String()

	token, err := s.authorizationRepo.CreateSession(ctx, sessionId, newUser.UserId)

	return &authorizationProto.SignupResponse{
		UserId: newUser.UserId,
		Token:  token,
	}, err
}
func (s *AuthorizationService) Logout(ctx context.Context, payload *authorizationProto.LogoutRequest) (*authorizationProto.LogoutResponse, error) {
	_, err := s.authorizationRepo.DeleteSession(ctx, payload.Token)

	if err != nil {
		return &authorizationProto.LogoutResponse{
			Message: "Failed to delete session",
		}, err
	}

	return &authorizationProto.LogoutResponse{
		Message: "Session deleted succesfully",
	}, nil
}
