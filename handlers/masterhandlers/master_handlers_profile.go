package masterhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"net/http"
	_ "net/url"

	"github.com/gin-gonic/gin"
)

func (h *studentHandlers) GetStudentProfile(c *gin.Context) {

	token, err := middlewares.GetHeaderAuthorization(c)

	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ไม่พบ token login."})
		c.Abort()
		return
	}

	claim, err := middlewares.GetClaims(token)

	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ไม่พบ claims user."})
		c.Abort()
		return
	}

	STD_CODE := claim.StudentCode

	studentProfileResponse, err := h.studentService.GetStudentProfile(STD_CODE)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ไม่พบข้อมูลประวัตินักศึกษา."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, studentProfileResponse)

}
