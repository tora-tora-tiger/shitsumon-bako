package handler

import (
	"backend/pkg/schema"
)

// APIHandler combines all handlers to implement ServerInterface
type APIHandler struct {
	*UserHandler
	*QuestionHandler
	*FileHandler
}

// Ensure APIHandler implements schema.ServerInterface
var _ schema.ServerInterface = APIHandler{}

func NewAPIHandler() *APIHandler {
	return &APIHandler{
		UserHandler:     NewUserHandler(),
		QuestionHandler: NewQuestionHandler(),
		FileHandler:     NewFileHandler(),
	}
}