package service

import (
	"context"
	"fmt"
	"loaders/internal/models"
	"loaders/internal/repository"
)

type customerService struct {
	repo *repository.Repository
}

func newCustomerService (repo *repository.Repository) *customerService {
	return &customerService{repo: repo}
}

func (c *customerService) Start(ctx context.Context, loadersID []int64, username, passwd string) (bool, error) {
	ct, err := c.repo.GetCustomer(ctx, username, passwd)
	if err != nil {
		return false, fmt.Errorf("error when start game:%v", err)
	}

	if len (ct.Tasks) == 0 {
		return true, nil
	}

	task := ct.Tasks[0]
	allLoaders, err := c.repo.GetLoadersFull(ctx)
	if err != nil {
		return false, fmt.Errorf("error when get loaders for task:%v", err)
	}
	choosenLoaders := make([]models.Loader, 0)
	totalCost := 0
	totalWeight := 0
	for _, ld := range allLoaders {
		choosenLoaders = append(choosenLoaders, ld)
	}

	for _, ld := range choosenLoaders {
		totalCost += ld.Salary
		totalWeight += ld.MaxWeight
	}

	if totalCost > ct.Balance {
		return false, fmt.Errorf("game failed: not enough balance")
	}
	if totalWeight < task.Weight {
		return false, fmt.Errorf("game failed: loaders can't handle such a huge weight")
	}

	ct.Balance -= totalCost
	c.repo.UpdateCustomer(ctx, totalCost, ct)
	c.repo.CompleteTask(ctx, &task)
	for _, ld := range choosenLoaders {
		ld.CompletedTasks = append(ld.CompletedTasks, task)
		if err := c.repo.UpdateLoader(ctx, &ld); err != nil {
			return false, fmt.Errorf("error when update loaders:%v" , err)
		}
	}
	return true, nil
}
