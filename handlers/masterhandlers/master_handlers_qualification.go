package masterhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"fmt"
	"net/http"
	_ "net/url"

	"github.com/gin-gonic/gin"
)

func (h *studentHandlers) GetQualification(c *gin.Context) {

	token, err := middlewares.GetHeaderAuthorization(c)

	if err != nil {
		ErrTokenLogin(c)
		return
	}

	claim, err := middlewares.GetClaims(token)

	if err != nil {
		ErrTokenClaim(c)
		return
	}

	role := claim.Role

	if role == "Bachelor" {
		ErrRoleBachelor(c)
		return
	}

	STD_CODE := claim.StudentCode

	fmt.Println(claim.StudentCode)

	qf, err := h.studentService.GetQualification(STD_CODE)

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

func (h *studentHandlers) AddQualification(c *gin.Context) {

	token, err := middlewares.GetHeaderAuthorization(c)

	if err != nil {
		ErrTokenLogin(c)
		return
	}

	claim, err := middlewares.GetClaims(token)

	if err != nil {
		ErrTokenClaim(c)
		return
	}

	role := claim.Role

	if role == "Bachelor" {
		ErrRoleBachelor(c)
		return
	}

	STD_CODE := claim.StudentCode

	fmt.Println(claim.StudentCode)

	qf, err := h.studentService.AddQualification(STD_CODE)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "รหัสนักศึกษา " + STD_CODE + " ยื่นขอเอกสารซ้ำกันในระบบ."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": qf, "message": "รหัสนักศึกษา " + STD_CODE + " ยื่นขอเอกสารสำเร็จ."})

}
