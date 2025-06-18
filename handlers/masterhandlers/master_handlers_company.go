package masterhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"RU-Smart-Workspace/ru-smart-api/services/masterservice"
	"fmt"
	"net/http"
	_ "net/url"

	"github.com/gin-gonic/gin"
)

func (h *studentHandlers) AddCommpany(c *gin.Context) {
	var requestBody masterservice.CompanyRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		c.Abort()
		return
	}

	qf, err := h.studentService.AddCommpany(requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "บันทึกผู้ขอตรวจสอบข้อมูล รหัสนักศึกษา " + requestBody.STD_CODE + " ซ้ำกันในระบบ."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"companys": qf, "message": "บันทึกผู้ขอตรวจสอบข้อมูล รหัสนักศึกษา " + requestBody.STD_CODE + " สำเร็จ."})

}
