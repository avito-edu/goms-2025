// user_service_manual_test.go
package user

import (
	"context"
	"errors"
	"testing"
)

// ManualUserRepository - ручная реализация мока
type ManualUserRepository struct {
	users      map[int]*User
	nextID     int
	createErr  error
	getByIDErr error
	deleteErr  error

	// Для проверки вызовов
	createCalls  []*User
	getByIDCalls []int
	deleteCalls  []int
}

func NewManualUserRepository() *ManualUserRepository {
	return &ManualUserRepository{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

func (m *ManualUserRepository) Create(ctx context.Context, user *User) error {
	m.createCalls = append(m.createCalls, user)

	if m.createErr != nil {
		return m.createErr
	}

	user.ID = m.nextID
	m.users[user.ID] = user
	m.nextID++
	return nil
}

func (m *ManualUserRepository) GetByID(ctx context.Context, id int) (*User, error) {
	m.getByIDCalls = append(m.getByIDCalls, id)

	if m.getByIDErr != nil {
		return nil, m.getByIDErr
	}

	user, exists := m.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	// Возвращаем копию чтобы избежать модификации
	return &User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (m *ManualUserRepository) Delete(ctx context.Context, id int) error {
	m.deleteCalls = append(m.deleteCalls, id)

	if m.deleteErr != nil {
		return m.deleteErr
	}

	delete(m.users, id)
	return nil
}

// Вспомогательные методы для тестов
func (m *ManualUserRepository) WithCreateError(err error) *ManualUserRepository {
	m.createErr = err
	return m
}

func (m *ManualUserRepository) WithGetByIDError(err error) *ManualUserRepository {
	m.getByIDErr = err
	return m
}

// Тесты с Manual Mock
func TestUserService_ManualMock(t *testing.T) {
	t.Run("successful user registration", func(t *testing.T) {
		mockRepo := NewManualUserRepository()
		service := NewUserService(mockRepo)

		user, err := service.RegisterUser(context.Background(), "John", "john@test.com")

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if user.ID != 1 {
			t.Errorf("Expected user ID 1, got %d", user.ID)
		}
		if user.Name != "John" {
			t.Errorf("Expected user name John, got %s", user.Name)
		}

		// Проверяем вызовы
		if len(mockRepo.createCalls) != 1 {
			t.Errorf("Expected 1 call to Create, got %d", len(mockRepo.createCalls))
		}
	})

	t.Run("create user error", func(t *testing.T) {
		mockRepo := NewManualUserRepository().
			WithCreateError(errors.New("database error"))
		service := NewUserService(mockRepo)

		user, err := service.RegisterUser(context.Background(), "John", "john@test.com")

		if err == nil {
			t.Error("Expected error, got nil")
		}
		if user != nil {
			t.Error("Expected nil user, got user")
		}
	})

	t.Run("get user success", func(t *testing.T) {
		mockRepo := NewManualUserRepository()
		service := NewUserService(mockRepo)

		// Сначала создаем пользователя
		createdUser, _ := service.RegisterUser(context.Background(), "Jane", "jane@test.com")

		// Затем получаем его
		foundUser, err := service.GetUser(context.Background(), createdUser.ID)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if foundUser.Name != "Jane" {
			t.Errorf("Expected user name Jane, got %s", foundUser.Name)
		}

		// Проверяем вызовы
		if len(mockRepo.getByIDCalls) != 1 {
			t.Errorf("Expected 1 call to GetByID, got %d", len(mockRepo.getByIDCalls))
		}
		if mockRepo.getByIDCalls[0] != 1 {
			t.Errorf("Expected GetByID call with ID 1, got %d", mockRepo.getByIDCalls[0])
		}
	})
}
