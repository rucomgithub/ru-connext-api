package masterhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
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

func (h *studentHandlers) GetPhotoGraduateByStudentCode(c *gin.Context) {
	service_token := viper.GetString("token.eservice")

	std_code := c.Param("id")

	token, err := middlewares.GenerateToken(std_code, "admin", h.redis_cache)

	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Set request error"})
		return
	}

	url := "http://10.2.1.155:9100/student/photograduate"

	client := &http.Client{
		Timeout: 60 * time.Second, // Set a higher timeout value
	}

	req, err := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("id_token", token.AccessToken)
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