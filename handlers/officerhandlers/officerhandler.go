package officerhandlers

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"errors"
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

func getLineNumber() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

// Helper function to get the file name where the function is called
func getFileName() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

func ErrTokenLogin(c *gin.Context) {
	err := errors.New("ไม่พบ token login.")
	c.Error(err)
	c.Set("line", handlers.GetLineNumber())
	c.Set("file", handlers.GetFileName())
	c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ token login."})
	c.Abort()
}

func ErrTokenClaim(c *gin.Context) {
	err := errors.New("ไม่พบ claims user.")
	c.Error(err)
	c.Set("line", handlers.GetLineNumber())
	c.Set("file", handlers.GetFileName())
	c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "ไม่พบ claims user."})
	c.Abort()
}

func ErrRoleBachelor(c *gin.Context) {
	err := errors.New("สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้")
	c.Error(err)
	c.Set("line", handlers.GetLineNumber())
	c.Set("file", handlers.GetFileName())
	c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้"})
	c.Abort()
}

func ErrValidateRequest(c *gin.Context) {
	err := errors.New("โปรดระบุค่าให้ถูกต้อง")
	c.Error(err)
	c.Set("line", handlers.GetLineNumber())
	c.Set("file", handlers.GetFileName())
	c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "ระบุค่า parameter ไม่ถูกต้อง."})
	c.Abort()
}
