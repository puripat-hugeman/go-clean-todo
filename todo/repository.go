package todo

import (
	"context"

	"github.com/puripat-hugeman/go-clean-todo/todo/datamodel"
)

type Repository interface {
	CreateTodo(ctx context.Context, todo datamodel.TodoCreateEntity) (response *datamodel.TodoCreateEntity, err error)
	GetTodos(ctx context.Context) (response []datamodel.TodoGetEntity, err error)
}
