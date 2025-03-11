package db

import (
	"GoForBeginner/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg *config.DBConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.USERNAME,
		cfg.PASSWORD,
		cfg.HOST,
		cfg.PORT,
		cfg.DbName,
	)

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("could not create connection pool: %w", err)
	}

	if err := db.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("could not ping database: %w", err)
	}

	return db, nil
}

//TODO close conn pool
