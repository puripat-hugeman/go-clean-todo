package restful

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *TodoHandler) GetTodoHandler(c *gin.Context) {
	ctx := c.Request.Context()
	response, err := h.usecase.GetTodos(ctx)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "failed to get todos",
			"error":  fmt.Sprintf("file read error: %v", err.Error()),
		})
	}
	c.JSON(http.StatusOK, response)
}
