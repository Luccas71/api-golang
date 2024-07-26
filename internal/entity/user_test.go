package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Joe Doe", "j@j.com", "12345")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, user.Name, "Joe Doe")
	assert.Equal(t, user.Email, "j@j.com")
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Joe Doe", "j@j.com", "12345")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("12345"))
	assert.False(t, user.ValidatePassword("123456"))
	assert.NotEqual(t, "12345", user.Password)
}
