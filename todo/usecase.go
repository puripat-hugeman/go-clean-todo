package todo

import (
	"context"

	"github.com/puripat-hugeman/go-clean-todo/todo/datamodel"
)

type UseCase interface {
	CreateTodo(ctx context.Context, todo datamodel.TodoRequestEntity) (result *datamodel.TodoCreateEntity, err error)
	GetTodos(ctx context.Context) (results []datamodel.TodoGetEntity, err error)
}
