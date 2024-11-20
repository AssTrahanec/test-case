package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"testCase"
)

type UserPostgres struct {
	db *sqlx.DB
}

var (
	ErrReferrerNotFound   = errors.New("referrer not found")
	ErrReferrerAlreadySet = errors.New("referrer already set")
)

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetUserById(userId string) (testCase.User, error) {
	var user testCase.User

	query := `
		SELECT id, username, balance, referrer_id 
		FROM users 
		WHERE id = $1
	`
	err := r.db.QueryRow(query, userId).Scan(
		&user.ID,
		&user.UserName,
		&user.Balance,
		&user.ReferrerID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return testCase.User{}, ErrUserNotFound
		}
		return testCase.User{}, err
	}

	return user, nil
}
func (r *UserPostgres) GetLeaderboard() ([]testCase.User, error) {
	var users []testCase.User

	query := `
		SELECT id, username, balance, referrer_id
		FROM users
		ORDER BY balance DESC
		LIMIT 10
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user testCase.User
		err := rows.Scan(
			&user.ID,
			&user.UserName,
			&user.Balance,
			&user.ReferrerID,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
func (r *UserPostgres) AddReferrer(userID, referrerID string) error {
	var userExists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", userID).Scan(&userExists)
	if err != nil {
		return err
	}
	if !userExists {
		return ErrUserNotFound
	}

	var referrerExists bool
	err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", referrerID).Scan(&referrerExists)
	if err != nil {
		return err
	}
	if !referrerExists {
		return ErrReferrerNotFound
	}

	var referrerAlreadySet bool
	err = r.db.QueryRow("SELECT referrer_id IS NOT NULL FROM users WHERE id = $1", userID).Scan(&referrerAlreadySet)
	if err != nil {
		return err
	}
	if referrerAlreadySet {
		return ErrReferrerAlreadySet
	}

	query := "UPDATE users SET referrer_id = $1 WHERE id = $2"
	_, err = r.db.Exec(query, referrerID, userID)
	return err
}
