package services

import (
	"errors"
	"furious/iam-api/internal/models"
	"furious/iam-api/internal/repositories"
	"furious/iam-api/pkg/utils"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) Persist(entity *models.User) error {
	if entity.Email == "" {
		return errors.New("email cannot be empty")
	}

	if entity.Username == "" {
		return errors.New("username cannot be empty")
	}

	if entity.FirstName == "" {
		return errors.New("first name cannot be empty")
	}

	if entity.LastName == "" {
		return errors.New("last name cannot be empty")
	}

	return s.userRepository.Persist(entity)
}

func (s *UserService) Search(params *models.User, pagination utils.Pagination) (*utils.Pagination, error) {
	users, err := s.userRepository.Search(params, pagination)
	return users, err
}

func (s *UserService) Delete(id uint) error {
	return s.userRepository.Delete(id)
}
