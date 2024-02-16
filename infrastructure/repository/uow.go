package repository

import (
	"context"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"sternx/domain"
	"sternx/infrastructure/logger"
)

var ErrSetupUnitOfWork = errors.New("could not setup unit of work")

type UnitOfWork interface {
	Do(ctx context.Context, isTransactional bool, fn UnitOfWorkBlock) error
	Setup(connString string) error
}

type unitOfWork struct {
	conn   *gorm.DB
	logger *logger.SubLogger
}

type UnitOfWorkStore interface {
	UserRepository() User
	Rollback() error
	Commit() error
}

type uowStore struct {
	userRepository User
	tx             *gorm.DB
}

type UnitOfWorkBlock func(UnitOfWorkStore) error

func (u *uowStore) UserRepository() User {
	return u.userRepository
}

func (u *uowStore) Rollback() error {
	return u.tx.Rollback().Error
}

func (u *uowStore) Commit() error {
	return u.tx.Commit().Error
}

func NewUnitOfWork() UnitOfWork {
	return &unitOfWork{
		logger: logger.NewSubLogger("_uow", nil),
	}
}

func (uow *unitOfWork) Setup(connString string) error {
	var err error
	if uow.conn, err = gorm.Open(
		postgres.Open(connString),
		&gorm.Config{
			SkipDefaultTransaction:   true,
			PrepareStmt:              true,
			DisableNestedTransaction: true,
			Logger:                   gormLogger.Default.LogMode(gormLogger.Silent),
		},
	); err != nil {
		uow.logger.Error("could not setup database connection", "error", err.Error())

		return ErrSetupUnitOfWork
	}

	uow.logger.Info("unit of work layer connected to database successfully")

	err = uow.autoMigrate()
	if err != nil {
		uow.logger.Error("could not migrate tables", "error", err.Error())

		return err
	}

	uow.logger.Info("tables has been migrated successfully")

	return nil
}

func (uow *unitOfWork) Do(ctx context.Context, isTransactional bool, fn UnitOfWorkBlock) error {
	var tx *gorm.DB

	if isTransactional {
		tx = uow.conn.WithContext(ctx).Begin()
	} else {
		tx = uow.conn.WithContext(ctx)
	}

	uows := &uowStore{
		userRepository: NewUser(tx),
		tx:             tx,
	}

	return fn(uows)
}

func (uow *unitOfWork) autoMigrate() error {
	return uow.conn.AutoMigrate(
		&domain.User{},
	)
}
