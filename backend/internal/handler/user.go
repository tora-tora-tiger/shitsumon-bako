package handler

import (
	"net/http"

	"backend/pkg/schema"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) RegisterUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *UserHandler) LoginUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *UserHandler) LogoutUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *UserHandler) GetCurrentUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *UserHandler) UpdateUserProfile(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *UserHandler) GetPublicUserInfo(ctx echo.Context, userId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *UserHandler) BlockUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *UserHandler) GetBlockedUsers(ctx echo.Context, params schema.GetBlockedUsersParams) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *UserHandler) UnblockUser(ctx echo.Context, userId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *UserHandler) AddNgWord(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *UserHandler) GetNgWords(ctx echo.Context, params schema.GetNgWordsParams) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *UserHandler) DeleteNgWord(ctx echo.Context, ngWordId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}