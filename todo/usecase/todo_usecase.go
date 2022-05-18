package usecase

import (
	"context"
	"time"

	"github.com/puripat-hugeman/go-clean-todo/todo"
	"github.com/puripat-hugeman/go-clean-todo/todo/datamodel"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type TodoUseCase struct {
	todoRepo todo.Repository
}

func NewTodoUseCase(todoRepo todo.Repository) *TodoUseCase {
	return &TodoUseCase{
		todoRepo: todoRepo,
	}
}

func (t *TodoUseCase) CreateTodo(ctx context.Context, todo datamodel.TodoRequestEntity) (result *datamodel.TodoCreateEntity, err error) {
	result, err = t.todoRepo.CreateTodo(ctx, datamodel.TodoCreateEntity{
		Uuid:      uuid.NewString(),
		Title:     todo.Title,
		Image:     todo.Image,
		CreatedAt: time.Now(),
		// UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "UsecaseError")
	}
	return result, nil
}

func (t *TodoUseCase) GetTodos(ctx context.Context) ([]datamodel.TodoGetEntity, error) {
	todos, err := t.todoRepo.GetTodos(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "UsecaseError")
	}
	return todos, nil
}
