package service

import (
	"context"
	"fmt"
	"loaders/internal/models"
	"loaders/internal/repository"
)

type taskService struct {
	repo *repository.Repository
}

func newTaskService (repo *repository.Repository) *taskService {
	return &taskService{repo: repo}
}

func (t *taskService) GenerateRandomTasks(ctx context.Context) ([]int64, error) {
	idList, err := t.repo.GetCustomersList(ctx)
	if err != nil {
		return nil, fmt.Errorf("error when get customers list: %v", err)
	}

	if len(idList) == 0 {
		return nil, fmt.Errorf("error: no customers exist")
	}
	for _ , id := range idList {
		for  i:= 0; i < 4; i++ {
			task := models.NewTask()
			t.repo.CreateTask(ctx, id, task)
		}
	}

	return idList, nil
}