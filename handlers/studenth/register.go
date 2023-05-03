package studenth

import (
	"RU-Smart-Workspace/ru-smart-api/services/students"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *studentHandlers) GetRegisterAll(c *gin.Context) {

	var requestBody students.RegisterAllRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err.Error())
		c.Abort()
		return
	}

	registerResponse, err := h.studentService.GetRegisterAll(requestBody.STD_CODE, requestBody.YEAR)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err.Error())
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, registerResponse)

}
