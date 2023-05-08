package repository

import (
	"context"
	"database/sql"

	helper "github.com/FadhliAbrar/financial-tracker-apps/Helper"
	domain "github.com/FadhliAbrar/financial-tracker-apps/Model/Domain"
)

type transactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepositoryImpl{}
}

func (repository *transactionRepositoryImpl) GetAllTransaction(ctx context.Context, tx *sql.Tx) []domain.Transaction {

	var query string = "SELECT (id, type, amount, category_id, date) FROM Transaction"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()
	arrTransaction := []domain.Transaction{}

	if rows.Next() {
		transaction := domain.Transaction{}
		err := rows.Scan(&transaction.Id, &transaction.Type, &transaction.Amount, &transaction.Category_id, &transaction.Date)
		helper.PanicIfError(err)
		arrTransaction = append(arrTransaction, transaction)
	} else {
		return arrTransaction
	}
	return arrTransaction
}
func (repository *transactionRepositoryImpl) AddTransaction(ctx context.Context, tx *sql.Tx, trx domain.Transaction) domain.Transaction {
	var query string = "INSERT INTO Transaction(user_id, category_id, amount, description, date, type) VALUES(?, ?, ?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, query, trx.User_id, trx.Category_id, trx.Amount, trx.Description, trx.Date, trx.Type)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	Trxc := domain.Transaction{
		Id: int(id),
	}
	return Trxc
}
func (repository *transactionRepositoryImpl) UpdateTransaction(ctx context.Context, tx *sql.Tx, trx domain.Transaction) domain.Transaction {
	var query string = "UPDATE Transaction SET user_id = ?, category_id = ?, amount = ?, description = ?, date = ?, type = ? WHERE id = ?"

	result, err := tx.ExecContext(ctx, query, trx.User_id, trx.Category_id, trx.Amount, trx.Description, trx.Date, trx.Type, trx.Id)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	Trxc := domain.Transaction{
		Id: int(id),
	}
	return Trxc
}
func (repository *transactionRepositoryImpl) DeleteTransactionById(ctx context.Context, tx *sql.Tx, id int) {
	var query string = "DELETE FROM Transaction WHERE id = ?;"

	_, err := tx.ExecContext(ctx, query, id)
	helper.PanicIfError(err)
}
