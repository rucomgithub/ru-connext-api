package handlers

import (
	"net/http"
	"strconv"

	"RU-Smart-Workspace/ru-smart-api/application/usecases"
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"RU-Smart-Workspace/ru-smart-api/handlers"

	"github.com/gin-gonic/gin"
)

type PublicationHandler struct {
	useCase *usecases.PublicationUseCase
}

func NewPublicationHandler(useCase *usecases.PublicationUseCase) *PublicationHandler {
	return &PublicationHandler{useCase: useCase}
}

func (h *PublicationHandler) CreatePublication(c *gin.Context) {
	var publication entities.Publication

	if err := c.ShouldBindJSON(&publication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.CreatePublication(&publication); err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "ไม่สามารถบันทึกแบบ ว.9 ของนักศึกษารหัส " + publication.StudentCode + " ในระบบ. " + err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, publication)
}

func (h *PublicationHandler) GetPublication(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "ป้อนรหัสนักศึกษาไม่ถูกต้อง"})
		c.Abort()
		return
	}

	publication, err := h.useCase.GetPublication(id)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "ไม่สามารถดึงข้อมูล ว.9 ของนักศึกษารหัส " + publication.StudentCode + " ในระบบ."})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, publication)
}

func (h *PublicationHandler) UpdatePublication(c *gin.Context) {
	student_code := c.Param("id")
	if student_code == "" {
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "ป้อนรหัสนักศึกษาไม่ถูกต้อง"})
		c.Abort()
		return
	}

	var publication entities.Publication
	if err := c.ShouldBindJSON(&publication); err != nil {
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "ป้อนข้อมูลไม่ถูกต้อง"})
		c.Abort()
		return
	}

	publication.StudentCode = student_code
	if err := h.useCase.UpdatePublication(&publication); err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "ไม่สามารถแก้ไขแบบ ว.9 ของนักศึกษารหัส " + publication.StudentCode + " ในระบบ. " + err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, publication)
}

func (h *PublicationHandler) DeletePublication(c *gin.Context) {
	student_code := c.Param("id")
	if student_code == "" {
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "ป้อนรหัสนักศึกษาไม่ถูกต้อง"})
		c.Abort()
		return
	}

	if err := h.useCase.DeletePublication(student_code); err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "ไม่สามารถลบรายการแบบ ว.9 ของนักศึกษารหัส " + student_code + " ในระบบ. " + err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Publication deleted successfully"})
}

func (h *PublicationHandler) ListPublications(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "25")

	offset, _ := strconv.Atoi(offsetStr)
	limit, _ := strconv.Atoi(limitStr)

	publications, err := h.useCase.ListPublications(offset, limit)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "ไม่พบรายการแบบ ว.9 ของนักศึกษาในระบบ."})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"publications": publications})
}
