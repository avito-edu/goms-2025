package mocks_stubs

import "github.com/stretchr/testify/mock"

type User struct{}

type UserRepositoryMock struct {
	mock.Mock
}

func (u *UserRepositoryMock) GetUser(id int) (*User, error) {
	args := u.Called(id)                      // Фиксируем факт вызова
	return args.Get(0).(*User), args.Error(1) // Возвращаем то, что настроили в тесте
}
