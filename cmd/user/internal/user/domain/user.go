package user

import "time"

type User struct {
	Id        uint64
	Username  string
	Password  string
	Email     string
	Image     string
	CreatedAt time.Time
}
