package service

import (
	"testCase"
	"testCase/pkg/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CompleteTask(userId, taskId string) error {
	return s.repo.CompleteTask(userId, taskId)
}
func (s *TaskService) GetCompletedTasksByUserID(userID string) ([]testCase.Task, error) {
	return s.repo.GetCompletedTasksByUserID(userID)
}
