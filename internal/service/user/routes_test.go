package user

import (
	"GoForBeginner/internal/db/models"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleRegister_Success(t *testing.T) {
	// Мокируем userRepo
	userRepo := &mockUserStore{
		GetUserByEmailFunc: func(email string) (*models.User, error) {
			return nil, nil // Пользователь не существует
		},
		CreateUserFunc: func(user models.User) error {
			return nil // Успешное создание пользователя
		},
	}

	// Создаем обработчик
	handler := NewHandler(userRepo)

	// Создаем тестовый запрос
	payload := `{
		"first_name": "John",
		"last_name": "Doe",
		"email": "john.doe@example.com",
		"password": "password123",
		"nickname": "johndoe"
	}`
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	// Создаем ResponseRecorder
	rr := httptest.NewRecorder()

	// Вызываем обработчик
	handler.handleRegister(rr, req)

	// Проверяем статус код
	if rr.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, rr.Code)
	}

	// Проверяем тело ответа
	expectedResponse := "null" // Ожидаем "null" вместо пустой строки
	if strings.TrimSpace(rr.Body.String()) != expectedResponse {
		t.Errorf("expected response %s, got %s", expectedResponse, rr.Body.String())
	}
}

type mockUserStore struct {
	GetUserByEmailFunc func(email string) (*models.User, error)
	GetUserByIDFunc    func(id int) (*models.User, error)
	CreateUserFunc     func(user models.User) error
}

func (m mockUserStore) GetUserByEmail(email string) (*models.User, error) {
	if m.GetUserByEmailFunc != nil {
		return m.GetUserByEmailFunc(email)
	}
	return nil, nil
}

func (m mockUserStore) GetUserByID(id int) (*models.User, error) {
	if m.GetUserByIDFunc != nil {
		return m.GetUserByIDFunc(id)
	}
	return nil, nil
}

func (m mockUserStore) CreateUser(user models.User) error {
	if m.CreateUserFunc != nil {
		return m.CreateUserFunc(user)
	}
	return nil
}
