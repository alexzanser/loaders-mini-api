package service

import (
	"loaders/internal/repository"
	"context"
)

type Service struct {
	repository.Repository
	Tasker
	Customer
}

type Tasker interface {
	GenerateRandomTasks(ctx context.Context) ([]int64, error)
}

type Customer interface {
	Start(ctx context.Context, loadersID []int64, username, passwd string) (string, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Repository: *repo,
		Tasker: newTaskService(repo),
		Customer: newCustomerService(repo),
	}
}
