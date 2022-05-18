package repository

import (
	"context"
	"log"

	"github.com/puripat-hugeman/go-clean-todo/todo/datamodel"
	repository "github.com/puripat-hugeman/go-clean-todo/todo/repository/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&repository.TodoCreateRepository{}); err != nil {
		log.Panicln(err)

	}
}

func (repo *TodoRepository) CreateTodo(ctx context.Context, todo datamodel.TodoCreateEntity) (*datamodel.TodoCreateEntity, error) {
	req := repository.TodoCreateRepository{
		Uuid:      todo.Uuid,
		Title:     todo.Title,
		Image:     todo.Image,
		CreatedAt: todo.CreatedAt,
		// UpdatedAt: todo.UpdatedAt,
	}
	if err := repo.db.WithContext(ctx).Create(req).Error; err != nil {
		return nil, errors.New("RepositoryError")
	}
	return &datamodel.TodoCreateEntity{
		Uuid:      req.Uuid,
		Title:     req.Title,
		Image:     req.Image,
		CreatedAt: req.CreatedAt,
		// UpdatedAt: req.UpdatedAt,
	}, nil
}

func (repo *TodoRepository) GetTodos(ctx context.Context) ([]datamodel.TodoGetEntity, error) {
	var todos []repository.TodoCreateRepository
	if err := repo.db.WithContext(ctx).Find(&todos).Error; err != nil {
		return nil, errors.New("RepositoryError")
	}

	var results = make([]datamodel.TodoGetEntity, len(todos))
	//package copier
	for i, todo := range todos {
		results[i] = datamodel.TodoGetEntity{
			Uuid:      todo.Uuid,
			Title:     todo.Title,
			Image:     todo.Image,
			CreatedAt: todo.CreatedAt,
			// UpdatedAt: todo.UpdatedAt,
		}
	}

	return results, nil
}
