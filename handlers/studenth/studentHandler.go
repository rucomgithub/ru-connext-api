package studenth

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"RU-Smart-Workspace/ru-smart-api/services/students"
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	_ "net/url"
	"strings"
	"time"

	"github.com/spf13/viper"

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
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Authorization fail becourse content type not json format."})
		c.Abort()
		return
	}

	tokenResponse, err := h.studentService.Authentication(requestBody.Std_code)
	if err != nil {
		c.Error(errors.New(err.Error() + ", " + requestBody.Std_code))
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnprocessableEntity, tokenResponse)
		c.Abort()
		return
	}
	c.IndentedJSON(http.StatusOK, tokenResponse)
}

func (h *studentHandlers) AuthenticationService(c *gin.Context) {

	var requestBody students.AuthenServicePlayload

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Authorization fail becourse content type not json format."})
		c.Abort()
		return
	}

	tokenResponse, err := h.studentService.AuthenticationService(requestBody.ServiveId)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
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
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Authorization fail becourse content type not json format."})
		c.Abort()
		return
	}

	tokenResponse, err := h.studentService.Authentication(requestBody.Std_code)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
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
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Authorization fail becourse content type not json format."})
		c.Abort()
		return
	}

	tokenResponse, err := h.studentService.AuthenticationRedirect(requestBody.Std_code, requestBody.Access_token)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
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
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Authorization fail becourse content type not json format."})
		c.Abort()
		return
	}

	tokenRespone, err := h.studentService.RefreshAuthentication(requestBody.Refresh_token)
	if err != nil {

		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusForbidden, tokenRespone)
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, tokenRespone)

}

func (h *studentHandlers) Unauthorization(c *gin.Context) {
	token, err := middlewares.GetHeaderAuthorization(c)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Authorization key in header not found"})
		c.Abort()
		return
	}

	// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
	isToken := h.studentService.Unauthorization(token)
	if !isToken {
		err := errors.New("Authorization falil because of timeout...")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Authorization falil because of timeout..."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Unauthorization successfuly."})
}

func (h *studentHandlers) CheckToken(c *gin.Context) {
	token := c.Param("token")
	// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
	tokenRec, err := h.studentService.CheckToken(token)
	if err != nil {
		err := errors.New("Authorization falil because of timeout...")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName()) // Register the error
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Authorization falil because of timeout..."})
		c.Abort()
		return
	}

	studentProfileResponse, err := h.studentService.GetStudentProfile(tokenRec.StudentCode)
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

func (h *studentHandlers) ExistsToken(c *gin.Context) {
	token, err := middlewares.GetHeaderAuthorization(c)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Authorization key in header not found"})
		c.Abort()
		return
	}

	// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
	isToken := h.studentService.CheckExistsToken(token)
	if !isToken {
		err := errors.New("Authorization falil because of timeout...")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName()) // Register the error
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Authorization falil because of timeout..."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Authorization successfuly and Exists Token."})
}

func (h *studentHandlers) GetStudentProfile(c *gin.Context) {

	STD_CODE := c.Param("std_code")
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

func (h *studentHandlers) GetRegister(c *gin.Context) {

	var payload students.RegisterPlayload
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Authorization fail becourse content type not json format."})
		c.Abort()
		return
	}

	registerResponse, err := h.studentService.GetRegister(payload.Std_code, payload.Course_year, payload.Course_semester)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusNoContent, gin.H{"message": "ไม่พบข้อมูลลงทะเบียนของนักศึกษา."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, registerResponse)
}

func (h *studentHandlers) GetPhoto(c *gin.Context) {
	service_token := viper.GetString("token.eservice")

	ID_TOKEN, err := middlewares.GetHeaderAuthorization(c)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "authorization key in header not found"})
		c.Abort()
		return
	}

	url := "http://10.2.1.155:9100/student/photo"

	client := &http.Client{
		Timeout: 60 * time.Second, // Set a higher timeout value
	}

	req, err := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("id_token", ID_TOKEN)
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Set request error"})
		return
	}

	req.Header.Set("Authorization", "Bearer "+service_token)

	response, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get image error"})
		return
	}
	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		err := errors.New("Invalid image format.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName()) // Register the error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format."})
		return
	}

	fmt.Println(contentType)

	// Decode the image
	var img image.Image
	switch contentType {
	case "image/jpeg":
		img, err = jpeg.Decode(response.Body)
		if err != nil {
			c.Error(err)
			c.Set("line", handlers.GetLineNumber())
			c.Set("file", handlers.GetFileName())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Get decode jpeg error"})
			return
		}
	case "image/png":
		img, err = jpeg.Decode(response.Body)
		if err != nil {
			c.Error(err)
			c.Set("line", handlers.GetLineNumber())
			c.Set("file", handlers.GetFileName())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Get decode png error"})
			return
		}
	default:
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported image format"})
		return
	}

	// outputImg := new(bytes.Buffer)
	outputImg := bytes.NewBuffer(nil)

	if err := jpeg.Encode(outputImg, img, nil); err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get resize error" + err.Error()})
		return
	}

	c.Data(http.StatusOK, contentType, outputImg.Bytes())
}

