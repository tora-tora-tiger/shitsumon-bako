package handler

import (
	"fmt"
	"net/http"

	"backend/pkg/schema"

	"backend/internal/db"

	"github.com/labstack/echo/v4"
)

var DB = db.DB

type QuestionHandler struct{}

func NewQuestionHandler() *QuestionHandler {
	return &QuestionHandler{}
}

func (h *QuestionHandler) CreateQuestion(ctx echo.Context) error {
	req := schema.QuestionCreateRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &schema.ErrorDetails{
			Message: "Invalid request payload",
		})
	}

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
	// debug
	if params.Limit != nil {
		fmt.Printf("limit: %v\n", *params.Limit)
	} else {
		fmt.Println("limit: nil")
	}
	if params.Page != nil {
		fmt.Printf("page: %v\n", *params.Page)
	} else {
		fmt.Println("page: nil")
	}
	if params.Status != nil {
		fmt.Printf("status: %v\n", *params.Status)
	} else {
		fmt.Println("status: nil")
	}
	if params.SortBy != nil {
		fmt.Printf("sortBy: %v\n", *params.SortBy)
	} else {
		fmt.Println("sortBy: nil")
	}
	if params.SortOrder != nil {
		fmt.Printf("sortOrder: %v\n", *params.SortOrder)
	} else {
		fmt.Println("sortOrder: nil")
	}
	p := schema.GetReceivedQuestionsParams(params)
	fmt.Printf("params: %+v\n", p)
	// var questions []db.Question

	// クエリを作る
	// var order string
	// if params.SortBy != nil {
	// 	field := "createdAt"
		
	// 	if params.SortBy != nil && *params.SortBy != schema.GetReceivedQuestionsParamsSortBy {
	// 		field = *params.SortBy
	// 	}
	// 	orderDir := ""
	// 	if params.SortOrder != nil {
	// 		orderDir = *params.SortOrder
	// 	}
	// 	order = fmt.Sprintf("%s %s", field, orderDir)
	// }
	// var query string
	// if params.Status != nil {
	// 	query = fmt.Sprintf("status = '%s'", *params.Status)
	// }

	// error := DB.Where(query).Order(order).Find(&questions)

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
