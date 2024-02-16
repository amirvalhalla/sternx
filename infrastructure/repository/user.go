package repository

import (
	"gorm.io/gorm"
	"sternx/domain"
)

type User interface {
	BaseRepository[domain.User]
}

func NewUser(tx *gorm.DB) User {
	return NewBaseRepository[domain.User](tx)
}
