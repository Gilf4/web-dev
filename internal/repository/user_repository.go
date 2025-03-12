package repository

import (
	"GoForBeginner/internal/db/models"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

type UserStoreRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(user models.User) error
}

type UserStore struct {
	db *pgxpool.Pool
}

func NewUserStore(pool *pgxpool.Pool) *UserStore {
	return &UserStore{
		db: pool,
	}
}

func (s *UserStore) GetUserByEmail(email string) (*models.User, error) {
	user := new(models.User)

	query := "SELECT id, first_name, last_name, email, password, nickname FROM users WHERE email = $1"

	err := s.db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Nickname,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // Пользователь не найден
		}
		return nil, fmt.Errorf("failed to query user by email: %w", err)
	}

	return user, nil
}

func (s *UserStore) GetUserByID(id int) (*models.User, error) {
	user := new(models.User)

	query := "SELECT * FROM users WHERE id = $1"

	rows, err := s.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query user by id: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(user)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		return user, nil
	}

	return nil, nil
}

func (s *UserStore) CreateUser(user models.User) error {
	query := "INSERT INTO users (first_name, last_name, email, password, nickname) VALUES ($1, $2, $3, $4, $5)"

	_, err := s.db.Exec(context.Background(), query, user.FirstName, user.LastName, user.Email, user.Password, user.Nickname)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
