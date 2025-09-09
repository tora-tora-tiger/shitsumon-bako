package handler

import (
	"net/http"

	"backend/pkg/schema"

	"github.com/labstack/echo/v4"
)

type QuestionHandler struct{}

func NewQuestionHandler() *QuestionHandler {
	return &QuestionHandler{}
}

func (h *QuestionHandler) CreateQuestion(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
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