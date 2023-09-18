package database

import (
	"github.com/labstack/echo/v4"

	"sample/pkg/domain/model"
	"sample/pkg/domain/repository"
	"sample/pkg/infrastructure/database"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (r *userRepository) List(c echo.Context, limit int, offset int) (*model.Users, error) {
	var users model.Users
	result := database.DB.Offset(offset).Limit(limit).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}
func (r *userRepository) FindById(c echo.Context, userID int) (*model.User, error) {
	var user model.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) Create(c echo.Context, user *model.User) (*model.User, error) {
	result := database.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
