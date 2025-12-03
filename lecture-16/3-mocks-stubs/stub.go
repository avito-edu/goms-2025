package mocks_stubs

type UserRepositoryStub struct{}

func (u *UserRepositoryStub) GetUser(id int) (*User, error) {
	return &User{ID: 1, Name: "Test User"}, nil // Всегда возвращает одного и того же пользователя
}
