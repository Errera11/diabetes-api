package repository

import (
	"context"
	userProto "github.com/Errera11/user/internal/protogen"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *userProto.CreateUserRequest) (int32, error)
	GetUserById(ctx context.Context, id int32) (*userProto.GetUserByIdResponse, error)
	GetUserByEmail(ctx context.Context, email string) (*userProto.GetUserByIdResponse, error)
	UpdateUser(ctx context.Context, user *userProto.User) error
	DeleteUser(ctx context.Context, id int32) error
	GetAllUsers(ctx context.Context) (*userProto.GetAllUsersResponse, error)
}
