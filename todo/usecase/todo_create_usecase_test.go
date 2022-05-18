package usecase_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/puripat-hugeman/go-clean-todo/todo"
	"github.com/puripat-hugeman/go-clean-todo/todo/datamodel"
	mock "github.com/puripat-hugeman/go-clean-todo/todo/test/mockdata"
	"github.com/puripat-hugeman/go-clean-todo/todo/test/mocktodo"
	"github.com/puripat-hugeman/go-clean-todo/todo/usecase"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestTodosUsecase(t *testing.T) {
	suite.Run(t, new(TestTodosUsecaseTestSuite))
}

type TestTodosUsecaseTestSuite struct {
	suite.Suite

	sqlMock    sqlmock.Sqlmock
	sqlMockDB  *sql.DB
	mockGormDB *gorm.DB

	mockRepository *mocktodo.Repository
	usecase        todo.UseCase
	request        datamodel.TodoRequestEntity
	datamodel      datamodel.TodoCreateEntity
}

func (suite *TestTodosUsecaseTestSuite) SetupSuite() {

}

func (suite *TestTodosUsecaseTestSuite) TeardownSuite() {

}

func (suite *TestTodosUsecaseTestSuite) SetupTest() {
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
	suite.request = mock.TodoRequestEntityMockData()
	suite.datamodel = mock.TodoCreatetEntityMockData()

}

func (suite *TestTodosUsecaseTestSuite) TeardownTest() {
	assert.NoError(suite.T(), suite.sqlMock.ExpectationsWereMet())
}

func (suite *TestTodosUsecaseTestSuite) TestSuccessCreateTodo() {
	suite.sqlMock.MatchExpectationsInOrder(true)
	suite.mockRepository.On("CreateTodo", context.Background(), testifyMock.MatchedBy(func(datamodel datamodel.TodoCreateEntity) bool {
		return suite.datamodel.CreatedAt.Day() == datamodel.CreatedAt.Day() &&
			suite.datamodel.CreatedAt.Hour() == datamodel.CreatedAt.Hour() &&
			suite.datamodel.CreatedAt.Minute() == datamodel.CreatedAt.Minute()
	})).Return(&suite.datamodel, nil)
	result, err := suite.usecase.CreateTodo(context.Background(), suite.request)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.EqualValues(suite.T(), suite.datamodel, *result)
}

func (suite *TestTodosUsecaseTestSuite) TestFailedCreateTodo() {
	suite.mockRepository.On("CreateTodo", context.Background(), testifyMock.MatchedBy(func(datamodel datamodel.TodoCreateEntity) bool {
		return suite.datamodel.CreatedAt.Day() == datamodel.CreatedAt.Day() &&
			suite.datamodel.CreatedAt.Hour() == datamodel.CreatedAt.Hour() &&
			suite.datamodel.CreatedAt.Minute() == datamodel.CreatedAt.Minute()
	})).Return(nil, mock.UsecaseError)
	result, err := suite.usecase.CreateTodo(context.Background(), suite.request)

	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), mock.UsecaseError.Error())
}
