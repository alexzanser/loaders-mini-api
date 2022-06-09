package repository

import (
	"context"
	"fmt"
	"loaders/internal/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type customerRepo struct {
	pool *pgxpool.Pool
}

func newCustomerRepo(pool *pgxpool.Pool) *customerRepo {
	return &customerRepo{
		pool: pool,
	}
}

const (
	getCustomerQuery 		= "SELECT id, username, passwd_hash, balance FROM customers WHERE username=$1;"
	getCustomersListQuery 	= "SELECT id FROM CUSTOMERS;"
	getTasksQuery    		= "SELECT id, name, weight FROM tasks WHERE customer_id=$1 and completed=false;"
	customersUpdateQuery	= "UPDATE customers SET balance=$1 WHERE id=$2;"
)


func (c *customerRepo) GetCustomer(ctx context.Context, username, passwd string) (*models.Customer, error) {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	var ct models.Customer

	row := c.pool.QueryRow(ctx, getCustomerQuery, username)
	if err := row.Scan(&ct.ID, &ct.Username, &ct.PasswdHash, &ct.Balance); err != nil {
		return nil, fmt.Errorf("error receiving data from database: %w", err)
	}

	rows, err := c.pool.Query(ctx, getTasksQuery, ct.ID)
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
	   ct.Tasks = append(ct.Tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error receiving data from database: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("error commiting transaction: %w", err)
	}

	passwdHash, ok := ctx.Value("PasswdHash").(string)
	if ok && passwdHash == "" {
		if err := bcrypt.CompareHashAndPassword([]byte(ct.PasswdHash), []byte(passwd)); err != nil {
			return nil, fmt.Errorf("wrong username or password :%v", err)
		}
	}
	return &ct, nil
}	

func (c *customerRepo) GetCustomersList(ctx context.Context) ([]int64, error) {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	idList := make([]int64, 0)
	rows, err := c.pool.Query(ctx, getCustomersListQuery)
	if err != nil {
		return nil, fmt.Errorf("error receiving data from database: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("error receiving data from database: %w", err)
		}
	   idList = append(idList, id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error receiving data from database: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("error commiting transaction: %w", err)
	}

	return idList, nil
}

func (t *customerRepo) UpdateCustomer(ctx context.Context, taskCost int, ct *models.Customer) (error) {
	tx, err := t.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	res, err := tx.Exec(ctx, customersUpdateQuery, taskCost, ct.ID)
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