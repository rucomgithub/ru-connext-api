package services

import (
	"RU-Smart-Workspace/ru-smart-api/repositories"

	"github.com/go-redis/redis/v8"
)

type (
	rotcsServices struct {
		rotcsRepo   repositories.RotcsRepoInterface
		redis_cache *redis.Client
	}

	RotcsRequest struct {
		StudentCode string `json:"StudentCode" validate:"min=10,max=10,regexp=^[0-9]{10}$"`
	}

	RotcsRegisterResponse struct {
		StudentCode string              `json:"studentCode"`
		Total       int                 `json:"total"`
		Detail      []RotcsDetailRecord `json:"detail"`
	}

	RotcsDetailRecord struct {
		YearReport   string `json:"yearReport"`
		LocationArmy string `json:"locationArmy"`
		LayerArmy    string `json:"layerArmy"`
		LayerReport  string `json:"layerReport"`
		TypeReport   string `json:"typeReport"`
		Status       string `json:"status"`
	}

	RotcsRegisterRecord struct {
		StudentCode string `json:"studentCode"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		FacultyId   string `json:"facultyId"`
		FacultyName string `json:"facultyName"`
	}

	RotcsExtendResponse struct {
		StudentCode string                      `json:"studentCode"`
		ExtendYear  string                      `json:"extendYear"`
		Code9       string                      `json:"code9"`
		Option1     string                      `json:"option1"`
		Option2     string                      `json:"option2"`
		Option3     string                      `json:"option3"`
		Option4     string                      `json:"option4"`
		Option5     string                      `json:"option5"`
		Option6     string                      `json:"option6"`
		Option7     string                      `json:"option7"`
		Option8     string                      `json:"option8"`
		Option9     string                      `json:"option9"`
		OptionOther string                      `json:"optionOther"`
		Total       int                         `json:"total"`
		Detail      []RotcsExtendDetailResponse `json:"detail"`
	}

	RotcsExtendDetailResponse struct {
		Id               string `json:"id"`
		RegisterYear     string `json:"registerYear"`
		RegisterSemester string `json:"registerSemester"`
		Credit           string `json:"credit"`
		Created          string `json:"created"`
		Modified         string `json:"modified"`
	}

	RotcsServiceInterface interface {
		GetRotcsRegister(rotcsRequest RotcsRequest) (*RotcsRegisterResponse, error)
		GetRotcsExtend(rotcsRequest RotcsRequest) (*RotcsExtendResponse, error)
	}
)

func NewRotcsServices(rotcsRepo repositories.RotcsRepoInterface, redis_cache *redis.Client) RotcsServiceInterface {
	return &rotcsServices{
		rotcsRepo:   rotcsRepo,
		redis_cache: redis_cache,
	}
}
