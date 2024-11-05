package officerhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"RU-Smart-Workspace/ru-smart-api/services/masterservice"
	"errors"
	"fmt"
	"net/http"
	_ "net/url"

	"github.com/gin-gonic/gin"
)

func (h *officerHandlers) GetQualificationAll(c *gin.Context) {

	token, err := middlewares.GetHeaderAuthorization(c)

	if err != nil {
		ErrTokenLogin(c)
		return
	}

	claim, err := middlewares.GetClaimsOfficer(token)

	if err != nil {
		ErrTokenClaim(c)
		return
	}

	role := claim.Role

	if role == "Bachelor" {
		ErrRoleBachelor(c)
		return
	}

	qualificationResponse, total, err := h.officerServices.GetQualificationAll()
	if err != nil {
		err = errors.New("ไม่พบข้อมูลการยื่นขอเอกสาร.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ไม่พบข้อมูลการยื่นขอเอกสาร."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "all", "qualifications": qualificationResponse, "total": total, "message": "พบรายการข้อมูลยื่นขอเอกสารของนักศึกษาในระบบ."})

}

func (h *officerHandlers) GetQualification(c *gin.Context) {

	var STD_CODE string = c.Param("id")

	fmt.Println(STD_CODE)

	qf, err := h.officerServices.GetQualification(STD_CODE)

	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "ไม่พบข้อมูลยื่นขอเอกสารของนักศึกษารหัส " + STD_CODE + " ในระบบ."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": qf, "message": "พบข้อมูลยื่นขอเอกสารของนักศึกษารหัส " + STD_CODE + " ในระบบ."})

}

func (h *officerHandlers) UpdateQualification(c *gin.Context) {

	var requestBody masterservice.QualificationRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		ErrValidateRequest(c)
		return
	}

	var STD_CODE string = c.Param("id")

	fmt.Println(STD_CODE)

	qf, rowsAffected, err := h.officerServices.UpdateQualification(STD_CODE, requestBody.STATUS, requestBody.DESCRIPTION)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "ไม่สามารถดำเนินการปรับปรุงสถานะ " + requestBody.STATUS + " การยื่นขอเอกสารของนักศึกษา " + STD_CODE + " ได้."})
		c.Abort()
		return
	}

	if rowsAffected < 1 {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "ไม่สามารถดำเนินการปรับปรุงสถานะรายการจาก " + qf.STATUS + " เป็น " + requestBody.STATUS + " ได้."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": qf, "message": "ดำเนินการปรับปรุงสถานะ " + requestBody.STATUS + " การยื่นขอเอกสารของนักศึกษา " + STD_CODE + " สำเร็จ."})

}
