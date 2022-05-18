package restful_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/puripat-hugeman/go-clean-todo/todo/datamodel"
	"github.com/puripat-hugeman/go-clean-todo/todo/delivery/restful"
	mock "github.com/puripat-hugeman/go-clean-todo/todo/test/mockdata"
	"github.com/puripat-hugeman/go-clean-todo/todo/test/mocktodo"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestGetTodosHandler(t *testing.T) {

}

type TestGetTodosHandlerTestSuite struct {
	suite.Suite

	engine *gin.Engine

	mockUsecase *mocktodo.UseCase
	todos       []datamodel.TodoCreateEntity
}

//coverage 0% !!
func (suite *TestGetTodosHandlerTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	suite.engine = gin.New()
	suite.mockUsecase = new(mocktodo.UseCase)
	testHandler := restful.NewHandler(suite.mockUsecase)
	suite.engine.GET(mock.GetPath, testHandler.GetTodoHandler)
}

func (suite *TestGetTodosHandlerTestSuite) TeardownSuite() {

}

func (suite *TestGetTodosHandlerTestSuite) SetupTest() {
	suite.mockUsecase.ExpectedCalls = []*testifyMock.Call{}
	suite.todos = []datamodel.TodoCreateEntity{
		mock.TodoCreatetEntityMockData(),
		mock.TodoCreatetEntityMockData(),
	}

}

func (suite *TestGetTodosHandlerTestSuite) TeardownTest() {

}

func (suite *TestGetTodosHandlerTestSuite) TestSuccessGetTodos() {
	suite.mockUsecase.On("GetTodos", context.Background()).
		Return(&suite.todos, nil)

	res := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, mock.GetPath, nil)
	assert.NoError(suite.T(), err)

	suite.engine.ServeHTTP(res, req)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
}

func (suite *TestGetTodosHandlerTestSuite) TestFailedInternalServerError() {
	suite.mockUsecase.On("GetTodos", context.Background()).
		Return(nil, mock.UsecaseError)

	res := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, mock.GetPath, nil)
	assert.NoError(suite.T(), err)

	suite.engine.ServeHTTP(res, req)
	assert.Equal(suite.T(), http.StatusInternalServerError, res.Code)
	assert.Equal(suite.T(), err.Error(), mock.UsecaseError.Error())
}
