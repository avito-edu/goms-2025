// user_service.go
package user

import "context"

type User struct {
	ID    int
	Name  string
	Email string
}

// UserRepository - интерфейс, который мы будем мокать
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int) (*User, error)
	Delete(ctx context.Context, id int) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, name, email string) (*User, error) {
	user := &User{Name: name, Email: email}
	err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUser(ctx context.Context, id int) (*User, error) {
	return s.repo.GetByID(ctx, id)
}
