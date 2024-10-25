package masterhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/services/masterservice"
	"net/http"
	_ "net/url"

	"github.com/gin-gonic/gin"
)

func (h *studentHandlers) AcceptPrivacyPolicy(c *gin.Context) {

	var requestBody masterservice.PrivacyPolicyRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnprocessableEntity, err.Error())
		c.Abort()
		return
	}

	privacyResponse, err := h.studentService.SetPrivacyPolicy(requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnprocessableEntity, err.Error())
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, privacyResponse)

}
