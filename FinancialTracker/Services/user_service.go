package services

import (
	"context"

	web "github.com/FadhliAbrar/financial-tracker-apps/Model/Web"
)

type UserService interface {
	CreateUser(ctx context.Context, web web.UserRequest) web.UserResponse
	Login(ctx context.Context, web web.UserRequest) web.UserResponse
	FindUserById(ctx context.Context, id int) web.UserResponse
	UpdateUser(ctx context.Context, web web.UserRequest) web.UserResponse
	DeleteUserById(ctx context.Context, id int)
}
