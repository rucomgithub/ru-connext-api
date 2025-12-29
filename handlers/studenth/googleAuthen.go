package studenth

import (
    "log"
    "net/http"
	"RU-Smart-Workspace/ru-smart-api/services/students"
	"RU-Smart-Workspace/ru-smart-api/handlers"
    "github.com/gin-gonic/gin"
)

func (h *studentHandlers) AuthorizationGoogle(c *gin.Context) {
    var req students.AuthenPlayload

    if err := c.ShouldBindJSON(&req); err != nil { 
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request",
            "message": "กรุณาระบุข้อมูลให้ครบถ้วน",
        })
        return
    }

    // Verify Google token
    tokenInfo, err := h.studentService.VerifyToken(c.Request.Context(), req.Refresh_token)
    if err != nil {
        log.Printf("Token verification failed: %v", err)
        c.JSON(http.StatusUnauthorized, gin.H{
            "error":   "Token verification failed",
            "message": "ไม่สามารถยืนยันตัวตนได้",
        })
        return
    }

    // Verify student ID matches email
    extractedID := h.studentService.ExtractStudentID(tokenInfo.Email)
    if extractedID != req.Std_code {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Student ID mismatch",
            "message": "รหัสนักศึกษาไม่ตรงกับอีเมล",
        })
        return
    }

	tokenResponse, err := h.studentService.Authentication(extractedID)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnprocessableEntity, tokenResponse)
		c.Abort()
		return
	}

    c.JSON(http.StatusOK, tokenResponse)
}