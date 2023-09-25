package db

import (
	"context"
	"fmt"
	"homework-5/configs"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDB(ctx context.Context) (*Database, error) {
	dsn := generateDsn()
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return newDatabase(pool), nil
}

func generateDsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", configs.Host, configs.Port, configs.User, configs.Password, configs.DBname)
}
