package repository

type UserRepository struct{}

func New() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByID(id string) (string, error) {
	return "John Doe", nil
}
