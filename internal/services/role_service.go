package services

import (
	"errors"
	"furious/iam-api/internal/models"
	"furious/iam-api/internal/repositories"
	"furious/iam-api/pkg/utils"
)

type RoleService struct {
	repository *repositories.RoleRepository
}

func NewRoleService(repository *repositories.RoleRepository) *RoleService {
	return &RoleService{
		repository: repository,
	}
}

func (s *RoleService) Search(params *models.Role, pagination utils.Pagination) (*utils.Pagination, error) {
	return s.repository.Search(params, pagination)
}

func (s *RoleService) Persist(entity *models.Role) error {
	if entity.Name == "" {
		return errors.New("role name cannot be empty")
	}

	return s.repository.Persist(entity)
}

func (s *RoleService) Delete(id uint) error {
	if id == 0 {
		return errors.New("its not possible to delete a role with id 0")
	}

	return s.repository.Delete(id)
}
