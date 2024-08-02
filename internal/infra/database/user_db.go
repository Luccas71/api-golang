package database

import (
	"github.com/Luccas1/api-golang/internal/entity"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

// devolve uma conexao com banco de dados
func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{DB: db}
}

func (u *UserDB) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *UserDB) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := u.DB.Where("email= ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
