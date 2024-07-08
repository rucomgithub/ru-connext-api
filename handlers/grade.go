package handlers

import (
	"RU-Smart-Workspace/ru-smart-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type gradeHandlers struct {
	gradeServices services.GradeServiceInterface
}

func NewgradeHandlers(gradeServices services.GradeServiceInterface) gradeHandlers {
	return gradeHandlers{gradeServices: gradeServices}
}

func (h *gradeHandlers) GradeYear(c *gin.Context) {
	var requestBody services.GradeRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {

		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	gradeResponse, err := h.gradeServices.GradeYear(requestBody)
	if err != nil {

		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gradeResponse)

}

func (h *gradeHandlers) Grades(c *gin.Context) {
	var std_code string = c.Param("std_code")

	gradeResponse, err := h.gradeServices.GradeAll(std_code)
	if err != nil {

		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gradeResponse)

}
