package services

import (
	"context"
	"database/sql"

	helper "github.com/FadhliAbrar/financial-tracker-apps/Helper"
	domain "github.com/FadhliAbrar/financial-tracker-apps/Model/Domain"
	web "github.com/FadhliAbrar/financial-tracker-apps/Model/Web"
	repository "github.com/FadhliAbrar/financial-tracker-apps/Repository"
)

type transactionServiceImpl struct {
	DB         *sql.DB
	Repository repository.TransactionRepository
}

func NewTransactionService(db *sql.DB, repo repository.TransactionRepository) TransactionService {
	return &transactionServiceImpl{
		DB:         db,
		Repository: repo,
	}
}

func (service *transactionServiceImpl) CreateTransaction(ctx context.Context, web web.TransactionRequest) web.TransactionResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.Trxn(tx)

	flowType, err := domain.ToFlowData(web.Type)
	helper.PanicIfError(err)

	trx := domain.Transaction{
		Type:        flowType,
		Amount:      web.Amount,
		Description: web.Description,
		Category_id: web.Category_id,
		Date:        web.Date,
	}

	Transaction := service.Repository.AddTransaction(ctx, tx, trx)
	return helper.ToTransactionResponse(Transaction)

}
func (service *transactionServiceImpl) FindAllTransaction(ctx context.Context) []web.TransactionResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.Trxn(tx)

	Transaction := service.Repository.GetAllTransaction(ctx, tx)
	return helper.ToTransactionResponses(Transaction)
}
func (service *transactionServiceImpl) UpdateTransaction(ctx context.Context, web web.TransactionRequest) web.TransactionResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.Trxn(tx)

	flowType, err := domain.ToFlowData(web.Type)
	helper.PanicIfError(err)

	trx := domain.Transaction{
		Type:        flowType,
		Amount:      web.Amount,
		Description: web.Description,
		Category_id: web.Category_id,
		Date:        web.Date,
	}
	Transaction := service.Repository.UpdateTransaction(ctx, tx, trx)
	return helper.ToTransactionResponse(Transaction)
}
func (service *transactionServiceImpl) DeleteTransactionById(ctx context.Context, id int) {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.Trxn(tx)

	service.Repository.DeleteTransactionById(ctx, tx, id)
}
