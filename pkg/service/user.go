package service

import (
	"testCase"
	"testCase/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}
func (s *UserService) GetUserById(userId string) (testCase.User, error) {
	return s.repo.GetUserById(userId)
}
func (s *UserService) GetLeaderboard() ([]testCase.User, error) {
	return s.repo.GetLeaderboard()
}
func (s *UserService) AddReferrer(userID, ReferrerID string) error {
	return s.repo.AddReferrer(userID, ReferrerID)
}