func (h *studentHandlers) GetPhotoGraduate(c *gin.Context) {
	service_token := viper.GetString("token.eservice")

	ID_TOKEN, err := middlewares.GetHeaderAuthorization(c)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "authorization key in header not found"})
		c.Abort()
		return
	}

	url := "http://10.2.1.155:9100/student/photograduate"

	client := &http.Client{
		Timeout: 60 * time.Second, // Set a higher timeout value
	}

	req, err := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("id_token", ID_TOKEN)
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Set request error"})
		return
	}

	req.Header.Set("Authorization", "Bearer "+service_token)

	response, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get image error"})
		return
	}
	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		err := errors.New("Invalid image format.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName()) // Register the error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format."})
		return
	}

	fmt.Println(contentType)

	// Decode the image
	var img image.Image
	switch contentType {
	case "image/jpeg":
		img, err = jpeg.Decode(response.Body)
		if err != nil {
			c.Error(err)
			c.Set("line", handlers.GetLineNumber())
			c.Set("file", handlers.GetFileName())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Get decode jpeg error"})
			return
		}
	case "image/png":
		img, err = jpeg.Decode(response.Body)
		if err != nil {
			c.Error(err)
			c.Set("line", handlers.GetLineNumber())
			c.Set("file", handlers.GetFileName())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Get decode png error"})
			return
		}
	default:
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported image format"})
		return
	}

	// outputImg := new(bytes.Buffer)
	outputImg := bytes.NewBuffer(nil)

	if err := jpeg.Encode(outputImg, img, nil); err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get resize error" + err.Error()})
		return
	}

	c.Data(http.StatusOK, contentType, outputImg.Bytes())
}

func (h *studentHandlers) GetPhotoGraduateById(c *gin.Context) {
	service_token := viper.GetString("token.eservice")

	ID_TOKEN := c.Param("id")

	if ID_TOKEN == "" {
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "authorization key on request not found."})
		c.Abort()
		return
	}

	url := "http://10.2.1.155:9100/student/photograduate"

	client := &http.Client{
		Timeout: 60 * time.Second, // Set a higher timeout value
	}

	req, err := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("id_token", ID_TOKEN)
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Set request error"})
		return
	}

	req.Header.Set("Authorization", "Bearer "+service_token)

	response, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get image error"})
		return
	}
	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		err := errors.New("Invalid image format.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName()) // Register the error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format."})
		return
	}

	fmt.Println(contentType)

	// Decode the image
	var img image.Image
	switch contentType {
	case "image/jpeg":
		img, err = jpeg.Decode(response.Body)
		if err != nil {
			c.Error(err)
			c.Set("line", handlers.GetLineNumber())
			c.Set("file", handlers.GetFileName())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Get decode jpeg error"})
			return
		}
	case "image/png":
		img, err = jpeg.Decode(response.Body)
		if err != nil {
			c.Error(err)
			c.Set("line", handlers.GetLineNumber())
			c.Set("file", handlers.GetFileName())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Get decode png error"})
			return
		}
	default:
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported image format"})
		return
	}

	// outputImg := new(bytes.Buffer)
	outputImg := bytes.NewBuffer(nil)

	if err := jpeg.Encode(outputImg, img, nil); err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get resize error" + err.Error()})
		return
	}

	c.Data(http.StatusOK, contentType, outputImg.Bytes())
}

func (h *studentHandlers) GetPhotoById(c *gin.Context) {

	id := c.Param("id")

	url := "http://10.2.1.155:9100/student/photo"

	client := &http.Client{
		Timeout: 60 * time.Second, // Set a higher timeout value
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Set request error"})
		return
	}

	req.Header.Set("Authorization", "Bearer "+id)

	response, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get image error" + err.Error()})
		return
	}
	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format"})
		return
	}

	fmt.Println(contentType)

	// Decode the image
	var img image.Image
	switch contentType {
	case "image/jpeg":
		img, err = jpeg.Decode(response.Body)
		if err != nil {
			c.Error(err)
			c.Set("line", handlers.GetLineNumber())
			c.Set("file", handlers.GetFileName())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Get decode jpeg error" + err.Error()})
			return
		}
	case "image/png":
		img, err = jpeg.Decode(response.Body)
		if err != nil {
			c.Error(err)
			c.Set("line", handlers.GetLineNumber())
			c.Set("file", handlers.GetFileName())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Get decode png error" + err.Error()})
			return
		}
	default:
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported image format"})
		return
	}

	// outputImg := new(bytes.Buffer)
	outputImg := bytes.NewBuffer(nil)

	if err := jpeg.Encode(outputImg, img, nil); err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get resize error" + err.Error()})
		return
	}

	c.Data(http.StatusOK, contentType, outputImg.Bytes())
}
