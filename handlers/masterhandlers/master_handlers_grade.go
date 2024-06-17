package masterhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"errors"
	"net/http"
	_ "net/url"

	"github.com/gin-gonic/gin"
)

func (h *studentHandlers) GetGradeAll(c *gin.Context) {

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

	studentProfileResponse, err := h.studentService.GetGradeAll(STD_CODE)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ไม่พบข้อมูลผลการเรียนนักศึกษา " + STD_CODE + "."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, studentProfileResponse)

}

func (h *studentHandlers) GetGradeByYear(c *gin.Context) {

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

	std_code := claim.StudentCode

	var year string = c.Param("year")

	if year == "" {
		handleError(c, errors.New("โปรดระบุปีการศึกษา"))
		return
	}

	studentProfileResponse, err := h.studentService.GetGradeByYear(std_code, year)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ไม่พบข้อมูลผลการเรียนนักศึกษา " + std_code + " ประจำปีการศึกษา " + year})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, studentProfileResponse)

}
