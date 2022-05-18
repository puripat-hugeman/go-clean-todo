package restful

import (
	"github.com/puripat-hugeman/go-clean-todo/todo"
)

type TodoHandler struct {
	usecase todo.UseCase
}

func NewHandler(usecase todo.UseCase) *TodoHandler {
	return &TodoHandler{
		usecase: usecase,
	}
}
