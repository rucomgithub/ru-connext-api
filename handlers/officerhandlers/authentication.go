package officerhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/services/officerservices"
	"RU-Smart-Workspace/ru-smart-api/services/students"
	"net/http"
	"time"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func (h *officerHandlers) Authentication(c *gin.Context) {
	var requestBody officerservices.AuthenRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", getLineNumber())
		c.Set("file", getFileName())
		handleError(c, err)
		return
	}

	authenResponse, err := h.officerServices.AuthenticationOfficer(requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", getLineNumber())
		c.Set("file", getFileName())
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, authenResponse)

}

func (h *officerHandlers) RefreshAuthentication(c *gin.Context) {

	var requestBody students.AuthenPlayload

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", getLineNumber())
		c.Set("file", getFileName())
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Authorization fail becourse content type not json format."})
		c.Abort()
		return
	}

	tokenRespone, err := h.officerServices.RefreshAuthenticationOfficer(requestBody.Refresh_token)
	if err != nil {

		c.Set("line", getLineNumber())
		c.Set("file", getFileName())
		c.IndentedJSON(http.StatusForbidden, tokenRespone)
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, tokenRespone)

}

func (h *officerHandlers) GetPhoto(c *gin.Context) {

	var accessToken string = c.Param("id")

	timeout := 10 * time.Second
	client := &http.Client{
		Timeout: timeout,
	}

	surl := "https://graph.microsoft.com/v1.0/me/photos/48x48/$value"

	req, err := http.NewRequest("GET", surl, nil)
	if err != nil {
		c.Error(err)
		c.Set("line", getLineNumber())
		c.Set("file", getFileName())
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Cannot create request."})
		c.Abort()
		return
	}

	fmt.Println("AccessToken:", accessToken)
	req.Header.Set("Authorization", accessToken)

	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Set("line", getLineNumber())
		c.Set("file", getFileName())
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Cannot get response."})
		c.Abort()
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.Set("line", getLineNumber())
		c.Set("file", getFileName())
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "No photo found."})
		c.Abort()
		return
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Error(err)
		c.Set("line", getLineNumber())
		c.Set("file", getFileName())
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Cannot read photo data."})
		c.Abort()
		return
	}

	c.Data(http.StatusOK, "image/png", data)
}

