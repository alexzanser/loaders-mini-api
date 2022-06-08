package repository

import (
	"loaders/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"context"
	"fmt"
)

type registerRepo struct {
	pool *pgxpool.Pool
}

func newRegisterRepo(pool *pgxpool.Pool) *registerRepo{
	return &registerRepo{
		pool: pool,
	}
}

const createCustomerQuery = `INSERT INTO customers (username, passwd_hash, balance)
	VALUES ($1, $2, $3) RETURNING id;`

func (a registerRepo) CreateCustomer(ctx context.Context, c *models.Customer) (int64, error) {
	var customerID int64

	tx, err := a.pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	row := a.pool.QueryRow(ctx, createCustomerQuery, c.Username, c.PasswdHash, c.Balance)
	if err := row.Scan(&customerID); err != nil {
		return 0, fmt.Errorf("error adding data to database: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("error commiting transaction: %w", err)
	}

	return customerID, nil
}

const createLoaderQuery = `INSERT INTO loaders (username, passwd_hash, max_weight, drunk, fatigue, salary)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

func (a registerRepo) CreateLoader(ctx context.Context, l *models.Loader) (int64, error) {
	var loaderID int64

	tx, err := a.pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	row := a.pool.QueryRow(ctx, createLoaderQuery, l.Username, l.PasswdHash, l.MaxWeight, l.Drunk, l.Fatigue, l.Salary)
	if err := row.Scan(&loaderID); err != nil {
		return 0, fmt.Errorf("error adding data to database: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("error commiting transaction: %w", err)
	}
	
	return loaderID, nil
}
