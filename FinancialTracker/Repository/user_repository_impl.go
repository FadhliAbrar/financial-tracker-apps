package repository

import (
	"context"
	"database/sql"

	helper "github.com/FadhliAbrar/financial-tracker-apps/Helper"
	domain "github.com/FadhliAbrar/financial-tracker-apps/Model/Domain"
)

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (repository *userRepositoryImpl) AddUser(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	const query string = "INSERT INTO User(username, email, password) VALUES (?,?,?)"

	result, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	return domain.User{Id: int(id)}
}
func (repository *userRepositoryImpl) FindUserById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error) {
	const query string = "SELECT (username, email) FROM User WHERE id = ?"

	user := domain.User{}

	row := tx.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Username, &user.Email)
	defer helper.PanicIfError(err)
	if err != nil {
		return user, err
	}

	return user, nil
}
func (repository *userRepositoryImpl) DeleteUserById(ctx context.Context, tx *sql.Tx, id int) {
	const query string = "DELETE * FROM User WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, id)
	helper.PanicIfError(err)
}
func (repository *userRepositoryImpl) UpdateUser(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	const query string = "UPDATE User SET username = ?, email = ?, amount = ?, password = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.Id)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	return domain.User{Id: int(id)}
}
