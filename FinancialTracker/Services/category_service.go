package services

import (
	"context"

	web "github.com/FadhliAbrar/financial-tracker-apps/Model/Web"
)

type CategoryServices interface {
	CreateCategory(ctx context.Context, web web.CategoryWebRequest) web.CategoryResponses
	FindCategoryById(ctx context.Context, id int) (web.CategoryResponses, error)
	FindAllCategory(ctx context.Context) []web.CategoryResponses
	UpdateCategoryById(ctx context.Context, web web.CategoryWebRequest) web.CategoryResponses
	DeleteCategoryById(ctx context.Context, id int)
}
