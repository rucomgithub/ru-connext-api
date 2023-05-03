package services

import (
	"RU-Smart-Workspace/ru-smart-api/repositories"

	"github.com/go-redis/redis/v8"
)

type (
	onDemandServices struct {
		onDemandRepo repositories.OnDemandRepoInterface
		redis_cache  *redis.Client
	}

	OnDemandRequest struct {
		COURSE_NO string `json:"course_no" validate:"min=7,max=7,regexp=^[A-Z]{3}[0-9]{4}$"`
		YEAR      string `json:"year" validate:"min=4,max=4,regexp=^[0-9]"`
		SEMESTER  string `json:"semester" validate:"min=1,max=1,regexp=^[0-9]"`
	}

	OnDemandResponse struct {
		COURSE_NO string           `json:"course_no"`
		YEAR      string           `json:"year"`
		SEMESTER  string           `json:"semester"`
		COUNT     int              `json:"count"`
		RECORD    []onDemandRecord `json:"record"`
	}

	onDemandRecord struct {
		STUDY_SEMESTER string `db:"STUDY_SEMESTER"`
		STUDY_YEAR     string `db:"STUDY_YEAR"`
		COURSE_NO      string `db:"COURSE_NO"`
		DAY_CODE       string `db:"DAY_CODE"`
		TIME_CODE      string `db:"TIME_CODE"`
		BUILDING_CODE  string `db:"BUILDING_CODE"`
		ROOM_CODE      string `db:"ROOM_CODE"`
	}

	OnDemandServiceInterface interface {
		OnDemand(onDemandRequest OnDemandRequest) (*OnDemandResponse, error)
		OnDemandAll(onDemandRequest OnDemandRequest) (*OnDemandResponse, error)
	}
)

func NewOnDemandServices(onDemandRepo repositories.OnDemandRepoInterface, redis_cache *redis.Client) OnDemandServiceInterface {
	return &onDemandServices{
		onDemandRepo: onDemandRepo,
		redis_cache:  redis_cache,
	}
}
