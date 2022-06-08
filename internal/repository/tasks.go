package repository

import (
	"context"
	"fmt"
	"loaders/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type taskRepo struct {
	pool *pgxpool.Pool
}

func newTaskRepo(pool *pgxpool.Pool) *taskRepo{
	return &taskRepo{
		pool: pool,
	}
}

const (
	createTaskQuery = "INSERT INTO tasks (customer_id, name, weight) VALUES ($1, $2, $3);"
	completeTaskQuery = "INSERT INTO tasks (completed) VALUES ($1) WHERE id=$2;"
)

func (t *taskRepo) CreateTask(ctx context.Context, userID int64, task *models.Task) (error) {
	tx, err := t.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	res, err := tx.Exec(ctx, createTaskQuery, userID, task.Name, task.Weight)
	if err != nil {
		return fmt.Errorf("error adding data to database: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("no rows were affected")
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error commiting transaction: %w", err)
	}

	return nil
}

func (t *taskRepo) CompleteTask(ctx context.Context, task *models.Task) (error) {
	tx, err := t.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	res, err := tx.Exec(ctx, completeTaskQuery, true, task.ID)
	if err != nil {
		return fmt.Errorf("error adding data to database: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("no rows were affected")
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error commiting transaction: %w", err)
	}

	return nil
}
