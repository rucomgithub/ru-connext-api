package handlers

import (
	"RU-Smart-Workspace/ru-smart-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type registerHandlers struct {
	registerServices services.RegisterServiceInterface
}

func NewRegisterHandlers(registerServices services.RegisterServiceInterface) registerHandlers {
	return registerHandlers{registerServices: registerServices}
}

func (h *registerHandlers) GetRegister(c *gin.Context) {
	var requestBody services.RegisterRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		handleError(c, err)
		return
	}

	registerResponse, err := h.registerServices.GetRegister(requestBody)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, registerResponse)

}

func (h *registerHandlers) GetRegisterYear(c *gin.Context) {
	var std_code string = c.Param("std_code")
	registerResponse, err := h.registerServices.GetRegisterYear(std_code)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, registerResponse)
}

func (h *registerHandlers) GetRegisterGroupYearSemester(c *gin.Context) {
	var std_code string = c.Param("std_code")
	registerResponse, err := h.registerServices.GetRegisterGroupYearSemester(std_code)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, registerResponse)
}

func (h *registerHandlers) GetRegisterMr30(c *gin.Context) {
	var std_code string = c.Param("std_code")

	var requestBody services.RegisterMr30Request

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		handleError(c, err)
		return
	}

	registerResponse, err := h.registerServices.GetRegisterMr30(std_code, requestBody)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, registerResponse)
}

func (h *registerHandlers) GetRegisterMr30Latest(c *gin.Context) {
	var std_code string = c.Param("std_code")

	var requestBody services.RegisterMr30Request

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		handleError(c, err)
		return
	}

	registerResponse, err := h.registerServices.GetRegisterMr30Latest(std_code)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, registerResponse)
}
