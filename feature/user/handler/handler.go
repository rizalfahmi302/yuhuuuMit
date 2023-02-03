package handler

import (
	"log"
	"net/http"
	"yuhuuuMit/feature/user"
	"yuhuuuMit/helper"

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
			return c.JSON(helper.ErrorResponse(err))
		}

		newUser := user.Core{}
		copier.Copy(&newUser, &registerInput)

		_, err := uc.srv.RegisterServ(newUser)
		if err != nil {
			log.Println("error handler ==>", err)
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(helper.SuccessResponse(http.StatusCreated, "success register account"))
	}
}
