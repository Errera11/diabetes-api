package handler

import (
	"context"
	"fmt"
	"github.com/Errera11/authorization/internal/authorization/service"
	authorization "github.com/Errera11/authorization/internal/protogen/authorization"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
)

type AuthorizationGrpcHandler struct {
	authorizationService service.AuthorizationService
	validator            *validator.Validate
	authorization.UnimplementedAuthServiceServer
}

func (a AuthorizationGrpcHandler) Signin(ctx context.Context, request *authorization.SigninRequest) (*authorization.SigninResponse, error) {
	parsedReq := &SigninValidator{
		Email:    request.Email,
		Password: request.Password,
	}

	err := a.validator.Struct(parsedReq)

	if err != nil {
		return nil, err
	}

	creds, err := a.authorizationService.SignIn(ctx, request)
	if err != nil {
		fmt.Println(err)
	}

	return creds, err
}

func (a AuthorizationGrpcHandler) Signup(ctx context.Context, request *authorization.SignupRequest) (*authorization.SignupResponse, error) {
	parsedReq := &SignupValidator{
		Email:    request.Email,
		Password: request.Password,
		Username: request.Username,
	}
	err := a.validator.Struct(parsedReq)
	if err != nil {
		return nil, err
	}

	return a.authorizationService.SignUp(ctx, request)
}

func (a AuthorizationGrpcHandler) Logout(ctx context.Context, request *authorization.LogoutRequest) (*authorization.LogoutResponse, error) {
	parsedReq := &LogoutValidator{
		Token: request.Token,
	}
	err := a.validator.Struct(parsedReq)
	if err != nil {
		return nil, err
	}

	return a.authorizationService.Logout(ctx, request)
}

func NewGrpcAuthorizationService(grpc *grpc.Server, authorizationService service.AuthorizationService) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	gRPCHandler := &AuthorizationGrpcHandler{
		authorizationService: authorizationService,
		validator:            validate,
	}
	authorization.RegisterAuthServiceServer(grpc, gRPCHandler)
}
