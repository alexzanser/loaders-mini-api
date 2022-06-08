package postgres

import (
	"context"
	"runtime"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPool(connstr string) (*pgxpool.Pool, error) {
	connConf, err := pgxpool.ParseConfig(connstr)
	if err != nil {
		return nil, err
	}

	connConf.MaxConns = int32(runtime.NumCPU())
	pool, err := pgxpool.ConnectConfig(context.Background(), connConf)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
