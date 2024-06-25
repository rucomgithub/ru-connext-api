package handlers

import (
	"RU-Smart-Workspace/ru-smart-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandlers struct {
	eventServices services.EventServiceInterface
}

func NewEventHandlers(eventServices services.EventServiceInterface) EventHandlers {
	return EventHandlers{eventServices: eventServices}
}

func (h *EventHandlers) GetEventListAll(c *gin.Context) {
	var requestBody services.EventRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	eventResponse, err := h.eventServices.GetEventListAll(requestBody)

	if err != nil {
		c.Error(err)
		c.Set("line", GetLineNumber())
		c.Set("file", GetFileName())
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, eventResponse)

}
