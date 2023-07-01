package studenth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *studentHandlers) GetStudentAll(c *gin.Context) {

	studentResponse, err := h.studentService.GetStudentAll()
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err.Error())
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, studentResponse)

}
