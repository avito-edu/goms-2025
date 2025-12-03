//go:generate mockgen -package=mocks -source=$GOFILE -destination=mocks/$GOFILE
package service

type UserRepository interface {
	GetAllBy() ([]string, error)
}

type AutoRepository interface {
	Search(string2 string) ([]string, error)
}

type Service struct {
	uRepo UserRepository
	aRepo AutoRepository
}

func New(uRepo UserRepository, aRepo AutoRepository) *Service {
	return &Service{uRepo: uRepo, aRepo: aRepo}
}

func (s *Service) Check() ([]string, error) {
	all, err := s.uRepo.GetAllBy()
	if err != nil {
		return nil, err
	}

	got, err := s.aRepo.Search(all[0])
	if err != nil {
		return nil, err
	}

	return got, nil
}
