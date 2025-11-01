package handler

import (
	"fmt"
	"net/http"

	"backend/pkg/schema"

	"backend/internal/db"

	"github.com/labstack/echo/v4"
)

type QuestionHandler struct{}

func NewQuestionHandler() *QuestionHandler {
	return &QuestionHandler{}
}

func (h *QuestionHandler) CreateQuestion(ctx echo.Context) error {
	DB := db.DB

	req := schema.QuestionCreateRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &schema.ErrorDetails{
			Message: "Invalid request payload",
		})
	}

	fmt.Println(req)

	q := db.Question{
		Content:     req.Content,
		RecipientId: req.RecipientId,
		// 省略したいフィールドは設定しない（ポインタなら nil のまま）
	}

	if req.AttachedImageIdList != nil && len(*req.AttachedImageIdList) > 0 {
		list := make([]db.ImageFile, 0, len(*req.AttachedImageIdList))
		for _, id := range *req.AttachedImageIdList {
			list = append(list, db.ImageFile{Id: id})
		}
		q.AttachedImageList = &list
	}

	if err := DB.Create(&q).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, &schema.ErrorDetails{
			Message: "Failed to create question",
			Details: err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, q)
}

func (h *QuestionHandler) GetReceivedQuestions(ctx echo.Context, params schema.GetReceivedQuestionsParams) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *QuestionHandler) GetSentQuestions(ctx echo.Context, params schema.GetSentQuestionsParams) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *QuestionHandler) GetQuestionDetail(ctx echo.Context, questionId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *QuestionHandler) UpdateQuestion(ctx echo.Context, questionId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *QuestionHandler) DeleteQuestion(ctx echo.Context, questionId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *QuestionHandler) MarkQuestionAsRead(ctx echo.Context, questionId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *QuestionHandler) CreateAnswer(ctx echo.Context, questionId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *QuestionHandler) GetPublicQAs(ctx echo.Context, userId string, params schema.GetPublicQAsParams) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (h *QuestionHandler) ReportContent(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}
