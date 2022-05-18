package server

import "github.com/puripat-hugeman/go-clean-todo/todo/delivery/restful"

func (s *Server) SetUpRoutes(handler restful.TodoHandler) {
	todoAPI := s.engine.Group("/todo")
	todoAPI.GET("/", handler.GetTodoHandler)
	todoAPI.POST("/create", handler.CreateTodoHandler)
}

// import (
// 	"encoding/base64"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/puripat-hugeman/go-clean-todo/todo"
// 	"github.com/puripat-hugeman/go-clean-todo/todo/datamodel"
// 	httpmodel "github.com/puripat-hugeman/go-clean-todo/todo/delivery/http_model"
// 	"github.com/puripat-hugeman/go-clean-todo/todo/delivery/utils"
// 	"github.com/puripat-hugeman/go-clean-todo/todo/enums"

// 	"github.com/gin-gonic/gin"
// )

// func NewRoutesFactory(group *gin.RouterGroup) func(usecase todo.UseCase) {
// 	todoRoutesFactory := func(usecase todo.UseCase) {

// 		group.GET("/", func(c *gin.Context) {
// 			ctx := c.Request.Context()
// 			response, err := usecase.GetTodos(ctx)
// 			if err != nil {
// 				log.Println(err)
// 			}
// 			c.JSON(http.StatusOK, response)
// 		})
// 		group.POST("/create", func(c *gin.Context) {
// 			formData, err := utils.ExtractTodoMultipartFileAndData(c)
// 			if err != nil && !errors.Is(err, enums.ErrFile) {
// 				status := utils.ErrStatus(enums.MapErrHandler.MultipartError, err)
// 				c.AbortWithStatusJSON(http.StatusBadRequest, status)
// 				return
// 			}
// 			var req httpmodel.TodoRequestDelivery
// 			if err := json.Unmarshal(formData.JSONData, &req); err != nil {
// 				status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
// 				c.AbortWithStatusJSON(http.StatusBadRequest, status)
// 				return
// 			}
// 			var imgBytes []byte
// 			if formData.FileData != nil {
// 				imgBytes = formData.FileData
// 			} else {
// 				// Default to-do image
// 				imgBytes, err = os.ReadFile("./assets/default.jpg")
// 				if err != nil {
// 					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 						"status": "failed to set default image for todo",
// 						"error":  fmt.Sprintf("file read error: %v", err.Error()),
// 					})
// 				}
// 			}
// 			// Convert to Base64 string
// 			imgBase64Str := base64.StdEncoding.EncodeToString(imgBytes)
// 			if l := len(imgBase64Str); l > enums.POSTGRES_MAX_STRLEN {
// 				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 					"status": "failed to set image for todo",
// 					"error":  fmt.Sprintf("image file too large: %d > %d", l, enums.POSTGRES_MAX_STRLEN),
// 				})
// 				return
// 			}
// 			todo, err := datamodel.NewTodo(
// 				req.Title,
// 				// enums.InProgress,
// 				imgBase64Str)

// 			if err != nil {
// 				log.Println("here")
// 				c.AbortWithStatusJSON(
// 					http.StatusBadRequest,
// 					gin.H{
// 						"status": "todo creation failed",
// 						"error":  err.Error()},
// 				)
// 			}
// 			ctx := c.Request.Context()

// 			resp, err := usecase.CreateTodo(ctx, todo)
// 			if err != nil {
// 				c.AbortWithStatusJSON(
// 					http.StatusBadRequest,
// 					gin.H{
// 						"status": "todo creation failed",
// 						"error":  err.Error()},
// 				)
// 			}
// 			c.JSON(http.StatusCreated, resp)
// 		})
// 	}
// 	return todoRoutesFactory
// }
