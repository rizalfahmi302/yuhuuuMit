package repository

import (
	"yuhuuuMit/feature/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname    string
	Username    string `gorm:"unique"`
	Email       string `gorm:"unique"`
	Password    string
	Gender      string
	Avatar      string
	Sampul      string
	DateOfBirth string
	Bio         string
}

func CoreToUser(data user.Core) User {
	return User{
		Model:       gorm.Model{ID: data.ID},
		Fullname:    data.Fullname,
		Username:    data.Username,
		Email:       data.Email,
		Password:    data.Password,
		Gender:      data.Gender,
		Avatar:      data.Avatar,
		Sampul:      data.Sampul,
		DateOfBirth: data.DateOfBirth,
		Bio:         data.Bio,
	}
}

func UserToCore(data User) user.Core {
	return user.Core{
		ID:          data.ID,
		Fullname:    data.Fullname,
		Username:    data.Username,
		Email:       data.Email,
		Password:    data.Password,
		Gender:      data.Gender,
		Avatar:      data.Avatar,
		Sampul:      data.Sampul,
		DateOfBirth: data.DateOfBirth,
		Bio:         data.Bio,
	}
}
