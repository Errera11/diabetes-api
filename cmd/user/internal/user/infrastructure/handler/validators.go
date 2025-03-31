package handler

type UserRequestValidator struct {
	UserId int32 `validate:"required,gte=0"`
}

type CreateUserValidator struct {
	Username string `validate:"required,gte=3"`
	Password string `validate:"required,gte=8"`
	Email    string `validate:"required,email"`
	Image    string
}
