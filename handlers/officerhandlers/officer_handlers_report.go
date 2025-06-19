package officerhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"net/http"
	"errors"
	_ "net/url"
	"RU-Smart-Workspace/ru-smart-api/services/officerservices"
	"github.com/gin-gonic/gin"
)
func (h *officerHandlers) GetReport(c *gin.Context) {

	var requestBody officerservices.ReportRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		ErrValidateRequest(c)
		return
	}

	reportResponse, err := h.officerServices.GetReport(&requestBody)
	if err != nil {
		err = errors.New("ไม่พบข้อมูลรายงาน.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ไม่พบข้อมูลรายงาน."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"report": reportResponse, "message": "พบข้อมูลรายงาน."})

}