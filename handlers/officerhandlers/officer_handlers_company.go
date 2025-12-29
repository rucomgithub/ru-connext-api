package officerhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"fmt"
	"net/http"
	_ "net/url"

	"github.com/gin-gonic/gin"
)

func (h *officerHandlers) GetCommpanyList(c *gin.Context) {

	var std_code string = c.Param("id")

	fmt.Println(std_code)

	companys,count, err := h.officerServices.GetCommpanyList(std_code)

	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "ไม่พบข้อมูล บริษัท/ห้างร้าน/หน่วยงาน/ส่วนราชการ ของนักศึกษารหัส " + std_code + " ในระบบ."})
		c.Abort()
		return
	}

	message := fmt.Sprintf(`พบข้อมูล บริษัท/ห้างร้าน/หน่วยงาน/ส่วนราชการ ของนักศึกษารหัส %s ในระบบ จำนวน %d รายการ.`,std_code,count)

	c.IndentedJSON(http.StatusOK, gin.H{"companys": companys,"count": count,"message": message})

}