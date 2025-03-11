package repository

import (
	"GoForBeginner/internal/db/models"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

type UserStoreRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) error
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

	query := "SELECT * FROM users WHERE email = $1"

	rows, err := s.db.Query(context.Background(), query, email)
	if err != nil {
		return nil, fmt.Errorf("failed to query user by email: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(user)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
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

	for rows.Next() {
		err := rows.Scan(user)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *UserStore) CreateUser(user *models.User) error {
	query := "INSERT INTO users (firstName, lastName, email, password) VALUES ($1, $2, $3, $4)"

	_, err := s.db.Exec(context.Background(), query, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
