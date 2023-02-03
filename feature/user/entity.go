package user

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint
	Fullname    string `validate:"required,min=6"`
	Username    string `validate:"required,min=6"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required,min=3"`
	Gender      string `validate:"required"`
	Avatar      string
	Sampul      string
	DateOfBirth string
	Bio         string
}

type UserHandler interface {
	RegisterHdl() echo.HandlerFunc
}

type UserService interface {
	RegisterServ(newUser Core) (Core, error)
}

type UserRepository interface {
	RegisterRepo(newUser Core) (Core, error)
}