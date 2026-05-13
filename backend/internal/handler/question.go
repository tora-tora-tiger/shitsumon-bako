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
		Content: req.Content,
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
	q := query.Question.WithContext(ctx.Request().Context())
	q = q.Where(query.Question.Id.Eq(questionId))

	questions, err := q.Find()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &schema.ErrorDetails{
			Message: "Failed to retrieve question",
			Details: err,
		})
	}

	if len(questions) == 0 {
		return ctx.JSON(http.StatusNotFound, &schema.ErrorDetails{
			Message: "Question not found",
		})
	}

	if _, err := q.Delete(questions[0]); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &schema.ErrorDetails{
			Message: "Failed to delete question",
			Details: err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *QuestionHandler) MarkQuestionAsRead(ctx echo.Context, questionId string) error {
	q := query.Question.WithContext(ctx.Request().Context())
	q = q.Where(query.Question.Id.Eq(questionId))
	q = q.Preload(query.Question.AttachedImageList)

	questions, err := q.Find()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &schema.ErrorDetails{
			Message: "Failed to retrieve question",
			Details: err,
		})
	}

	if len(questions) == 0 {
		return ctx.JSON(http.StatusNotFound, &schema.ErrorDetails{
			Message: "Question not found",
		})
	}

	question := questions[0]
	if question.Status != string(schema.Answered) {
		if _, err := q.Update(query.Question.Status, string(schema.Read)); err != nil {
			return ctx.JSON(http.StatusInternalServerError, &schema.ErrorDetails{
				Message: "Failed to mark question as read",
				Details: err,
			})
		}
		question.Status = string(schema.Read)
	}

	return ctx.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) CreateAnswer(ctx echo.Context, questionId string) error {
	req := schema.AnswerCreateRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &schema.ErrorDetails{
			Message: "Invalid request payload",
		})
	}

	q := query.Question.WithContext(ctx.Request().Context())
	q = q.Where(query.Question.Id.Eq(questionId))
	q = q.Preload(query.Question.Answer)

	questions, err := q.Find()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &schema.ErrorDetails{
			Message: "Failed to retrieve question",
			Details: err,
		})
	}

	if len(questions) == 0 {
		return ctx.JSON(http.StatusNotFound, &schema.ErrorDetails{
			Message: "Question not found",
		})
	}

	question := questions[0]
	if question.Answer != nil {
		return ctx.JSON(http.StatusBadRequest, &schema.ErrorDetails{
			Message: "Question already has an answer",
		})
	}

	answer := model.Answer{
		QuestionId: questionId,
		Content:    req.Content,
		IsPublic:   req.IsPublic,
	}

	if req.AttachedImageIdList != nil && len(*req.AttachedImageIdList) > 0 {
		list := make([]model.ImageFile, 0, len(*req.AttachedImageIdList))
		for _, id := range *req.AttachedImageIdList {
			list = append(list, model.ImageFile{Id: id, QuestionId: questionId})
		}
		answer.AttachedImageList = &list
	}

	if err := query.Q.Transaction(func(tx *query.Query) error {
		txCtx := tx.WithContext(ctx.Request().Context())
		if err := txCtx.Answer.Create(&answer); err != nil {
			return err
		}
		_, err := txCtx.Question.
			Where(tx.Question.Id.Eq(questionId)).
			Update(tx.Question.Status, string(schema.Answered))
		return err
	}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &schema.ErrorDetails{
			Message: "Failed to create answer",
			Details: err,
		})
	}

	return ctx.JSON(http.StatusCreated, answer)
}

func (h *QuestionHandler) GetPublicQAs(ctx echo.Context, userId string, params schema.GetPublicQAsParams) error {
	q := query.Question.WithContext(ctx.Request().Context())
	ans := query.Answer
	qst := query.Question
	subq := query.Answer.WithContext(ctx.Request().Context()).
		Where(ans.QuestionId.EqCol(qst.Id), ans.IsPublic.Is(true))

	q = q.Where(query.Question.RecipientId.Eq(userId), gen.Exists(subq))
	q = q.Preload(query.Question.AttachedImageList)
	q = q.Preload(query.Question.Answer)
	q = q.Preload(query.Question.Answer.AttachedImageList)

	if params.SortBy != nil {
		field := schema.GetPublicQAsParamsSortByCreatedAt
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

	if params.Limit != nil && params.Page != nil {
		q = q.Limit(int(*params.Limit)).Offset(int(*params.Page))
	}

	questions, err := q.Find()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &schema.ErrorDetails{
			Message: "Failed to retrieve public questions",
			Details: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, questions)
}

func (h *QuestionHandler) ReportContent(ctx echo.Context) error {
	req := schema.ReportRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &schema.ErrorDetails{
			Message: "Invalid request payload",
		})
	}

	if req.TargetType != schema.ReportRequestTargetTypeQuestion && req.TargetType != schema.ReportRequestTargetTypeUser {
		return ctx.JSON(http.StatusBadRequest, &schema.ErrorDetails{
			Message: "Invalid target type",
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{})
}
