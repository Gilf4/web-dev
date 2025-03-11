package migrations

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

func CreateTables(ctx context.Context, pool *pgxpool.Pool) error {
	if err := createUsersTable(ctx, pool); err != nil {
		return fmt.Errorf("не удалось создать таблицу users: %w", err)
	}
	if err := createBooksTable(ctx, pool); err != nil {
		return fmt.Errorf("не удалось создать таблицу books: %w", err)
	}
	return nil
}

func createUsersTable(ctx context.Context, db *pgxpool.Pool) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("could not create users table: %w", err)
	}
	return nil
}

func createBooksTable(ctx context.Context, db *pgxpool.Pool) error {
	query := `
	CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		authors VARCHAR(255) UNIQUE NOT NULL,
		publication_year INTEGER NOT NULL,
		isbn VARCHAR(20) NOT NULL,
		cover_image_url VARCHAR(255) NOT NULL
	);
	`
	_, err := db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("could not create books table: %w", err)
	}
	return nil
}
