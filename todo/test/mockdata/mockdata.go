package mock

import (
	"errors"
	"time"

	"github.com/puripat-hugeman/go-clean-todo/todo/datamodel"

	"github.com/google/uuid"
	"github.com/goombaio/namegenerator"
)

const (
	DriverName     = "postgres"
	DSN            = "mockDB"
	ImageMock      = "imgBase64Str"
	MultipartImage = "../../test/multipart/ImageMock"
	GetPath        = "/todo"
	CreatePath     = "/todo/create"
)

var (
	UsecaseError    = errors.New("UsecaseError")
	RepositoryError = errors.New("RepositoryError")
)

func NewName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	name := nameGenerator.Generate()
	return name
}

func TodoGetEntityMockData() datamodel.TodoGetEntity {
	return datamodel.TodoGetEntity{
		Uuid:      uuid.NewString(),
		Title:     NewName(),
		Image:     ImageMock,
		CreatedAt: time.Now(),
		// UpdatedAt: time.Now(),
	}
}

func TodoCreatetEntityMockData() datamodel.TodoCreateEntity {
	return datamodel.TodoCreateEntity{
		Uuid:      uuid.NewString(),
		Title:     NewName(),
		Image:     ImageMock,
		CreatedAt: time.Now(),
		// UpdatedAt: time.Now(),
	}
}

func TodoRequestEntityMockData() datamodel.TodoRequestEntity {
	return datamodel.TodoRequestEntity{
		Title: NewName(),
		Image: ImageMock,
	}
}

func TodoRequestHandlerCreateMockData() datamodel.TodoRequestEntity {
	return datamodel.TodoRequestEntity{
		Title: "test",
		Image: "",
	}
}

func TodoRequestBodyCreateMockData(request datamodel.TodoRequestEntity) string {
	return `
	{
	"title":"` + request.Title + `",
	"imageUrl":"` + request.Image + `"
	}
`
}
