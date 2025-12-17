package handlers

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *journalHandler) CreateSimilarity(c *gin.Context) {
	var thesisJournal entities.ThesisSimilarity
	if err := c.ShouldBindJSON(&thesisJournal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "check error" + err.Error(),
		})
		return
	}

	if err := h.thesisJournalService.CreateThesisSimilarity(c.Request.Context(), &thesisJournal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "check error" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Similarity created successfully",
		"data":    thesisJournal,
	})
}

func (h *journalHandler) UpdateSimilarity(c *gin.Context) {
	id := c.Param("studentId")

	var thesisJournal entities.ThesisSimilarity
	if err := c.ShouldBindJSON(&thesisJournal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
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
		"success": true,
		"message": "Similarity updated successfully.",
		"data":    thesisJournal,
	})
}

func (h *journalHandler) UpdateSimilarityStatus(c *gin.Context) {
	id := c.Param("id")

	journal, err := h.thesisJournalService.UpdateThesisSimilarityStatus(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Journal similarity updated status successfully",
		"data":    journal,
	})
}

func (h *journalHandler) GetSimilarityByID(c *gin.Context) {
	studentID := c.Param("studentId")

	thesisJournal, err := h.thesisJournalService.GetSimilarityByStudentID(c.Request.Context(), studentID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "check error" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
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
			"success": false,
			"message": "check error" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "list similarity retrieved successfully",
		"data":    students,
	})
}

func (h *journalHandler) DeleteThesisSimilarity(c *gin.Context) {
	id := c.Param("id")

	if err := h.thesisJournalService.DeleteThesisSimilarity(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "check error" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Similarity deleted successfully",
	})
}

func (h *journalHandler) CreateSimilarityMaster(c *gin.Context) {

	token, err := middlewares.GetHeaderAuthorization(c)

	fmt.Println(token)

	if err != nil {
		err = errors.New("ไม่พบ token login.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "ไม่พบ token login."})
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
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "ไม่พบ claims user."})
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
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้."})
		c.Abort()
		return
	}

	var thesisSimilarity entities.ThesisSimilarity
	if err := c.ShouldBindJSON(&thesisSimilarity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "โปรดระบุแบบฟอร์มให้ครบ " + err.Error(),
		})
		return
	}

	thesisSimilarity.StudentID = claim.StudentCode

	if err := h.thesisJournalService.CreateThesisSimilarity(c.Request.Context(), &thesisSimilarity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "ไม่สามารถบันทึกข้อมูลได้ " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Similarity created successfully",
		"data":    thesisSimilarity,
	})
}

func (h *journalHandler) UpdateSimilarityMaster(c *gin.Context) {
	token, err := middlewares.GetHeaderAuthorization(c)

	fmt.Println(token)

	if err != nil {
		err = errors.New("ไม่พบ token login.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "ไม่พบ token login."})
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
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "ไม่พบ claims user."})
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
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้."})
		c.Abort()
		return
	}

	var thesisSimilarity entities.ThesisSimilarity
	if err := c.ShouldBindJSON(&thesisSimilarity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "โปรดระบุแบบฟอร์มให้ครบ " + err.Error(),
		})
		return
	}

	thesisSimilarity.StudentID = claim.StudentCode
	log.Print(thesisSimilarity.StudentID)
	if err := h.thesisJournalService.UpdateThesisSimilarity(c.Request.Context(), &thesisSimilarity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Similarity updated successfully.",
		"data":    thesisSimilarity,
	})
}

func (h *journalHandler) GetSimilarityMaster(c *gin.Context) {

	token, err := middlewares.GetHeaderAuthorization(c)

	fmt.Println(token)

	if err != nil {
		err = errors.New("ไม่พบ token login.")
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "ไม่พบ token login."})
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
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "ไม่พบ claims user."})
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
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "สิทธิ์ไม่สามารถเข้าถึงข้อมูลส่วนนี้ได้."})
		c.Abort()
		return
	}

	studentID := claim.StudentCode

	log.Print(studentID)

	thesisJournal, err := h.thesisJournalService.GetSimilarityByStudentID(c.Request.Context(), studentID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "check error" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Journal retrieved successfully",
		"data":    thesisJournal,
	})
}
