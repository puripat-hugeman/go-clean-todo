package usecase_test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/puripat-hugeman/go-clean-todo/todo"
	"github.com/puripat-hugeman/go-clean-todo/todo/datamodel"
	mock "github.com/puripat-hugeman/go-clean-todo/todo/test/mockdata"
	"github.com/puripat-hugeman/go-clean-todo/todo/test/mocktodo"
	"github.com/puripat-hugeman/go-clean-todo/todo/usecase"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetTodosUsecase(t *testing.T) {
	suite.Run(t, new(TestGetTodosUsecaseTestSuite))
}

type TestGetTodosUsecaseTestSuite struct {
	suite.Suite

	sqlMock    sqlmock.Sqlmock
	sqlMockDB  *sql.DB
	mockGormDB *gorm.DB

	mockRepository *mocktodo.Repository
	usecase        todo.UseCase
	datamodel      []datamodel.TodoGetEntity
}

func (suite *TestGetTodosUsecaseTestSuite) SetupSuite() {

}

func (suite *TestGetTodosUsecaseTestSuite) TeardownSuite() {

}

func (suite *TestGetTodosUsecaseTestSuite) SetupTest() {
	var err error

	suite.sqlMockDB, suite.sqlMock, err = sqlmock.New()
	assert.NoError(suite.T(), err)

	suite.mockGormDB, err = gorm.Open(postgres.New(postgres.Config{
		DriverName:           mock.DriverName,
		DSN:                  mock.DSN,
		PreferSimpleProtocol: true,
		Conn:                 suite.sqlMockDB,
	}), &gorm.Config{})
	assert.NoError(suite.T(), err)

	suite.mockRepository = new(mocktodo.Repository)
	suite.usecase = usecase.NewTodoUseCase(suite.mockRepository)
	suite.datamodel = []datamodel.TodoGetEntity{
		mock.TodoGetEntityMockData(),
		mock.TodoGetEntityMockData(),
	}
	log.Println("usecase: ", &suite.usecase)
}

func (suite *TestGetTodosUsecaseTestSuite) TeardownTest() {

}

func (suite *TestGetTodosUsecaseTestSuite) TestGetPass() {
	suite.sqlMock.MatchExpectationsInOrder(true)
	suite.mockRepository.On("GetTodos", context.Background()).Return(suite.datamodel, nil)
	suite.usecase = usecase.NewTodoUseCase(suite.mockRepository)
	results, err := suite.usecase.GetTodos(context.Background())

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), results)
	assert.EqualValues(suite.T(), suite.datamodel, results)
}

func (suite *TestGetTodosUsecaseTestSuite) TestGetFailed() {
	suite.mockRepository.On("GetTodos", context.Background()).Return(nil, mock.UsecaseError)
	results, err := suite.usecase.GetTodos(context.Background())

	assert.Nil(suite.T(), results)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), mock.UsecaseError.Error())
}
