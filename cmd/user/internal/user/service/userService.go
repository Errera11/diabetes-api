package service

import (
	"context"
	"fmt"
	userProto "github.com/Errera11/user/internal/protogen"
	"github.com/Errera11/user/internal/user/repository"
)

type UserService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserById(ctx context.Context, id int32) *userProto.GetUserByIdResponse {
	fmt.Println("Getting user by id...")

	userRecord, err := s.userRepo.GetUserById(ctx, id)

	if err != nil {
		fmt.Println(`Error: error getting user by id %v`, id)
	}

	return userRecord
}

func (s *UserService) CreateUser(ctx context.Context, user *userProto.CreateUserRequest) *userProto.CreateUserResponse {
	userId, err := s.userRepo.CreateUser(ctx, user)

	if err != nil {
		fmt.Println(`Cant create user `, user.Email, err)
	}

	return &userProto.CreateUserResponse{
		UserId: userId,
	}
}

func (s *UserService) GetAllUsers(ctx context.Context, pagination *userProto.Pagination) *userProto.GetAllUsersResponse {
	users, err := s.userRepo.GetAllUsers(ctx)

	if err != nil {
		fmt.Println(`Cant get all users `, err)
	}

	return users
}
