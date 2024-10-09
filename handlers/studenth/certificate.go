package studenth

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *studentHandlers) Certifiate(c *gin.Context) {

	ID_TOKEN, err := middlewares.GetHeaderAuthorization(c)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "authorization key in header not found"})
		c.Abort()
		return
	}

	tokenResponse, err := h.studentService.Certificate(ID_TOKEN)
	if err != nil {
		c.Error(errors.New(err.Error() + ", " + ID_TOKEN))
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnprocessableEntity, tokenResponse)
		c.Abort()
		return
	}
	c.IndentedJSON(http.StatusOK, tokenResponse)
}
