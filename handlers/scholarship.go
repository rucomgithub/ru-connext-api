package handlers

import (
	"RU-Smart-Workspace/ru-smart-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type scholarShipHandlers struct {
	scholarShipServices services.ScholarShipServiceInterface
}

func NewScholarShipHandlers(scholarShipServices services.ScholarShipServiceInterface) scholarShipHandlers {
	return scholarShipHandlers{scholarShipServices: scholarShipServices}
}

func (h *scholarShipHandlers) GetScholarshipAll(c *gin.Context) {

	var requestBody services.ScholarShipRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		handleError(c, err)
		return
	}
	scholarShipResponse, err := h.scholarShipServices.GetScholarshipAll(requestBody)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, scholarShipResponse)

}
