package repository

import (
	"context"
	"database/sql"

	helper "github.com/FadhliAbrar/financial-tracker-apps/Helper"
	domain "github.com/FadhliAbrar/financial-tracker-apps/Model/Domain"
)

type categoryRepository struct {
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (repository *categoryRepository) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	const query string = "INSERT INTO Category(name, type) VALUES(?,?)"

	result, err := tx.ExecContext(ctx, query, category.Name, category.Type)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	return domain.Category{Id: int(id)}
}

func (repository *categoryRepository) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	const query string = "UPDATE Category SET name = ?, type = ? WHERE id = ?"

	result, err := tx.ExecContext(ctx, query, category.Name, category.Type, category.Id)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	return domain.Category{Id: int(id)}
}

func (repository *categoryRepository) DeleteById(ctx context.Context, tx *sql.Tx, categoryId int) {
	const query string = "DELETE FROM Category WHERE id = ?;"

	_, err := tx.ExecContext(ctx, query, categoryId)
	helper.PanicIfError(err)
}

func (repository *categoryRepository) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	const query string = "SELECT (id, name, type, created_at, updated_at) FROM Category WHERE id = ?"
	category := domain.Category{}

	row := tx.QueryRowContext(ctx, query, categoryId)
	err := row.Scan(&category.Id, &category.Name, &category.Type, &category.Created_at, &category.Updated_at)
	if err != nil {
		defer helper.PanicIfError(err)
		return domain.Category{}, err
	}

	return category, nil
}

func (repository *categoryRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	const query string = "SELECT (id, name, type, created_at, updated_at) FROM Category"
	categories := []domain.Category{}
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()
	if rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name, &category.Type, &category.Created_at, &category.Updated_at)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
