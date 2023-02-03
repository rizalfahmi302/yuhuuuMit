package service

import (
	"log"
	"yuhuuuMit/feature/user"
	"yuhuuuMit/helper"

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

func (us *userService) RegisterServ(newUser user.Core) (user.Core, error) {
	// check input validation
	if err := us.vld.Struct(newUser); err != nil {
		log.Println("error service validation ==>", err)
		return user.Core{}, err
	}

	// bcrypt password before insert into database
	passBcrypt, err := helper.PassBcrypt(newUser.Password)
	if err != nil {
		log.Println("error service bcrypt ==>", err)
		return user.Core{}, err
	}
	newUser.Password = passBcrypt

	res, err := us.qry.RegisterRepo(newUser)
	if err != nil {
		log.Println("error service ==>", err)
		return user.Core{}, err
	}
	return res, nil
}
