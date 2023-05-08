package services

import (
	"context"

	web "github.com/FadhliAbrar/financial-tracker-apps/Model/Web"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, web web.TransactionRequest) web.TransactionResponse
	FindAllTransaction(ctx context.Context) []web.TransactionResponse
	UpdateTransaction(ctx context.Context, web web.TransactionRequest) web.TransactionResponse
	DeleteTransactionById(ctx context.Context, id int)
}
