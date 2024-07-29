package entity

import (
	"github.com/Luccas1/api-golang/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

// type User struct {
// 	ID       entity.ID `json:"id"`
// 	Name     string    `json:"name"`
// 	Email    string    `json:"email"`
// 	Password string    `json:"-"` //password nunca mostrado para o usuário final
// }

// func NewUser(name, email, password string) (*User, error) {
// 	//protegendo a senha com hash
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &User{
// 		ID:       entity.NewID(),
// 		Name:     name,
// 		Email:    email,
// 		Password: string(hash),
// 		//transformando o slice de byte em string para salvar no DB
// 	}, nil
// }

// // comparando e validando se a senha guardada é igual a senha passada pelo usuario
// func (u *User) ValidatePassword(password string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
// 	return err == nil
// }

type User struct {
	ID       entity.ID
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) (*User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
