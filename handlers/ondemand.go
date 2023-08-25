package handlers

import (
	"RU-Smart-Workspace/ru-smart-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ondemandHandlers struct {
	ondemandServices services.OndemandServiceInterface
}

func NewOndemandHandlers(ondemandServices services.OndemandServiceInterface) ondemandHandlers {
	return ondemandHandlers{ondemandServices: ondemandServices}
}

func (h *ondemandHandlers) GetOndemandAll(c *gin.Context) {
	var requestBody services.OndemandRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	ondemandResponse, err := h.ondemandServices.GetOndemandAll(requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, ondemandResponse)

}

func (h *ondemandHandlers) GetOndemandSubjectCode(c *gin.Context) {
	var requestBody services.OndemandSubjectCodeRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	ondemandResponse, err := h.ondemandServices.GetOndemandSubjectCode(requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, ondemandResponse)

}
