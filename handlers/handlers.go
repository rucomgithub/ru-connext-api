package handlers

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case errs.AppError:
		c.JSON(e.Code, gin.H{"message": e.Message})
	case error:
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
}

func GetLineNumber() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

// Helper function to get the file name where the function is called
func GetFileName() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}
