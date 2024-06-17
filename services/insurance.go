package services

import (
	"RU-Smart-Workspace/ru-smart-api/repositories"

	"github.com/go-redis/redis/v8"
)

type (
	InsuranceServices struct {
		insuranceRepo repositories.InsuranceRepoInterface
		redis_cache   *redis.Client
	}

	InsuranceRequest struct {
		StudentCode string `json:"StudentCode" validate:"min=10,max=10,regexp=^[0-9]{10}$"`
	}

	InsuranceResponse struct {
		StudentCode string            `json:"StudentCode"`
		Total       int               `json:"Total"`
		Detail      []InsuranceRecord `json:"Detail"`
	}

	InsuranceRecord struct {
		StudentCode     string `json:"StudentCode"`
		PersonCode      string `json:"PersonCode"`
		NameInsurance   string `json:"NameInsurance"`
		StartDate       string `json:"StartDate"`
		EndDate         string `json:"EndDate"`
		StatusInsurance string `json:"StatusInsurance"`
		TypeInsurance   string `json:"TypeInsurance"`
		YearMonth       string `json:"YearMonth"`
		Expire          string `json:"Expire"`
	}

	InsuranceServiceInterface interface {
		GetInsuranceListAll(insuranceRequest InsuranceRequest) (*InsuranceResponse, error)
	}
)

func NewInsuranceServices(insuranceRepo repositories.InsuranceRepoInterface, redis_cache *redis.Client) InsuranceServiceInterface {
	return &InsuranceServices{
		insuranceRepo: insuranceRepo,
		redis_cache:   redis_cache,
	}
}
