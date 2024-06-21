package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
}

func TestNewUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456")
	assert.NotEmpty(t, user)
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("12345"))
	assert.NotEqual(t, user.Password, "123456") //garante que o password nao é o passado acima e sim o que passou pelo bcrypt
}
