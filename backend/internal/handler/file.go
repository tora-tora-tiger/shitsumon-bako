package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type FileHandler struct{}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

func (h *FileHandler) UploadImage(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *FileHandler) GetImage(ctx echo.Context, imageId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *FileHandler) DeleteImage(ctx echo.Context, imageId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}