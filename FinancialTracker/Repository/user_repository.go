package repository

import (
	"context"
	"database/sql"

	domain "github.com/FadhliAbrar/financial-tracker-apps/Model/Domain"
)

type UserRepository interface {
	AddUser(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindUserById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error)
	DeleteUserById(ctx context.Context, tx *sql.Tx, id int)
	UpdateUser(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}
