package repository

import (
	"github.com/labstack/echo/v4"

	"sample/pkg/domain/model"
)

// UserRepository NOTE: DIのためにここではインターフェイスのみ
type UserRepository interface {
	List(c echo.Context, limit int, offset int) (*model.Users, error)
	FindById(c echo.Context, userID int) (*model.User, error)
	Create(c echo.Context, user *model.User) (*model.User, error)
}
