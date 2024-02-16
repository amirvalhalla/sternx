package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;unique"`
	Email    string
	Password []byte
	gorm.Model
}

func NewUser(email string, password []byte) *User {
	return &User{
		ID:       uuid.New(),
		Email:    email,
		Password: password,
	}
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
}

func (u *User) UpdatePassword(password []byte) {
	u.Password = password
}
