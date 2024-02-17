package repository

import (
	"gorm.io/gorm"
	"sternx/domain"
	"sternx/infrastructure/logger"
)

type User interface {
	BaseRepository[domain.User]
}

func NewUser(tx *gorm.DB) User {
	return NewBaseRepository[domain.User](
		tx,
		logger.NewSubLogger("_user repository", nil),
	)
}
