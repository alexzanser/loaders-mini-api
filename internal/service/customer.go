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

func (c *customerService) Start(ctx context.Context, loadersID []int64, username, passwd string) (string, error) {
	ct, err := c.repo.GetCustomer(ctx, username, passwd)
	if err != nil {
		return "Game failed!", fmt.Errorf("error when start game:%v", err)
	}

	if len (ct.Tasks) == 0 {
		return "Game successfully completed! All tasks are done!", nil
	}

	task := ct.Tasks[0]
	allLoaders, err := c.repo.GetLoadersFull(ctx)
	if err != nil {
		return "Game failed!", fmt.Errorf("error when get loaders for task:%v", err)
	}
	choosenLoaders := make([]models.Loader, 0)
	totalCost := 0
	totalWeight := 0
	
	for _, ld := range allLoaders {
		if contains(loadersID, ld.ID) {
			choosenLoaders = append(choosenLoaders, ld)
		}
	}

	for _, ld := range choosenLoaders {
		totalCost += ld.Salary
		totalWeight += ld.MaxWeight
	}

	if totalCost > ct.Balance {
		return "Game failed!", fmt.Errorf("game failed: not enough balance")
	}
	if totalWeight < task.Weight {
		return "Game failed!", fmt.Errorf("game failed: loaders can't handle such a huge weight")
	}

	err = c.repo.CompleteTask(ctx, &task)
	if err != nil {
		return "Game failed!", fmt.Errorf("game failed: can't complete task")
	}

	ct.UpdateBalance(totalCost)
	err = c.repo.UpdateCustomer(ctx, totalCost, ct)
	if err != nil {
		return "Game failed!", fmt.Errorf("game failed: can't update cutomer info")
	}

	for _, ld := range choosenLoaders {
		ld.CompletedTasks = append(ld.CompletedTasks, task)
		ld.Update()
		if err := c.repo.UpdateLoader(ctx, &ld); err != nil {
			return "Game failed!", fmt.Errorf("error when update loaders:%v" , err)
		}
	}

	if len (ct.Tasks) == 0 {
		return "Game successfully completed! All tasks are done!", nil
	}
	
	return "Task finished, congratulations! Still have some tasks to deal with.", nil
}

func contains(s []int64, e int64) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
