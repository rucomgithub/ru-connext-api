package services

import (
	"RU-Smart-Workspace/ru-smart-api/repositories"

	"github.com/go-redis/redis/v8"
)

type (
	scholarShipServices struct {
		scholarShipRepo repositories.ScholarshipRepoInterface
		redis_cache     *redis.Client
	}

	ScholarShipRequest struct {
		STD_CODE string `json:"std_code" validate:"min=9,max=10,regexp=^[0-9]"`
	}

	ScholarShipResponse struct {
		STD_CODE string `json:"std_code"`
		Total    int    `json:"total"`
		RECORD   []scholarShipRecord
	}

	scholarShipRecord struct {
		ID                  int     `json:"id"`
		StudentCode         string  `json:"std_code"`
		ScholarshipYear     int     `json:"scholarship_year"`
		ScholarshipSemester string  `json:"scholarship_semester"`
		ScholarshipType     string  `json:"scholarship_type"`
		Amount              float64 `json:"amount"`
		Created             string  `json:"created"`
		Modified            string  `json:"modified"`
		Username            string  `json:"username"`
		SubsidyNumber       string  `json:"subsidy_no"`
		SubsidyNameThai     string  `json:"subsidy_name_thai"`
	}

	ScholarShipServiceInterface interface {
		GetScholarshipAll(scholarShipRequest ScholarShipRequest) (*ScholarShipResponse, error)
	}
)

func NewScholarShipServices(scholarShipRepo repositories.ScholarshipRepoInterface, redis_cache *redis.Client) ScholarShipServiceInterface {
	return &scholarShipServices{
		scholarShipRepo: scholarShipRepo,
		redis_cache:     redis_cache,
	}
}
