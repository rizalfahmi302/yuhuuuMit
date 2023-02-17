package service

import (
	"errors"
	"log"
	"syedara/feature/user"
	h "syedara/helper"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	qry user.UserRepository
	vld *validator.Validate
}

func New(ur user.UserRepository, v *validator.Validate) user.UserService {
	return &userService{
		qry: ur,
		vld: v,
	}
}

func (us *userService) RegisterSrv(newUser user.Core) (user.Core, error) {
	// check input validation
	if err := us.vld.Struct(newUser); err != nil {
		log.Println("error service validation ==>", err)
		return user.Core{}, err
	}

	// bcrypt password before insert into database
	passBcrypt, err := h.PassBcrypt(newUser.Password)
	if err != nil {
		log.Println("error service bcrypt ~", err)
		return user.Core{}, err
	}
	newUser.Password = passBcrypt

	res, err := us.qry.RegisterRepo(newUser)
	if err != nil {
		log.Println("error service ~", err)
		return user.Core{}, err
	}

	return res, nil
}

func (us *userService) LoginSrv(email, password string) (string, user.Core, error) {
	if email == "" {
		return "", user.Core{}, errors.New("Email cannot be empty")
	} else if password == "" {
		return "", user.Core{}, errors.New("Password cannot be empty")
	}

	res, err := us.qry.LoginRepo(email)
	if err != nil {
		log.Println("error service ~", err)
		return "", user.Core{}, err
	}

	if err := h.PassCompare(res.Password, password); err != nil {
		log.Println("error pass compare service ~", err)
		return "", user.Core{}, err
	}

	token, err := h.TokenGenerate(res.ID)
	if err != nil {
		log.Println("error generate token service ~", err)
		return "", user.Core{}, err
	}

	return token, res, nil
}
