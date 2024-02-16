package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;unique"`
	Email    string
	Password string
}

func NewUser(email, password string) *User {
	return &User{
		ID:       uuid.New(),
		Email:    email,
		Password: password,
	}
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
}

func (u *User) UpdatePassword(password string) {
	u.Password = password
}
