package handler

type SigninValidator struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,gte=8"`
}

type SignupValidator struct {
	Username string `validate:"required"`
	Password string `validate:"required,gte=8"`
	Email    string `validate:"required,email"`
	// TODO
	//optional string Image = 4;
}

type LogoutValidator struct {
	Token string `validate:"required"`
}
