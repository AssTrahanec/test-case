package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testCase"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user testCase.User) (string, error) {
	var id string

	query := fmt.Sprintf("insert into users (username, password_hash) values ($1, $2) returning id")
	row := r.db.QueryRow(query, user.UserName, user.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}
func (r *AuthPostgres) GetUser(username, password string) (testCase.User, error) {
	var user testCase.User
	query := fmt.Sprintf("select id from users where username = $1 and password_hash = $2")
	err := r.db.Get(&user, query, username, password)
	return user, err
}
