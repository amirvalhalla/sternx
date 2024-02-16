package grpc

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/scrypt"
	"sternx/config"
	"sternx/domain"
	"sternx/infrastructure/logger"
	"sternx/infrastructure/query"
	"sternx/infrastructure/repository"
	"sternx/infrastructure/util"
	sternx "sternx/www/grpc/gen/go"
)

var ErrInternalServerError = errors.New("something went wrong, please try again later")

type userServer struct {
	*Server
	*logger.SubLogger
	*config.Config
	repository.UnitOfWork
}

func newUserServer(
	server *Server,
	uow repository.UnitOfWork,
	cfg *config.Config,
) *userServer {
	return &userServer{
		Server:     server,
		SubLogger:  logger.NewSubLogger("_user", nil),
		Config:     cfg,
		UnitOfWork: uow,
	}
}

func (s *userServer) CreateUser(ctx context.Context, request *sternx.CreateUserRequest) (*sternx.UserResponse, error) {
	if util.IsStringEmpty(request.Email) || util.IsStringEmpty(request.Password) {
		return nil, errors.New("email or password could not be empty")
	}

	hashedPassword, err := scrypt.Key([]byte(request.Password), s.config.PasswordSalt, 16384, 8, 1, 32)
	if err != nil {
		return nil, ErrInternalServerError
	}

	userEntity := domain.NewUser(request.Email, hashedPassword)

	err = s.uow.Do(ctx, true, func(uows repository.UnitOfWorkStore) error {
		findUser, _ := uows.UserRepository().FindOne(query.FindByCustomColumn("email", request.Email))

		if findUser.ID != uuid.Nil {
			return errors.New("the email is duplicate")
		}

		_, err = uows.UserRepository().Insert(userEntity)
		if err != nil {
			_ = uows.Rollback()

			return err
		}

		return uows.Commit()
	})
	if err != nil {
		return nil, err
	}

	return &sternx.UserResponse{
		Id:    userEntity.ID.String(),
		Email: userEntity.Email,
	}, nil
}

func (s *userServer) GetUserByID(ctx context.Context, request *sternx.GetUserRequest) (*sternx.UserResponse, error) {
	if util.IsStringEmpty(request.Id) {
		return nil, errors.New("id could not be empty")
	}

	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	var userEntity domain.User
	err = s.uow.Do(ctx, false, func(uows repository.UnitOfWorkStore) error {
		userEntity, err = uows.UserRepository().FindOne(query.FindByID(id))

		return err
	})
	if err != nil {
		return nil, err
	}

	return &sternx.UserResponse{
		Id:    userEntity.ID.String(),
		Email: userEntity.Email,
	}, nil
}

func (s *userServer) GetUsers(ctx context.Context, request *sternx.GetUsersRequest) (*sternx.GetUsersResponse, error) {
	var userEntities []domain.User
	var totalRecords int64
	var err error

	err = s.uow.Do(ctx, false, func(uows repository.UnitOfWorkStore) error {
		userEntities, totalRecords, err = uows.UserRepository().FindAllWithTotalRecords(
			query.WithoutAnySearch(),
			query.WithOrderBy("created_at", query.DESC),
			query.WithOffset(int(request.PageIndex), int(request.PageSize)),
			query.WithLimit(int(request.PageSize)),
		)

		return err
	})
	if err != nil {
		return nil, err
	}

	users := make([]*sternx.UserResponse, 0, len(userEntities))
	for i := range userEntities {
		ur := sternx.UserResponse{
			Id:    userEntities[i].ID.String(),
			Email: userEntities[i].Email,
		}
		users = append(users, &ur)
	}

	return &sternx.GetUsersResponse{
		Users:        users,
		PageIndex:    request.PageIndex,
		PageSize:     request.PageSize,
		TotalRecords: uint64(totalRecords),
	}, nil
}

func (s *userServer) UpdateUser(ctx context.Context, request *sternx.UpdateUserRequest) (*sternx.UserResponse, error) {
	if util.IsStringEmpty(request.Id) {
		return nil, errors.New("id could not be empty")
	}

	if util.IsStringEmpty(request.Email) && util.IsStringEmpty(request.Password) {
		return nil, errors.New("both email and password cannot be empty together")
	}

	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	var findUser domain.User
	err = s.uow.Do(ctx, true, func(uows repository.UnitOfWorkStore) error {
		findUser, err = uows.UserRepository().FindOne(query.FindByID(id))
		if err != nil {
			return err
		}

		if !util.IsStringEmpty(request.Email) {
			// check email exists before or not
			tempUser, _ := uows.UserRepository().FindOne(query.FindByCustomColumn("email", request.Email))
			if tempUser.ID != uuid.Nil {
				return errors.New("email is duplicate")
			}

			findUser.Email = request.Email
		}

		if !util.IsStringEmpty(request.Password) {
			hashedPassword, err := scrypt.Key([]byte(request.Password), s.config.PasswordSalt, 16384, 8, 1, 32)
			if err != nil {
				return ErrInternalServerError
			}
			findUser.Password = hashedPassword
		}

		result, err := uows.UserRepository().Update(&findUser)
		if err != nil {
			_ = uows.Rollback()

			return err
		}

		findUser.Email = result.Email

		return uows.Commit()
	})
	if err != nil {
		return nil, err
	}

	return &sternx.UserResponse{
		Id:    findUser.ID.String(),
		Email: findUser.Email,
	}, nil
}

func (s *userServer) DeleteUser(ctx context.Context, request *sternx.DeleteUserRequest) (*sternx.DeleteUserResponse,
	error,
) {
	if util.IsStringEmpty(request.Id) {
		return nil, errors.New("id could not be empty")
	}

	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	err = s.uow.Do(ctx, true, func(uows repository.UnitOfWorkStore) error {
		findUser, err := uows.UserRepository().FindOne(query.FindByID(id))
		if err != nil {
			return err
		}

		err = uows.UserRepository().SoftDelete(&findUser)
		if err != nil {
			_ = uows.Rollback()

			return err
		}

		return uows.Commit()
	})
	if err != nil {
		return nil, err
	}

	return &sternx.DeleteUserResponse{}, nil
}
