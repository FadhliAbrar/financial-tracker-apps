package repository

import (
	"context"
	"database/sql"

	domain "github.com/FadhliAbrar/financial-tracker-apps/Model/Domain"
)

type TransactionRepository interface {
	GetAllTransaction(ctx context.Context, tx *sql.Tx) []domain.Transaction
	AddTransaction(ctx context.Context, tx *sql.Tx, trx domain.Transaction) domain.Transaction
	UpdateTransaction(ctx context.Context, tx *sql.Tx, trx domain.Transaction) domain.Transaction
	DeleteTransactionById(ctx context.Context, tx *sql.Tx, id int)
}
