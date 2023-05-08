package services

import (
	"context"
	"database/sql"

	helper "github.com/FadhliAbrar/financial-tracker-apps/Helper"
	domain "github.com/FadhliAbrar/financial-tracker-apps/Model/Domain"
	web "github.com/FadhliAbrar/financial-tracker-apps/Model/Web"
	Repository "github.com/FadhliAbrar/financial-tracker-apps/Repository"
)

type CategoryServiceImpl struct {
	DB         *sql.DB
	Repository Repository.CategoryRepository
}

func NewCategoryService(DB *sql.DB, Repository Repository.CategoryRepository) CategoryServices {
	return &CategoryServiceImpl{
		DB:         DB,
		Repository: Repository,
	}
}

func (service *CategoryServiceImpl) CreateCategory(ctx context.Context, web web.CategoryWebRequest) web.CategoryResponses {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.Trxn(tx)

	flowData, err := domain.ToFlowData(web.Type)
	helper.PanicIfError(err)
	categoryDomain := domain.Category{
		Name: web.Name,
		Type: flowData,
	}
	category := service.Repository.Save(ctx, tx, categoryDomain)

	return helper.ToCategoryWebResponse(category)
}
func (service *CategoryServiceImpl) FindCategoryById(ctx context.Context, id int) (web.CategoryResponses, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.Trxn(tx)

	category, err := service.Repository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	return helper.ToCategoryWebResponse(category), nil
}
func (service *CategoryServiceImpl) FindAllCategory(ctx context.Context) []web.CategoryResponses {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.Trxn(tx)

	category := service.Repository.FindAll(ctx, tx)

	return helper.ToCategoryWebResponses(category)

}
func (service *CategoryServiceImpl) UpdateCategoryById(ctx context.Context, web web.CategoryWebRequest) web.CategoryResponses {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.Trxn(tx)

	flowData, err := domain.ToFlowData(web.Type)
	helper.PanicIfError(err)
	categoryDomain := domain.Category{
		Name: web.Name,
		Type: flowData,
	}
	category := service.Repository.Update(ctx, tx, categoryDomain)

	return helper.ToCategoryWebResponse(category)
}
func (service *CategoryServiceImpl) DeleteCategoryById(ctx context.Context, id int) {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.Trxn(tx)

	service.Repository.DeleteById(ctx, tx, id)
}
