package services

import (
	"context"
	"database/sql"

	helper "github.com/FadhliAbrar/financial-tracker-apps/Helper"
	domain "github.com/FadhliAbrar/financial-tracker-apps/Model/Domain"
	web "github.com/FadhliAbrar/financial-tracker-apps/Model/Web"
	repository "github.com/FadhliAbrar/financial-tracker-apps/Repository"
)

type userServiceImpl struct {
	DB         *sql.DB
	Repository repository.UserRepository
}

func NewUserService(db *sql.DB, repo repository.UserRepository) UserService {
	return &userServiceImpl{
		DB:         db,
		Repository: repo,
	}
}

func (service *userServiceImpl) CreateUser(ctx context.Context, web web.UserRequest) web.UserResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.Trxn(tx)

	user := domain.User{
		Username: web.Username,
		Email:    web.Email,
		Password: web.Password,
	}

	domainUser := service.Repository.AddUser(ctx, tx, user)
	return helper.ToUserResponse(domainUser)
}
func (service *userServiceImpl) Login(ctx context.Context, web web.UserRequest) web.UserResponse {
	domain := domain.User{
		Id:    123123,
		Email: "I'LL HANDLE IT LATER",
	}
	return helper.ToUserResponse(domain)
}
func (service *userServiceImpl) FindUserById(ctx context.Context, id int) web.UserResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.Trxn(tx)

	domainUser, err := service.Repository.FindUserById(ctx, tx, id)
	helper.PanicIfError(err)
	return helper.ToUserResponse(domainUser)
}
func (service *userServiceImpl) UpdateUser(ctx context.Context, web web.UserRequest) web.UserResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.Trxn(tx)

	user := domain.User{
		Username: web.Username,
		Email:    web.Email,
		Password: web.Password,
	}

	domainUser := service.Repository.UpdateUser(ctx, tx, user)
	return helper.ToUserResponse(domainUser)
}
func (service *userServiceImpl) DeleteUserById(ctx context.Context, id int) {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.Trxn(tx)

	service.Repository.DeleteUserById(ctx, tx, id)
}
