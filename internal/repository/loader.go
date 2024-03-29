package repository

import (
	"context"
	"fmt"
	"loaders/internal/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type loaderRepo struct {
	pool *pgxpool.Pool
}

func newLoaderRepo(pool *pgxpool.Pool) *loaderRepo {
	return &loaderRepo{
		pool: pool,
	}
}

const (
	getLoaderQuery 			= "SELECT id, username, passwd_hash, max_weight, drunk, fatigue, salary, balance FROM loaders WHERE username=$1"
	getCompletedTasksQuery  = `SELECT tasks.id, tasks.name, tasks.weight FROM tasks, loaders  
									WHERE loaders.id = $1 and tasks.id = ANY (loaders.completed_tasks) and completed=true`;
	getLoadersListQuery		= "SELECT  id, username, max_weight, salary FROM loaders;"
	getLoadersFullQuery		= "SELECT id, username, passwd_hash, max_weight, drunk, fatigue, salary, balance FROM loaders;"
	loadersUpdateQuery		= "UPDATE loaders SET  max_weight = $1, fatigue=$2, balance=$3, completed_tasks=completed_tasks || $4 WHERE id=$5;"
)

func (l *loaderRepo) GetLoader(ctx context.Context, username, passwd string) (*models.Loader, error) {
	tx, err := l.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	
	var ld models.Loader
	var id	int64

	row := l.pool.QueryRow(ctx, getLoaderQuery, username)
	if err := row.Scan(&id, &ld.Username, &ld.PasswdHash, &ld.MaxWeight, 
						&ld.Drunk, &ld.Fatigue, &ld.Salary, &ld.Balance); err != nil {
		return nil, fmt.Errorf("error receiving data from database: %w", err)
	}

	rows, err := l.pool.Query(ctx, getCompletedTasksQuery, id)
	if err != nil {
		return nil, fmt.Errorf("error receiving data from database: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.ID, &t.Name, &t.Weight)
		if err != nil {
			return nil, fmt.Errorf("error receiving data from database: %w", err)
		}
	   ld.CompletedTasks = append(ld.CompletedTasks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error receiving data from database: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("error commiting transaction: %w", err)
	}

	passwdHash, ok := ctx.Value("PasswdHash").(string)
	if ok && passwdHash == "" {
		if err := bcrypt.CompareHashAndPassword([]byte(ld.PasswdHash), []byte(passwd)); err != nil {
			return nil, fmt.Errorf("wrong username or password :%v", err)
		}
	}
	return &ld, nil
}

func (l *loaderRepo) GetLoadersList(ctx context.Context) ([]models.Loader, error) {
	tx, err := l.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	
	var ld []models.Loader

	rows, err := l.pool.Query(ctx, getLoadersListQuery)
	if err != nil {
		return nil, fmt.Errorf("error receiving data from database: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var l models.Loader
		err := rows.Scan(&l.ID, &l.Username, &l.MaxWeight, &l.Salary)
		if err != nil {
			return nil, fmt.Errorf("error receiving data from database: %w", err)
		}
	   ld = append(ld, l)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error receiving data from database: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("error commiting transaction: %w", err)
	}

	return ld, nil
}


func (l *loaderRepo) GetLoadersFull(ctx context.Context) ([]models.Loader, error) {
	tx, err := l.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	
	var ld []models.Loader

	rows, err := l.pool.Query(ctx, getLoadersFullQuery)
	if err != nil {
		return nil, fmt.Errorf("error receiving data from database: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var l models.Loader
		err := rows.Scan(&l.ID, &l.Username, &l.PasswdHash, &l.MaxWeight, &l.Drunk, &l.Fatigue, &l.Salary, &l.Balance)
		if err != nil {
			return nil, fmt.Errorf("error receiving data from database: %w", err)
		}
	   ld = append(ld, l)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error receiving data from database: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("error commiting transaction: %w", err)
	}

	return ld, nil
}

func (l *loaderRepo) UpdateLoader(ctx context.Context, ld *models.Loader) (error) {
	tx, err := l.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	completedID := make([]int64, 0)
	for _, val := range ld.CompletedTasks {
		completedID = append(completedID, val.ID)
	}
	res, err := tx.Exec(ctx, loadersUpdateQuery, ld.MaxWeight, ld.Fatigue, ld.Balance, completedID, ld.ID)
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
