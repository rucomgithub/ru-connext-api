package handlers

import (
	"RU-Smart-Workspace/ru-smart-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type rotcsHandlers struct {
	rotcsServices services.RotcsServiceInterface
}

func NewRotcsHandlers(rotcsServices services.RotcsServiceInterface) rotcsHandlers {
	return rotcsHandlers{rotcsServices: rotcsServices}
}

func (h *rotcsHandlers) GetRotcsRegister(c *gin.Context) {
	var requestBody services.RotcsRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {

		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	rotcsResponse, err := h.rotcsServices.GetRotcsRegister(requestBody)
	if err != nil {

		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, rotcsResponse)

}

func (h *rotcsHandlers) GetRotcsExtend(c *gin.Context) {
	var requestBody services.RotcsRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {

		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	rotcsResponse, err := h.rotcsServices.GetRotcsExtend(requestBody)
	if err != nil {

		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, rotcsResponse)

}
