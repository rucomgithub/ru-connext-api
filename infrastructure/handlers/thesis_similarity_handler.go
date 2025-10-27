package handlers

import (
	"net/http"
	"strconv"
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"log"
	"github.com/gin-gonic/gin"
)

func (h *journalHandler) CreateSimilarity(c *gin.Context) {
	var thesisJournal entities.ThesisSimilarity
	if err := c.ShouldBindJSON(&thesisJournal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success" : false,
			"message": "check error" + err.Error(),
		})
		return
	}

	if err := h.thesisJournalService.CreateThesisSimilarity(c.Request.Context(), &thesisJournal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success" : false,
			"message": "check error" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success" : true,
		"message": "Similarity created successfully",
		"data":    thesisJournal,
	})
}

func (h *journalHandler) UpdateSimilarity(c *gin.Context) {
	id := c.Param("studentId")

	var thesisJournal entities.ThesisSimilarity
	if err := c.ShouldBindJSON(&thesisJournal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success" : false,
			"message": "check error" + err.Error(),
		})
		return
	}

	thesisJournal.StudentID = id
	log.Print(thesisJournal.StudentID)
	if err := h.thesisJournalService.UpdateThesisSimilarity(c.Request.Context(), &thesisJournal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success" : true,
		"message": "Similarity updated successfully.",
		"data":    thesisJournal,
	})
}

func (h *journalHandler) GetSimilarityByID(c *gin.Context) {
	studentID := c.Param("studentId")

	thesisJournal, err := h.thesisJournalService.GetSimilarityByStudentID(c.Request.Context(), studentID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success" : false,
			"message": "check error" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success" : true,
		"message": "Journal retrieved successfully",
		"data":    thesisJournal,
	})
}

func (h *journalHandler) ListSimilaritys(c *gin.Context) {
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

	students, err := h.thesisJournalService.ListThesisSimilaritys(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success" : false,
			"message": "check error" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success" : true,
		"message": "list similarity retrieved successfully",
		"data":    students,
	})
}

func (h *journalHandler) DeleteThesisSimilarity(c *gin.Context) {
	id := c.Param("id")

	if err := h.thesisJournalService.DeleteThesisSimilarity(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success" : false,
			"message": "check error" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success" : true,
		"message": "Similarity deleted successfully",
	})
}