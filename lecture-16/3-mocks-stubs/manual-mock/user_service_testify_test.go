// user_service_testify_test.go
package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestifyUserRepository - мок с использованием testify
type TestifyUserRepository struct {
	mock.Mock
}

func (m *TestifyUserRepository) Create(ctx context.Context, user *User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *TestifyUserRepository) GetByID(ctx context.Context, id int) (*User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*User), args.Error(1)
}

func (m *TestifyUserRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Тесты с Testify Mock
func TestUserService_TestifyMock(t *testing.T) {
	t.Run("successful user registration", func(t *testing.T) {
		mockRepo := new(TestifyUserRepository)
		service := NewUserService(mockRepo)

		// Настраиваем ожидания
		mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*user.User")).
			Return(nil).
			Run(func(args mock.Arguments) {
				user := args.Get(1).(*User)
				user.ID = 1 // Симулируем присвоение ID
			})

		user, err := service.RegisterUser(context.Background(), "John", "john@test.com")

		assert.NoError(t, err)
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "John", user.Name)

		mockRepo.AssertExpectations(t)
	})

	t.Run("create user error", func(t *testing.T) {
		mockRepo := new(TestifyUserRepository)
		service := NewUserService(mockRepo)

		mockRepo.On("Create",
			mock.Anything, mock.AnythingOfType("*user.User")).
			Return(errors.New("database error"))

		user, err := service.
			RegisterUser(context.Background(), "John", "john@test.com")

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "database")

		mockRepo.AssertExpectations(t)
	})

	t.Run("get user success", func(t *testing.T) {
		mockRepo := new(TestifyUserRepository)
		service := NewUserService(mockRepo)

		expectedUser := &User{ID: 1, Name: "Jane", Email: "jane@test.com"}

		// Настраиваем ожидания для GetByID
		mockRepo.On("GetByID", mock.Anything, 1).
			Return(expectedUser, nil)

		foundUser, err := service.GetUser(context.Background(), 1)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, foundUser)

		mockRepo.AssertExpectations(t)
		mockRepo.AssertCalled(t, "GetByID", mock.Anything, 1)
	})

	t.Run("get user not found", func(t *testing.T) {
		mockRepo := new(TestifyUserRepository)
		service := NewUserService(mockRepo)

		mockRepo.On("GetByID", mock.Anything, 999).
			Return((*User)(nil), errors.New("user not found"))

		user, err := service.GetUser(context.Background(), 999)

		assert.Error(t, err)
		assert.Nil(t, user)

		mockRepo.AssertExpectations(t)
	})
}
