package repositories

import (
	"errors"
	"furious/iam-api/internal/configs"
	"furious/iam-api/internal/models"
	"furious/iam-api/pkg/utils"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: DB,
	}
}

func (r *UserRepository) Persist(user *models.User) error {
	var result *gorm.DB
	if user.ID != 0 {
		result = r.DB.Model(user).Omit("ID").Updates(user)
	} else {
		result = r.DB.Create(user)
	}
	return result.Error
}

func (r *UserRepository) Search(params *models.User, pagination utils.Pagination) (*utils.Pagination, error) {
	var users []models.User

	result := r.DB.Scopes(configs.Paginate(users, &pagination, r.DB)).Where(&params).Find(&users)

	pagination.Rows = users

	return &pagination, result.Error
}

func (r *UserRepository) Delete(id uint) error {
	if id == 0 {
		return errors.New("it's not possible to delete a user with id 0")
	}

	result := r.DB.Delete(&models.User{}, id)

	return result.Error
}
