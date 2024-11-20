package repository

import (
	"github.com/jmoiron/sqlx"
	"testCase"
)

type User interface {
	GetUserById(userId string) (testCase.User, error)
	GetLeaderboard() ([]testCase.User, error)
	AddReferrer(userID, referrerID string) error
}

type Task interface {
	CompleteTask(userId, taskId string) error
	GetCompletedTasksByUserID(userID string) ([]testCase.Task, error)
}

type Authorization interface {
	CreateUser(user testCase.User) (string, error)
	GetUser(username, password string) (testCase.User, error)
}

type Repository struct {
	Authorization
	Task
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Task:          NewTaskPostgres(db),
		User:          NewUserPostgres(db),
	}
}
