package handlers

import (
	"net/http"
	"strconv"

	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"RU-Smart-Workspace/ru-smart-api/domain/services"

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

	// // Parse conference date if present
	// if thesisJournal.ConferencePresentation != nil {
	// 	// The JSON might have the date as string, convert it
	// 	if dateStr, ok := c.GetPostForm("conferencePresentation.conferenceDate"); ok {
	// 		if date, err := time.Parse("2006-01-02", dateStr); err == nil {
	// 			thesisJournal.ConferencePresentation.ConferenceDate = date
	// 		}
	// 	}
	// }

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

func (h *journalHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")

	var thesisJournal entities.ThesisJournal
	if err := c.ShouldBindJSON(&thesisJournal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	thesisJournal.StudentID = id
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
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
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
