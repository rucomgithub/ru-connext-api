package handlers

import (
	"net/http"
	"strconv"
	"log"
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"RU-Smart-Workspace/ru-smart-api/domain/services"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type journalHandler struct {
	thesisJournalService services.ThesisJournalService
}

func NewJournalHandler(thesisJournalService services.ThesisJournalService) *journalHandler {
	return &journalHandler{
		thesisJournalService: thesisJournalService,
	}
}

func (h *journalHandler) CreateJournal(c *gin.Context) {
	var thesisJournal entities.ThesisJournal
	if err := c.ShouldBindJSON(&thesisJournal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.thesisJournalService.CreateThesisJournal(c.Request.Context(), &thesisJournal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Journal created successfully",
		"data":    thesisJournal,
	})
}

func (h *journalHandler) GetJournal(c *gin.Context) {
	id := c.Param("id")

	thesisJournal, err := h.thesisJournalService.GetThesisJournal(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Journal retrieved successfully",
		"data":    thesisJournal,
	})
}

func (h *journalHandler) GetJournalMaster(c *gin.Context) {
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

	thesisJournal, err := h.thesisJournalService.GetThesisJournal(c.Request.Context(), std_code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Journal retrieved successfully",
		"data":    thesisJournal,
	})
}

func (h *journalHandler) GetJournalByStudentID(c *gin.Context) {
	studentID := c.Param("studentId")

	thesisJournal, err := h.thesisJournalService.GetThesisJournalByStudentID(c.Request.Context(), studentID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Journal retrieved successfully",
		"data":    thesisJournal,
	})
}

func (h *journalHandler) UpdateJournal(c *gin.Context) {
	id := c.Param("id")

	var thesisJournal entities.ThesisJournal
	if err := c.ShouldBindJSON(&thesisJournal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	thesisJournal.StudentID = id
	log.Print(thesisJournal.JournalPublication)
	if err := h.thesisJournalService.UpdateThesisJournal(c.Request.Context(), &thesisJournal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Journal updated successfully",
		"data":    thesisJournal,
	})
}

func (h *journalHandler) UpdateJournalMaster(c *gin.Context) {
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
	
	var thesisJournal entities.ThesisJournal
	if err := c.ShouldBindJSON(&thesisJournal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	thesisJournal.StudentID = claim.StudentCode
	log.Print(thesisJournal.JournalPublication)
	if err := h.thesisJournalService.UpdateThesisJournal(c.Request.Context(), &thesisJournal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Journal updated successfully",
		"data":    thesisJournal,
	})
}

func (h *journalHandler) DeleteJoural(c *gin.Context) {
	id := c.Param("id")

	if err := h.thesisJournalService.DeleteThesisJournal(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "check error" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Journal deleted successfully",
	})
}

func (h *journalHandler) ListJournals(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10000")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10000
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	students, err := h.thesisJournalService.ListThesisJournals(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "list journal retrieved successfully",
		"data":    students,
	})
}
