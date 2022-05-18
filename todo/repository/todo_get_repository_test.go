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

func TestGetTodosRepository(t *testing.T) {
	suite.Run(t, new(TestGetTodosRepositoryTestSuite))
}

type TestGetTodosRepositoryTestSuite struct {
	suite.Suite

	sqlMock    sqlmock.Sqlmock
	sqlMockDB  *sql.DB
	mockGormDB *gorm.DB

	repository todo.Repository
	datamodel  []datamodel.TodoGetEntity
}

func (suite *TestGetTodosRepositoryTestSuite) SetupSuite() {

}

func (suite *TestGetTodosRepositoryTestSuite) TeardownSuite() {

}

func (suite *TestGetTodosRepositoryTestSuite) SetupTest() {
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

	suite.repository = repository.NewTodoRepository(suite.mockGormDB)
	suite.datamodel = []datamodel.TodoGetEntity{
		mock.TodoGetEntityMockData(),
		mock.TodoGetEntityMockData(),
	}
}

func (suite *TestGetTodosRepositoryTestSuite) TeardownTest() {
	assert.NoError(suite.T(), suite.sqlMock.ExpectationsWereMet())
}

func (suite *TestGetTodosRepositoryTestSuite) TestSuccessGetTodos() {
	suite.sqlMock.MatchExpectationsInOrder(true)
	suite.sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "` + etc.TodoTable + `"`)).
		WillReturnRows(sqlmock.NewRows([]string{
			etc.TodoIDColumn,
			etc.TodoTitleColumn,
			etc.TodoImageURLColumn,
			etc.TodoCreatedAtColumn,
		}).
			AddRow(
				suite.datamodel[0].Uuid,
				suite.datamodel[0].Title,
				suite.datamodel[0].Image,
				suite.datamodel[0].CreatedAt,
			).
			AddRow(suite.datamodel[1].Uuid,
				suite.datamodel[1].Title,
				suite.datamodel[1].Image,
				suite.datamodel[1].CreatedAt,
			))

	results, err := suite.repository.GetTodos(context.Background())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), results)
	assert.Equal(suite.T(), len(suite.datamodel), len(results))
	assert.Equal(suite.T(), suite.datamodel, results)
}

func (suite *TestGetTodosRepositoryTestSuite) TestFailedGetTodos() {

	suite.sqlMock.MatchExpectationsInOrder(true)
	suite.sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "` + etc.TodoTable + `"`)).
		WillReturnError(mock.RepositoryError)

	results, err := suite.repository.GetTodos(context.Background())

	assert.Nil(suite.T(), results)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), mock.RepositoryError.Error(), err.Error())
}
