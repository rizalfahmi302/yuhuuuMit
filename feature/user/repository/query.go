package repository

import (
	"log"
	"syedara/feature/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserRepository {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) RegisterRepo(newUser user.Core) (user.Core, error) {
	cnv := CoreToUser(newUser)
	query := uq.db.Create(&cnv)
	if query.Error != nil {
		log.Println("error Create query~", query.Error)
		return user.Core{}, query.Error
	}

	newUser.ID = cnv.ID
	return newUser, nil
}

func (uq *userQuery) LoginRepo(email string) (user.Core, error) {
	res := User{}
	query := uq.db.Where("email = ?", email).First(&res)
	if query.RowsAffected < 1 {
		log.Println("error RowsAffected query ~", query.Error)
		return user.Core{}, query.Error
	}
	if query.Error != nil {
		log.Println("error First query ~", query.Error)
		return user.Core{}, query.Error
	}

	return UserToCore(res), nil
}
