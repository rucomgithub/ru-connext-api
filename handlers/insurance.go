package handlers

import (
	"RU-Smart-Workspace/ru-smart-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InsuranceHandlers struct {
	insuranceServices services.InsuranceServiceInterface
}

func NewInsuranceHandlers(insuranceServices services.InsuranceServiceInterface) InsuranceHandlers {
	return InsuranceHandlers{insuranceServices: insuranceServices}
}

func (h *InsuranceHandlers) GetInsuranceListAll(c *gin.Context) {
	var requestBody services.InsuranceRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	insuranceResponse, err := h.insuranceServices.GetInsuranceListAll(requestBody)

	if err != nil {
		c.Error(err)
		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, insuranceResponse)

}
