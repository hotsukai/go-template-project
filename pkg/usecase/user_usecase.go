package usecase

import (
	"github.com/labstack/echo/v4"

	"sample/pkg/domain/model"
	"sample/pkg/domain/repository"
)

type UserUseCase interface {
	GetUserList(c echo.Context, limit int, offset int) (*model.Users, error)
	GetUserDetails(c echo.Context, id int) (*model.User, error)
	CreateUser(c echo.Context, user *model.User) (*model.User, error)
}

type userUseCase struct {
	repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (u *userUseCase) GetUserList(c echo.Context, limit int, offset int) (*model.Users, error) {
	repo := u.UserRepository
	users, err := repo.List(c, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUseCase) GetUserDetails(c echo.Context, id int) (*model.User, error) {
	repo := u.UserRepository
	user, err := repo.FindById(c, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUseCase) CreateUser(c echo.Context, user *model.User) (*model.User, error) {
	repo := u.UserRepository
	user, err := repo.Create(c, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
