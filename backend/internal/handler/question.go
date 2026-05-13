package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gen"

	"backend/internal/database"
	"backend/internal/database/model"
	"backend/internal/database/query"
	"backend/pkg/schema"
)

var DB = database.DB

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

	q := model.Question{
		RecipientId: req.RecipientId,
		// senderは認証情報から取る
		Content:     req.Content,
		// 省略したいフィールドは設定しない（ポインタなら nil のまま）
	}

	if req.AttachedImageIdList != nil && len(*req.AttachedImageIdList) > 0 {
		list := make([]model.ImageFile, 0, len(*req.AttachedImageIdList))
		for _, id := range *req.AttachedImageIdList {
			list = append(list, model.ImageFile{Id: id})
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
	q := query.Question.WithContext(ctx.Request().Context())
	// q := DB.WithContext(ctx.Request().Context())
	// クエリを作る
	// ステータスフィルター
	if params.Status != nil {
		q = q.Where(query.Question.Status.Eq(string(*params.Status)))
	}

	// 添付ファイル
	q = q.Preload(query.Question.AttachedImageList)

	// 回答
	q = q.Preload(query.Question.Answer)

	// ソート
	if params.SortBy != nil {
		field := schema.GetReceivedQuestionsParamsSortByCreatedAt
		if params.SortBy != nil {
			field = *params.SortBy
		}

		orderDir := schema.Desc
		if params.SortOrder != nil {
			orderDir = *params.SortOrder
		}

		// colName := q.NamingStrategy.ColumnName("questions", string(field))
		colName := query.Question.UnderlyingDB().NamingStrategy.ColumnName(query.Question.TableName(), string(field))
		if col, ok := query.Question.GetFieldByName(colName); ok {
			if orderDir == schema.Asc {
				q = q.Order(col.Asc())
			} else {
				q = q.Order(col.Desc())
			}
		}
		// q = q.Order(colName + " " + string(orderDir))
	}

	// ページネーション
	if params.Limit != nil && params.Page != nil {
		q = q.Limit(int(*params.Limit)).Offset(int(*params.Page))
	}

	var questions []*model.Question
	// result := q.Find(&questions)
	questions, err := q.Find()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &schema.ErrorDetails{
			Message: "Failed to retrieve questions",
			Details: err,
		})
	}

	return ctx.JSON(http.StatusOK, questions)
}

func (h *QuestionHandler) GetSentQuestions(ctx echo.Context, params schema.GetSentQuestionsParams) error {
	q := query.Question.WithContext(ctx.Request().Context())

	if params.HasAnswer != nil {
		ans := query.Answer
		qst := query.Question
		subq := query.Answer.WithContext(ctx.Request().Context()).Where(ans.QuestionId.EqCol(qst.Id))

		if *params.HasAnswer {
			// SELECT * FROM questions
			// WHERE EXISTS (
			//  SELECT * FROM answers
			//  WHERE answers.question_id = questions.id
			// )
			q = q.Where(gen.Exists(subq))
		} else {
			q = q.Not(gen.Exists(subq))
		}
		q.Preload(query.Question.Answer)
	}

	q.Preload(query.Question.AttachedImageList)

	// ソート
	if params.SortBy != nil {
		field := schema.CreatedAt
		if params.SortBy != nil {
			field = *params.SortBy
		}

		orderDir := schema.Desc
		if params.SortOrder != nil {
			orderDir = *params.SortOrder
		}

		colName := query.Question.UnderlyingDB().NamingStrategy.ColumnName(query.Question.TableName(), string(field))
		if col, ok := query.Question.GetFieldByName(colName); ok {
			if orderDir == schema.Asc {
				q = q.Order(col.Asc())
			} else {
				q = q.Order(col.Desc())
			}
		}
	}


	// ページネーション
	if params.Limit != nil && params.Page != nil {
		q = q.Limit(int(*params.Limit)).Offset(int(*params.Page))
	}
	
	var questions []*model.Question
	// result := q.Find(&questions)
	questions, err := q.Find()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &schema.ErrorDetails{
			Message: "Failed to retrieve questions",
			Details: err,
		})
	}

	return ctx.JSON(http.StatusOK, questions)
}

func (h *QuestionHandler) GetQuestionDetail(ctx echo.Context, questionId string) error {
	q := query.Question.WithContext(ctx.Request().Context())
	q = q.Where(query.Question.Id.Eq(questionId))

	q.Preload(query.Question.AttachedImageList)
	q.Preload(query.Question.Answer)

	question, err := q.Find()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &schema.ErrorDetails{
			Message: "Failed to retrieve questions",
			Details: err,
		})
	}

	if len(question) == 0 {
		return ctx.JSON(http.StatusNotFound, &schema.ErrorDetails{
			Message: "Question not found",
		})
	}

	return ctx.JSON(http.StatusOK, question[0])
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
