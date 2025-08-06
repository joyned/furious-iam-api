package repositories

import (
	"furious/iam-api/internal/configs"
	"furious/iam-api/internal/models"
	"furious/iam-api/pkg/utils"

	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(DB *gorm.DB) *RoleRepository {
	return &RoleRepository{
		DB: DB,
	}
}

func (r *RoleRepository) Persist(entity *models.Role) error {
	var result *gorm.DB
	if entity.ID != 0 {
		result = r.DB.Model(&entity).Omit("ID").Updates(entity)
	} else {
		result = r.DB.Create(entity)
	}

	return result.Error
}

func (r *RoleRepository) Search(params *models.Role, pagination utils.Pagination) (*utils.Pagination, error) {
	var roles []models.Role

	result := r.DB.Scopes(configs.Paginate(roles, &pagination, r.DB)).Where(&params).Find(&roles)

	pagination.Rows = roles

	return &pagination, result.Error
}

func (r *RoleRepository) Delete(id uint) error {
	result := r.DB.Delete(&models.Role{}, id)
	return result.Error
}
