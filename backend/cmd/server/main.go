package main

import (
	"net/http"
	
	"backend/pkg/schema"
	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ApiServer implements the generated ServerInterface
type ApiServer struct{}

// Users endpoints - minimal implementation
func (a *ApiServer) RegisterUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) LoginUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) LogoutUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) GetCurrentUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) UpdateUserProfile(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) GetPublicUserInfo(ctx echo.Context, userId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) BlockUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) GetBlockedUsers(ctx echo.Context, params schema.GetBlockedUsersParams) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) UnblockUser(ctx echo.Context, userId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) AddNgWord(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) GetNgWords(ctx echo.Context, params schema.GetNgWordsParams) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) DeleteNgWord(ctx echo.Context, ngWordId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

// Questions endpoints - minimal implementation
func (a *ApiServer) CreateQuestion(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) GetReceivedQuestions(ctx echo.Context, params schema.GetReceivedQuestionsParams) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) GetSentQuestions(ctx echo.Context, params schema.GetSentQuestionsParams) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) GetQuestionDetail(ctx echo.Context, questionId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) UpdateQuestion(ctx echo.Context, questionId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) DeleteQuestion(ctx echo.Context, questionId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) MarkQuestionAsRead(ctx echo.Context, questionId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) CreateAnswer(ctx echo.Context, questionId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) GetPublicQAs(ctx echo.Context, userId string, params schema.GetPublicQAsParams) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) ReportContent(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

// Files endpoints - minimal implementation
func (a *ApiServer) UploadImage(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) GetImage(ctx echo.Context, imageId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func (a *ApiServer) DeleteImage(ctx echo.Context, imageId string) error {
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"message": "Not implemented yet"})
}

func main() {
	e := echo.New()
	
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Health check endpoint
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "質問箱API Server is running"})
	})

	// Register API routes with the generated handlers
	apiServer := &ApiServer{}
	schema.RegisterHandlers(e, apiServer)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}