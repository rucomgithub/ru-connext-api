package masterhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"errors"
	"net/http"
	_ "net/url"

	"github.com/gin-gonic/gin"
)

func (h *studentHandlers) GetRegisterAll(c *gin.Context) {

	token, err := middlewares.GetHeaderAuthorization(c)

	if err != nil {
		err = errors.New("ไม่พบ token login.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ token login."})
		c.Abort()
		return
	}

	claim, err := middlewares.GetClaims(token)

	if err != nil {
		err = errors.New("ไม่พบ claims user.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ claims user."})
		c.Abort()
		return
	}

	role := claim.Role

	if role == "Bachelor" {
		err = errors.New("สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้."})
		c.Abort()
		return
	}

	std_code := claim.StudentCode

	studentRegisterResponse, err := h.studentService.GetRegisterAll(std_code)
	if err != nil {
		err = errors.New("ไม่พบข้อมูลลงทะเบียนนักศึกษา " + std_code + ".")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ไม่พบข้อมูลลงทะเบียนนักศึกษา " + std_code + "."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, studentRegisterResponse)

}

func (h *studentHandlers) GetRegisterFeeAll(c *gin.Context) {

	token, err := middlewares.GetHeaderAuthorization(c)

	if err != nil {
		err = errors.New("ไม่พบ token login.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ token login."})
		c.Abort()
		return
	}

	claim, err := middlewares.GetClaims(token)

	if err != nil {
		err = errors.New("ไม่พบ claims user.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ claims user."})
		c.Abort()
		return
	}

	role := claim.Role

	if role == "Bachelor" {
		err = errors.New("สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้."})
		c.Abort()
		return
	}

	std_code := claim.StudentCode

	studentRegisterFeeResponse, err := h.studentService.GetRegisterFeeAll(std_code,role)
	if err != nil {
		err = errors.New("ไม่พบข้อมูลค่าธรรมเนียมการลงทะเบียนนักศึกษา " + std_code + ".")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ไม่พบข้อมูลค่าธรรมเนียมการลงทะเบียนนักศึกษา " + std_code + "."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, studentRegisterFeeResponse)

}

func (h *studentHandlers) GetRegisterByYear(c *gin.Context) {

	token, err := middlewares.GetHeaderAuthorization(c)

	if err != nil {
		err = errors.New("ไม่พบ token login.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ token login."})
		c.Abort()
		return
	}

	claim, err := middlewares.GetClaims(token)

	if err != nil {
		err = errors.New("ไม่พบ claims user.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ claims user."})
		c.Abort()
		return
	}

	role := claim.Role

	if role == "Bachelor" {
		err = errors.New("สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้."})
		c.Abort()
		return
	}

	std_code := claim.StudentCode

	var year string = c.Param("year")

	if year == "" {
		err = errors.New("ไม่ระบุปีการศึกษา.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ไม่ระบุปีการศึกษา."})
		c.Abort()

		return
	}

	studentProfileResponse, err := h.studentService.GetRegisterByYear(std_code, year)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ไม่พบข้อมูลลงทะเบียนนักศึกษา " + std_code + " ประจำปีการศึกษา " + year})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, studentProfileResponse)

}
