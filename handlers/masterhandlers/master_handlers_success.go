package masterhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"errors"
	"fmt"
	"net/http"
	_ "net/url"

	"github.com/gin-gonic/gin"
)

func (h *studentHandlers) GetStudentSuccess(c *gin.Context) {

	token, err := middlewares.GetHeaderAuthorization(c)

	fmt.Println(token)

	if err != nil {
		err = errors.New("ไม่พบ token login.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ token login."})
		c.Abort()
		return
	}

	fmt.Println(token)

	claim, err := middlewares.GetClaims(token)

	if err != nil {
		err = errors.New("ไม่พบ claims user." + err.Error())
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ claims user."})
		c.Abort()
		return
	}

	role := claim.Role

	fmt.Println(role)

	if role == "Bachelor" {
		err = errors.New("สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้...")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้."})
		c.Abort()
		return
	}

	std_code := claim.StudentCode

	studentSuccessResponse, err := h.studentService.GetStudentSuccess(std_code)
	if err != nil {
		err = errors.New("ไม่พบข้อมูลสำเร็จการศึกษา " + std_code + ".")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ไม่พบข้อมูลสำเร็จการศึกษา " + std_code + "."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, studentSuccessResponse)

}

func (h *studentHandlers) GetStudentRequestSuccess(c *gin.Context) {

	token, err := middlewares.GetHeaderAuthorization(c)

	fmt.Println(token)

	if err != nil {
		err = errors.New("ไม่พบ token login.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ token login."})
		c.Abort()
		return
	}

	fmt.Println(token)

	claim, err := middlewares.GetClaims(token)

	if err != nil {
		err = errors.New("ไม่พบ claims user." + err.Error())
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ claims user."})
		c.Abort()
		return
	}

	role := claim.Role

	fmt.Println(role)

	if role == "Bachelor" {
		err = errors.New("สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้...")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้."})
		c.Abort()
		return
	}

	std_code := claim.StudentCode

	studentSuccessResponse, err := h.studentService.GetStudentRequestSuccess(std_code)
	if err != nil {
		err = errors.New("ไม่พบข้อมูลขอจบสำเร็จการศึกษา " + std_code + "." + err.Error())
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ไม่พบข้อมูลขอจบสำเร็จการศึกษา " + std_code + "."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, studentSuccessResponse)

}

func (h *studentHandlers) GetStudentSuccessById(c *gin.Context) {

	std_code := c.Param("id")

	if std_code == "" {
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบรหัสนักศึกษา."})
		c.Abort()
		return
	}

	studentSuccessResponse, err := h.studentService.GetStudentSuccess(std_code)
	if err != nil {
		err = errors.New("ไม่พบข้อมูลสำเร็จการศึกษา " + std_code + ".")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ไม่พบข้อมูลสำเร็จการศึกษา " + std_code + "."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, studentSuccessResponse)

}

func (h *studentHandlers) GetStudentSuccessCheck(c *gin.Context) {
	token := c.Param("id")

	if token == "" {
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ token certificate."})
		c.Abort()
		return
	}

	studentSuccessResponse, err := h.studentService.GetStudentSuccessCheck(token)
	if err != nil {
		//err = errors.New("ไม่พบข้อมูลสำเร็จการศึกษา " + err.Error() + std_code + ".")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, studentSuccessResponse)

}
