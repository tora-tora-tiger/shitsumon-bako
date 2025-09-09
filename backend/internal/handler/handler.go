package handler

// APIHandler combines all handlers to implement ServerInterface
type APIHandler struct {
	*UserHandler
	*QuestionHandler
	*FileHandler
}

func NewAPIHandler() *APIHandler {
	return &APIHandler{
		UserHandler:     NewUserHandler(),
		QuestionHandler: NewQuestionHandler(),
		FileHandler:     NewFileHandler(),
	}
}