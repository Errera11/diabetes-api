package service

import (
	"context"
	"fmt"
	userProto "github.com/Errera11/user/internal/protogen"
	"github.com/Errera11/user/internal/user/repository"
	"github.com/Errera11/user/utils"
)

type UserService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) *userProto.GetUserByIdResponse {
	fmt.Println("Getting user by id...")

	userRecord, err := s.userRepo.GetUserByEmail(ctx, email)

	if err != nil {
		fmt.Printf(`Error: error getting user by email %v %e`, email, err)
	}

	return userRecord
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
	passwordSalt := utils.GenRandomSalt()
	hashedUserPassword := utils.HashPassword(user.Password, passwordSalt)

	userPayload := &userProto.CreateUserRequest{
		Password: hashedUserPassword,
		Username: user.Username,
		Email:    user.Email,
		Image:    user.Image,
	}
	userId, err := s.userRepo.CreateUser(ctx, userPayload)

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
