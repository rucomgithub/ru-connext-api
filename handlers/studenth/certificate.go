package studenth

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/services/students"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *studentHandlers) Certifiate(c *gin.Context) {

	var requestBody students.CertifiatePlayload

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Certificate fail becourse content type not json format."})
		c.Abort()
		return
	}

	tokenResponse, err := h.studentService.Certificate(requestBody.Std_code, requestBody.Certifiate)
	if err != nil {
		c.Error(errors.New(err.Error() + ", " + requestBody.Std_code))
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnprocessableEntity, tokenResponse)
		c.Abort()
		return
	}
	c.IndentedJSON(http.StatusOK, tokenResponse)
}
