package services

import (
	"RU-Smart-Workspace/ru-smart-api/repositories"

	"github.com/go-redis/redis/v8"
)

type (
	EventServices struct {
		eventRepo   repositories.EventRepoInterface
		redis_cache *redis.Client
	}

	EventRequest struct {
		StdID string `json:"StdID" validate:"min=10,max=10,regexp=^[0-9]{10}$"`
	}

	EventResponse struct {
		StdID  string        `json:"StdID"`
		Total  int           `json:"Total"`
		Detail []EventRecord `json:"Detail"`
	}

	EventRecord struct {
		StdID    string `json:"StdID"`
		Title    string `json:"Title"`
		Time     string `json:"Time"`
		TypeName string `json:"TypeName"`
		Club     string `json:"Club"`
		Semester string `json:"Semester"`
		Year     string `json:"Year"`
	}

	EventServiceInterface interface {
		GetEventListAll(eventRequest EventRequest) (*EventResponse, error)
	}
)

func NewEventServices(eventRepo repositories.EventRepoInterface, redis_cache *redis.Client) EventServiceInterface {
	return &EventServices{
		eventRepo:   eventRepo,
		redis_cache: redis_cache,
	}
}
