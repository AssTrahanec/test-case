package service

import (
	"testCase"
	"testCase/pkg/repository"
)

type Task interface {
	CompleteTask(userId, taskId string) error
	GetCompletedTasksByUserID(userID string) ([]testCase.Task, error)
}

type User interface {
	GetUserById(userId string) (testCase.User, error)
	GetLeaderboard() ([]testCase.User, error)
	AddReferrer(userID, ReferrerID string) error
}

type Authorization interface {
	CreateUser(testCase.User) (string, error)
	GenerateToken(testCase.User) (string, error)
	ParseToken(accessToken string) (string, error)
}

type Service struct {
	Authorization
	Task
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Task:          NewTaskService(repos.Task),
		User:          NewUserService(repos.User),
	}
}
