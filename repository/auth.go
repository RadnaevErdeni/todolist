package repository

import (
	todo "TODO"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthDB struct {
	db *sqlx.DB
}

func NewAuthDB(db *sqlx.DB) *AuthDB {
	return &AuthDB{db: db}
}

func (r *AuthDB) CreateUser(user todo.User) (int, error) {
	var id int
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	searchuser := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password = $2", usersTable)
	row := tx.QueryRow(searchuser, user.Username, user.Password)
	if err := row.Scan(&id); err == nil {
		tx.Rollback()
		return 0, nil
	}
	createuser := fmt.Sprintf("INSERT INTO %s (username, password) values ($1, $2) RETURNING id", usersTable)
	row = tx.QueryRow(createuser, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, nil
	}
	return id, tx.Commit()
}

func (r *AuthDB) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password = $2", usersTable)

	err := r.db.Get(&user, query, username, password)

	return user, err
}
