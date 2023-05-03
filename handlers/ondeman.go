package handlers

import (
	"RU-Smart-Workspace/ru-smart-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type onDemandHandlers struct {
	ondemandServices services.OnDemandServiceInterface
}

func NewOnDeMandHandlers(ondemandServices services.OnDemandServiceInterface) onDemandHandlers {
	return onDemandHandlers{ondemandServices: ondemandServices}
}

func (h *onDemandHandlers) Healthz(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "The service works normally.",
	})
}

func (h *onDemandHandlers) GetOnDemandAll(c *gin.Context) {
	var requestBody services.OnDemandRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		handleError(c, err)
		return
	}

	gradeResponse, err := h.ondemandServices.OnDemandAll(requestBody)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gradeResponse)

}
