package service

import (
	"ITMO-students/lecture-8/myapp/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func New(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUser(id string) (string, error) {
	return s.repo.FindByID(id)
}
