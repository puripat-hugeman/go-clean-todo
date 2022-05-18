package restful_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/puripat-hugeman/go-clean-todo/todo/datamodel"
	"github.com/puripat-hugeman/go-clean-todo/todo/delivery/restful"
	mock "github.com/puripat-hugeman/go-clean-todo/todo/test/mockdata"
	"github.com/puripat-hugeman/go-clean-todo/todo/test/mocktodo"
	multiparttest "github.com/puripat-hugeman/go-clean-todo/todo/test/multipart"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestCreateTodoHandler(t *testing.T) {
	suite.Run(t, new(TestCreateTodoHandlerTestSuite))
}

type TestCreateTodoHandlerTestSuite struct {
	suite.Suite

	engine *gin.Engine

	mockUsecase *mocktodo.UseCase
	request     datamodel.TodoRequestEntity
	errRequest  datamodel.TodoRequestEntity
	requestBody string
	datamodel   datamodel.TodoCreateEntity
}

func (suite *TestCreateTodoHandlerTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	suite.engine = gin.New()
	suite.mockUsecase = new(mocktodo.UseCase)
	testHandler := restful.NewHandler(suite.mockUsecase)
	suite.engine.POST(mock.CreatePath, testHandler.CreateTodoHandler)

}

func (suite *TestCreateTodoHandlerTestSuite) TeardownSuite() {

}

//should coverage more than 80%
func (suite *TestCreateTodoHandlerTestSuite) SetupTest() {
	suite.mockUsecase.ExpectedCalls = []*testifyMock.Call{}
	suite.request = mock.TodoRequestHandlerCreateMockData()
	suite.errRequest = datamodel.TodoRequestEntity{}
	suite.datamodel = mock.TodoCreatetEntityMockData()
	suite.datamodel.Title = suite.request.Title
	suite.requestBody = mock.TodoRequestBodyCreateMockData(suite.request)
}

func (suite *TestCreateTodoHandlerTestSuite) TeardownTest() {
}

func (suite *TestCreateTodoHandlerTestSuite) TestSuccessCreateTodo() {
	request, err := multiparttest.PrepareMultipartRequest(suite.request, mock.MultipartImage)
	assert.NoError(suite.T(), err)

	suite.mockUsecase.On("CreateTodo",
		context.Background(),
		suite.request).Return(&suite.datamodel, nil)

	rw := httptest.NewRecorder()

	suite.engine.ServeHTTP(rw, request)
	assert.Equal(suite.T(), http.StatusCreated, rw.Code)
}

func (suite *TestCreateTodoHandlerTestSuite) TestFailedCreateTodoBadRequest() {
	suite.mockUsecase.On("CreateTodo",
		context.Background(),
		suite.request).Return(&suite.datamodel, nil)

	rw := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPost,
		mock.CreatePath,
		strings.NewReader(suite.requestBody))
	assert.NoError(suite.T(), err)

	suite.engine.ServeHTTP(rw, req)
	assert.Equal(suite.T(), http.StatusBadRequest, rw.Code)
}

func (suite *TestCreateTodoHandlerTestSuite) TestFailedInternalServerError() {
	request, err := multiparttest.PrepareMultipartRequest(suite.request, mock.MultipartImage)
	assert.NoError(suite.T(), err)

	suite.mockUsecase.On("CreateTodo",
		context.Background(),
		suite.request).Return(nil, mock.UsecaseError)

	rw := httptest.NewRecorder()

	suite.engine.ServeHTTP(rw, request)
	assert.Equal(suite.T(), http.StatusInternalServerError, rw.Code)
}
