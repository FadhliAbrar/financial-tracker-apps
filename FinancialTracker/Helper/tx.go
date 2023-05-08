package helper

import (
	"database/sql"
)

func Trxn(tx *sql.Tx) {
	err := recover()
	if err != nil {
		err := tx.Rollback()
		PanicIfError(err)
	} else {
		CommitErr := tx.Commit()
		PanicIfError(CommitErr)
	}
}
