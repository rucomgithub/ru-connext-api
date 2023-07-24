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

func (h *registerHandlers) Registers(c *gin.Context) {
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

func (h *registerHandlers) Years(c *gin.Context) {
	var std_code string = c.Param("std_code")

	if std_code == "" {
		handleError(c, errors.New("โปรดระบุรหัสนักศึกษา"))
		return
	}

	registerResponse, err := h.registerServices.GetListYear(std_code)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, registerResponse)
}

func (h *registerHandlers) YearSemesters(c *gin.Context) {
	var std_code string = c.Param("std_code")

	if std_code == "" {
		handleError(c, errors.New("โปรดระบุรหัสนักศึกษา"))
		return
	}

	registerResponse, err := h.registerServices.GetListYearSemester(std_code)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, registerResponse)
}

func (h *registerHandlers) ScheduleYearSemesters(c *gin.Context) {
	var std_code string = c.Param("std_code")

	if std_code == "" {
		handleError(c, errors.New("โปรดระบุรหัสนักศึกษา"))
		return
	}

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

func (h *registerHandlers) Schedules(c *gin.Context) {
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

func (h *registerHandlers) YearSemesterLates(c *gin.Context) {

	mr30Response, err := h.registerServices.GetYearSemesterLatest()
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, mr30Response)

}
