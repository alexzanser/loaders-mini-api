package repository

import (
	"context"
	"loaders/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type Customer interface{
	GetCustomer(ctx context.Context, username, passwd string) (*models.Customer, error)
	GetCustomersList(ctx context.Context) ([]int64, error)
}

type Loader interface{
	GetLoader(ctx context.Context, username, passwd string) (*models.Loader, error)
	GetLoadersList(ctx context.Context) ([]models.Loader, error) 
	UpdateLoader(ctx context.Context, ld *models.Loader) (error) 
}

type Registration interface {
	CreateCustomer(ctx context.Context, c *models.Customer) (int64, error)
	CreateLoader(ctx context.Context, l *models.Loader) (int64, error)
}

type Tasker interface{
	CreateTask(ctx context.Context, userID int64, task *models.Task) (error)
	CompleteTask(ctx context.Context, task *models.Task) (error) 
}

type Repository struct {
	Registration 
	Customer
	Loader
	Tasker
	logger	*log.Logger
}

func NewRepository(pgxPool *pgxpool.Pool, log *log.Logger) *Repository {
	return &Repository{
		logger: log,
		Registration: newRegisterRepo(pgxPool),
		Customer: newCustomerRepo(pgxPool),
		Loader: newLoaderRepo(pgxPool),
		Tasker: newTaskRepo(pgxPool),
	}
}
