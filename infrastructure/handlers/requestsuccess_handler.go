package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *journalHandler) UpdateRequestSuccessStatus(c *gin.Context) {
	id := c.Param("id")

	journal, err := h.thesisJournalService.UpdateRequestSuccessStatus(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Journal request success updated status successfully",
		"data":    journal,
	})
}

func (h *journalHandler) GetRequestSuccessByID(c *gin.Context) {
	studentID := c.Param("id")

	request, err := h.thesisJournalService.GetRequestSuccessByID(c.Request.Context(), studentID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "check error" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Request success retrieved successfully",
		"data":    request,
	})
}

func (h *journalHandler) ListRequestSuccess(c *gin.Context) {
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

	requests, err := h.thesisJournalService.ListRequestSuccesss(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "check error" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "list request success retrieved successfully",
		"data":    requests,
	})
}



