package web

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"sample/pkg/domain/model"
	database "sample/pkg/infrastructure/database/persistence"

	"net/http"

	"sample/pkg/usecase"
)

type UserHandler interface {
	ListUsers(c echo.Context) error
	CreateUser(echo.Context) error
	ShowUser(echo.Context) error
}

type userHandler struct {
	usecase.UserUseCase
}

func NewUserHandler() UserHandler {
	repo := database.NewUserRepository()
	uc := usecase.NewUserUseCase(repo)
	return &userHandler{uc}
}

func (h *userHandler) CreateUser(c echo.Context) error {
	uc := h.UserUseCase
	user := &model.User{}
	var err error
	if err = c.Bind(user); err != nil {
		return err
	}
	user, err = uc.CreateUser(c, user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

type listUserRequestDto struct {
	Pagination
}

func (h *userHandler) ListUsers(c echo.Context) error {
	dto := &listUserRequestDto{
		Pagination: NewPagination(),
	}
	var err error
	if err = c.Bind(dto); err != nil {
		return err
	}

	users, err := h.UserUseCase.GetUserList(c, dto.Limit, dto.Offset)
	if err != nil {
		return err
	}

	println("users", len(*users))

	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) ShowUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	user, err := h.UserUseCase.GetUserDetails(c, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
