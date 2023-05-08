package repository

import (
	"context"
	"database/sql"

	domain "github.com/FadhliAbrar/financial-tracker-apps/Model/Domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	DeleteById(ctx context.Context, tx *sql.Tx, categoryId int)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
