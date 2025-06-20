package officerhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/services/officerservices"
	"RU-Smart-Workspace/ru-smart-api/services/students"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *officerHandlers) Authentication(c *gin.Context) {
	var requestBody officerservices.AuthenRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", getLineNumber())
		c.Set("file", getFileName())
		handleError(c, err)
		return
	}

	authenResponse, err := h.officerServices.AuthenticationOfficer(requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", getLineNumber())
		c.Set("file", getFileName())
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, authenResponse)

}

func (h *officerHandlers) RefreshAuthentication(c *gin.Context) {

	var requestBody students.AuthenPlayload

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", getLineNumber())
		c.Set("file", getFileName())
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Authorization fail becourse content type not json format."})
		c.Abort()
		return
	}

	tokenRespone, err := h.officerServices.RefreshAuthenticationOfficer(requestBody.Refresh_token)
	if err != nil {

		c.Set("line", getLineNumber())
		c.Set("file", getFileName())
		c.IndentedJSON(http.StatusForbidden, tokenRespone)
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, tokenRespone)

}
