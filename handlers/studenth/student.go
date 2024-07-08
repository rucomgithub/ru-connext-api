package studenth

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *studentHandlers) GetStudentAll(c *gin.Context) {

	studentResponse, err := h.studentService.GetStudentAll()
	if err != nil {

		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnprocessableEntity, err.Error())
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, studentResponse)

}
