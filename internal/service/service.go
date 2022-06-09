package service

import (
	"loaders/internal/repository"
	"context"
)

//Service структура представляющая собой объект для работы с уровнем бизнес-логики
type Service struct {
	repository.Repository
	Tasker
	Customer
}


//Tasker описывает необходимые методы для генерации заданий
type Tasker interface {
	GenerateRandomTasks(ctx context.Context) ([]int64, error)
}

//Customer описывает необходимые для работы игры методы
type Customer interface {
	Start(ctx context.Context, loadersID []int64, username, passwd string) (string, error)
}

//NewService создает новую структуру типа Service
func NewService(repo *repository.Repository) *Service {
	return &Service{
		Repository: *repo,
		Tasker: newTaskService(repo),
		Customer: newCustomerService(repo),
	}
}
