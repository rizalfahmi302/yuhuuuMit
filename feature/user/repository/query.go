package repository

import (
	"log"
	"yuhuuuMit/feature/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserRepository {
	return &userQuery {
		db: db,
	}
}

func (uq *userQuery) RegisterRepo(newUser user.Core) (user.Core, error) {
	cnv := CoreToUser(newUser)
	if err := uq.db.Create(&cnv).Error; err != nil {
		log.Println("error query create", err)
		return user.Core{}, err
	}
	newUser.ID = cnv.ID
	return newUser, nil
}