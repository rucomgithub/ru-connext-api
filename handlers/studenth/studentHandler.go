package studenth

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"RU-Smart-Workspace/ru-smart-api/services/students"
	"net/http"

	"github.com/gin-gonic/gin"
)

type studentHandlers struct {
	studentService students.StudentServicesInterface
}

func NewStudentHandlers(studentService students.StudentServicesInterface) studentHandlers {
	return studentHandlers{studentService: studentService}
}

func (h *studentHandlers) AuthenticationTest(c *gin.Context) {

	var requestBody students.AuthenTestPlayload

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message":"Authorization fail becourse content type not json format."})
		c.Abort()
		return
	}

	tokenResponse, err := h.studentService.Authentication(requestBody.Std_code)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, tokenResponse)
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, tokenResponse)

}

func (h *studentHandlers) Authentication(c *gin.Context) {

	var requestBody students.AuthenPlayload

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message":"Authorization fail becourse content type not json format."})
		c.Abort()
		return
	}

	tokenResponse, err := h.studentService.Authentication(requestBody.Std_code)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, tokenResponse)
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, tokenResponse)

}


func (h *studentHandlers) AuthenticationRedirect(c *gin.Context) {

	var requestBody students.AuthenPlayloadRedirect

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message":"Authorization fail becourse content type not json format."})
		c.Abort()
		return
	}

	tokenResponse, err := h.studentService.AuthenticationRedirect(requestBody.Std_code, requestBody.Access_token)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, tokenResponse)
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, tokenResponse)

}

func (h *studentHandlers) RefreshAuthentication(c *gin.Context) {

	var requestBody students.AuthenPlayload

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message":"Authorization fail becourse content type not json format."})
		c.Abort()
		return
	}

	tokenRespone, err := h.studentService.RefreshAuthentication(requestBody.Refresh_token,requestBody.Std_code)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, tokenRespone)
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, tokenRespone)

}

func (h *studentHandlers) Unauthorization(c *gin.Context) {
		token, err := middlewares.GetHeaderAuthorization(c)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Authorization key in header not found"})
			c.Abort()
			return
		}

		// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
		isToken := h.studentService.Unauthorization(token)
		if !isToken {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"Authorization falil because of timeout."})
			c.Abort()
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message":"Unauthorization successfuly."})

}

func (h *studentHandlers) GetStudentProfile(c *gin.Context) {

	STD_CODE := c.Param("std_code")
	studentProfileResponse, err := h.studentService.GetStudentProfile(STD_CODE)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, studentProfileResponse)

}

func (h *studentHandlers) GetRegister(c *gin.Context) {
	
	var payload students.RegisterPlayload
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message":"Authorization fail becourse content type not json format."})
		c.Abort()
		return
	}

	registerResponse, err := h.studentService.GetRegister(payload.Std_code, payload.Course_year, payload.Course_semester)
	if err != nil {
		c.IndentedJSON(http.StatusNoContent, registerResponse)
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, registerResponse)
}
