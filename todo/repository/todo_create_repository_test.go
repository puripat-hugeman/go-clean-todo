package repository_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/puripat-hugeman/go-clean-todo/todo"
	"github.com/puripat-hugeman/go-clean-todo/todo/datamodel"
	"github.com/puripat-hugeman/go-clean-todo/todo/repository"
	"github.com/puripat-hugeman/go-clean-todo/todo/test/etc"
	mock "github.com/puripat-hugeman/go-clean-todo/todo/test/mockdata"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateTodoRepository(t *testing.T) {
	suite.Run(t, new(TestCreateTodoRepositoryTestSuite))
}

type TestCreateTodoRepositoryTestSuite struct {
	suite.Suite

	sqlMock    sqlmock.Sqlmock
	sqlMockDB  *sql.DB
	mockGormDB *gorm.DB

	repository todo.Repository
	datamodel  datamodel.TodoCreateEntity
}

func (suite *TestCreateTodoRepositoryTestSuite) SetupSuite() {

}

func (suite *TestCreateTodoRepositoryTestSuite) TeardownSuite() {

}

func (suite *TestCreateTodoRepositoryTestSuite) SetupTest() {
	var err error

	//Init DBMock
	suite.sqlMockDB, suite.sqlMock, err = sqlmock.New()
	assert.NoError(suite.T(), err)

	//Init Gorm
	suite.mockGormDB, err = gorm.Open(postgres.New(postgres.Config{
		DriverName:           mock.DriverName,
		DSN:                  mock.DSN,
		PreferSimpleProtocol: true,
		Conn:                 suite.sqlMockDB,
	}), &gorm.Config{})
	assert.NoError(suite.T(), err)

	//Init Repo
	suite.repository = repository.NewTodoRepository(suite.mockGormDB)
	suite.datamodel = mock.TodoCreatetEntityMockData()
}

func (suite *TestCreateTodoRepositoryTestSuite) TeardownTest() {
	assert.NoError(suite.T(), suite.sqlMock.ExpectationsWereMet())
}

func (suite *TestCreateTodoRepositoryTestSuite) TestSuccessCreateTodo() {
	suite.sqlMock.MatchExpectationsInOrder(true)
	suite.sqlMock.ExpectBegin()
	suite.sqlMock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "`+etc.TodoTable+`"`)).
		WithArgs(
			suite.datamodel.Uuid,
			suite.datamodel.Title,
			suite.datamodel.Image,
			suite.datamodel.CreatedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.sqlMock.ExpectCommit()
	result, err := suite.repository.CreateTodo(context.Background(), suite.datamodel)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.EqualValues(suite.T(), suite.datamodel, *result)
}

func (suite *TestCreateTodoRepositoryTestSuite) TestFailedCreateTodo() {
	suite.sqlMock.MatchExpectationsInOrder(true)
	suite.sqlMock.ExpectBegin()
	suite.sqlMock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "`+etc.TodoTable+`"`)).
		WithArgs(
			suite.datamodel.Uuid,
			suite.datamodel.Title,
			suite.datamodel.Image,
			suite.datamodel.CreatedAt,
		).
		WillReturnError(mock.RepositoryError)
	suite.sqlMock.ExpectCommit()
	result, err := suite.repository.CreateTodo(context.Background(), suite.datamodel)

	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), mock.RepositoryError.Error(), err.Error())
}
