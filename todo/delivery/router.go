package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/puripat-hugeman/go-clean-todo/todo"
)

func NewHttpHandler(todoUsecase todo.UseCase) http.Handler {
	router := gin.Default()

	// todoGroup := router.Group("/todo")
	// server.NewRoutesFactory(todoGroup)(todoUsecase)
	return router
}
