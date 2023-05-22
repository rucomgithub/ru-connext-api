package handlers

import (
	"RU-Smart-Workspace/ru-smart-api/services"
	"errors"
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

func (h *registerHandlers) GetYear(c *gin.Context) {
	var std_code string = c.Param("std_code")
	registerResponse, err := h.registerServices.GetListYear(std_code)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, registerResponse)
}

func (h *registerHandlers) GetYearSemester(c *gin.Context) {
	var std_code string = c.Param("std_code")
	registerResponse, err := h.registerServices.GetListYearSemester(std_code)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, registerResponse)
}

func (h *registerHandlers) GetScheduleYearSemester(c *gin.Context) {
	var std_code string = c.Param("std_code")

	var requestBody services.RegisterScheduleRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		handleError(c, err)
		return
	}

	registerResponse, err := h.registerServices.GetScheduleYearSemester(std_code, requestBody)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, registerResponse)
}

func (h *registerHandlers) GetSchedule(c *gin.Context) {
	var std_code string = c.Param("std_code")

	if std_code == "" {
		handleError(c, errors.New("โปรดระบุรหัสนักศึกษา"))
		return
	}

	registerResponse, err := h.registerServices.GetSchedule(std_code)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, registerResponse)
}
