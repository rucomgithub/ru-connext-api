package studenth

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"RU-Smart-Workspace/ru-smart-api/services/students"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
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
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Authorization fail becourse content type not json format."})
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
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Authorization fail becourse content type not json format."})
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
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Authorization fail becourse content type not json format."})
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
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Authorization fail becourse content type not json format."})
		c.Abort()
		return
	}

	tokenRespone, err := h.studentService.RefreshAuthentication(requestBody.Refresh_token, requestBody.Std_code)
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
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Authorization falil because of timeout."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Unauthorization successfuly."})

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
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Authorization fail becourse content type not json format."})
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

func (h *studentHandlers) GetImageProfile(c *gin.Context) {

	var token string
	const BEARER_SCHEMA = "Bearer "
	AUTH_HEADER := c.GetHeader("Authorization")

	if len(AUTH_HEADER) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	if strings.HasPrefix(AUTH_HEADER, BEARER_SCHEMA) {
		token = AUTH_HEADER[len(BEARER_SCHEMA):]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is not has"})
		return
	}

	url := "http://10.2.1.155:9100/student/image"

	client := &http.Client{
		Timeout: 60 * time.Second, // Set a higher timeout value
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Set request error"})
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)

	response, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get image error" + err.Error()})
		return
	}
	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format"})
		return
	}

	// Decode the image
	var img image.Image
	switch contentType {
	case "image/jpeg":
		img, err = jpeg.Decode(response.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Get decode jpeg error" + err.Error()})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported image format"})
		return
	}
	// Resize the image to a desired width and height
	resizedImg := resize.Resize(200, 0, img, resize.Lanczos3) // Adjust the width as per your requirement, set height to 0 to maintain aspect ratio

	// Create a new image with the same dimensions as the original image
	result := image.NewRGBA(resizedImg.Bounds())

	// Copy the original image to the new image
	draw.Draw(result, result.Bounds(), resizedImg, image.Point{0, 0}, draw.Src)

	// Set the text properties
	fontPath := "/app/fonts/Kanit-LightItalic.ttf" // Replace with the path to your desired font file
	fontSize := 10.0
	textColor := color.Black
	text := "@Ramkhamaeng University"

	// Load the font file
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading font file:" + err.Error()})
		return
	}
	font, err := truetype.Parse(fontData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing font:" + err.Error()})
		return
	}

	// Create the freetype context
	fctx := freetype.NewContext()
	fctx.SetDst(result)
	fctx.SetSrc(image.NewUniform(textColor))
	fctx.SetClip(result.Bounds())
	fctx.SetFont(font)
	fctx.SetFontSize(fontSize)

	// Calculate the text dimensions

	pt := freetype.Pt(10, 190)
	_, err = fctx.DrawString(text, pt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error drawing text: " + err.Error()})
		return
	}

	// Create a new in-memory buffer to store the resized image
	outputImg := new(bytes.Buffer)
	if err := jpeg.Encode(outputImg, result, nil); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get resize error" + err.Error()})
		return
	}

	// Set the appropriate headers for the response
	c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Disposition", "attachment; filename=resized_image.jpg") // Set the desired filename

	// Write the resized image to the response
	c.Data(http.StatusOK, "image/jpeg", outputImg.Bytes())
}

func (h *studentHandlers) GetPhoto(c *gin.Context) {

	ID_TOKEN, err := middlewares.GetHeaderAuthorization(c)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "authorization key in header not found"})
		c.Abort()
		return
	}

	url := "http://10.2.1.155:9100/student/photo"

	client := &http.Client{
		Timeout: 60 * time.Second, // Set a higher timeout value
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Set request error"})
		return
	}

	req.Header.Set("Authorization", "Bearer "+ID_TOKEN)

	response, err := client.Do(req)
	if err != nil {
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Get decode jpeg error" + err.Error()})
			return
		}
	case "image/png":
		img, err = jpeg.Decode(response.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Get decode png error" + err.Error()})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported image format"})
		return
	}

	// outputImg := new(bytes.Buffer)
	outputImg := bytes.NewBuffer(nil)

	if err := jpeg.Encode(outputImg, img, nil); err != nil {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Set request error"})
		return
	}

	req.Header.Set("Authorization", "Bearer "+id)

	response, err := client.Do(req)
	if err != nil {
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Get decode jpeg error" + err.Error()})
			return
		}
	case "image/png":
		img, err = jpeg.Decode(response.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Get decode png error" + err.Error()})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported image format"})
		return
	}

	// outputImg := new(bytes.Buffer)
	outputImg := bytes.NewBuffer(nil)

	if err := jpeg.Encode(outputImg, img, nil); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get resize error" + err.Error()})
		return
	}

	c.Data(http.StatusOK, contentType, outputImg.Bytes())
}

func (h *studentHandlers) GetPhotoAOD(c *gin.Context) {

	ID_TOKEN, err := middlewares.GetHeaderAuthorization(c)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "authorization key in header not found"})
		c.Abort()
		return
	}

	url := "http://10.2.1.155:9100/student/image"

	client := &http.Client{
		Timeout: 60 * time.Second, // Set a higher timeout value
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Set request error"})
		return
	}

	req.Header.Set("Authorization", "Bearer "+ID_TOKEN)

	response, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get image error" + err.Error()})
		return
	}
	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format" + contentType + response.Status})
		return
	}

	// Decode the image
	var img image.Image
	switch contentType {
	case "image/jpeg":
		img, err = jpeg.Decode(response.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Get decode jpeg error" + err.Error()})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported image format"})
		return
	}

	// Create a new in-memory buffer to store the resized image
	outputImg := new(bytes.Buffer)
	//outputImg := bytes.NewBuffer(nil)

	if err := jpeg.Encode(outputImg, img, nil); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get resize error" + err.Error()})
		return
	}

	// Set the appropriate headers for the image response
	// c.Header("Content-Type", "image/jpeg")

	// // Write the image data as the response body
	// c.Writer.WriteHeader(http.StatusOK)
	// c.Writer.Write(outputImg.Bytes())
	c.Data(http.StatusOK, "image/jpeg", outputImg.Bytes())
}
