package handler

import (
	"log"
	"net/http"
	"syedara/feature/user"
	h "syedara/helper"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type userControll struct {
	srv user.UserService
}

func New(srv *user.UserService) user.UserHandler {
	return &userControll{
		srv: *srv,
	}
}

func (uc *userControll) RegisterHdl() echo.HandlerFunc {
	return func(c echo.Context) error {
		registerInput := RegisterReq{}
		if err := c.Bind(&registerInput); err != nil {
			log.Println("error bind ==>", err)
			return c.JSON(h.ErrorResponse(err))
		}

		newUser := user.Core{}
		copier.Copy(&newUser, &registerInput)

		_, err := uc.srv.RegisterSrv(newUser)
		if err != nil {
			log.Println("error handler ==>", err)
			return c.JSON(h.ErrorResponse(err))
		}

		return c.JSON(h.SuccessResponse(http.StatusCreated, "Success register account"))
	}
}

func (uc *userControll) LoginHdl() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginInput := LoginReq{}
		if err := c.Bind(&loginInput); err != nil {
			log.Println("error bind ==>", err)
			return c.JSON(h.ErrorResponse(err))
		}

		token, res, err := uc.srv.LoginSrv(loginInput.Email, loginInput.Password)
		if err != nil {
			return c.JSON(h.ErrorResponse(err))
		}

		loginRes := LoginRes{}
		copier.Copy(&loginRes, &res)

		return c.JSON(h.SuccessResponse(http.StatusOK, "Login success", loginRes, token))
	}
}
