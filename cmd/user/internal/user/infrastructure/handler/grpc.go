package handler

import (
	"context"
	"fmt"
	userProto "github.com/Errera11/user/internal/protogen"
	"github.com/Errera11/user/internal/user/service"
	"github.com/go-playground/validator/v10"

	"google.golang.org/grpc"
)

type UserGrpcHandler struct {
	userService service.UserService
	userProto.UnimplementedUserServiceServer
}

var validate *validator.Validate

func (h *UserGrpcHandler) GetUserById(ctx context.Context, request *userProto.GetUserByIdRequset) (*userProto.GetUserByIdResponse, error) {
	validate = validator.New(validator.WithRequiredStructEnabled())

	incomingData := &UserRequestValidator{UserId: request.UserId}
	err := validate.Struct(incomingData)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	userRecord := h.userService.GetUserById(ctx, request.GetUserId())
	res := userProto.GetUserByIdResponse{
		Id:        userRecord.Id,
		Username:  userRecord.Username,
		CreatedAt: userRecord.CreatedAt,
		Email:     userRecord.Email,
		Image:     userRecord.Image,
	}

	return &res, nil
}

func (h *UserGrpcHandler) CreateUser(ctx context.Context, request *userProto.CreateUserRequest) (*userProto.CreateUserResponse, error) {
	validate = validator.New(validator.WithRequiredStructEnabled())

	incomingData := &CreateUserValidator{Username: request.Username, Password: request.Password, Email: request.Email}
	err := validate.Struct(incomingData)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	userRecord := h.userService.CreateUser(ctx, request)

	return userRecord, nil
}

func (h *UserGrpcHandler) GetAllUsers(ctx context.Context, request *userProto.Pagination) (*userProto.GetAllUsersResponse, error) {
	userRecord := h.userService.GetAllUsers(ctx, request)

	//var users []*userProto.GetUserByIdResponse
	//for _, user := range userRecord.Users {
	//	serializedUser := userProto.GetUserByIdResponse{
	//		Id:        user.Id,
	//		Username:  user.Username,
	//		CreatedAt: user.CreatedAt,
	//		Email:     user.Email,
	//		Image:     user.Image,
	//	}
	//
	//	users = append(users, &serializedUser)
	//}

	return userRecord, nil
}

func NewGrpcUserService(grpc *grpc.Server, userService service.UserService) {
	gRPCHandler := &UserGrpcHandler{
		userService: userService,
	}
	userProto.RegisterUserServiceServer(grpc, gRPCHandler)
}
