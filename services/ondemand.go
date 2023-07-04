package services

import (
	"RU-Smart-Workspace/ru-smart-api/repositories"

	"github.com/go-redis/redis/v8"
)

type (
	ondemandServices struct {
		ondemandRepo   repositories.OndemandRepoInterface
		redis_cache *redis.Client
	}

	OndemandRequest struct {
		SUBJECT_ID 	string `json:"subject_id" validate:"min=7,max=7,regexp=^[A-Z]{3}[0-9]{4}$"`
		SEMESTER   	string `json:"semester" validate:"min=1,max=1"`
		YEAR		string `json:"year" validate:"min=2,max=2,regexp=^[0-9]"`
	}

	OndemandSubjectCodeRequest struct {
		SUBJECT_CODE 	string `json:"subject_code" validate:"min=2,max=10,regexp=^[0-9]"`
	}

	OndemandResponse struct {
		SUBJECT_ID  string `json:"subject_id"`
		SEMESTER    string `json:"semester"`
		YEAR        string `json:"year"`
		RECORD   	ondemandRecord
	} 

	ondemandRecord struct {
		SUBJECT_CODE	string `json:"subject_code"`
		SUBJECT_ID     	string `json:"subject_id"`
		SUBJECT_NAME_ENG string `json:"subject_name_eng"`
		SEMESTER        string `json:"semester"`
		YEAR          	string `json:"year"`
		TOTAL			int `json:"total"`
		DETAIL   	[]ondemandSubjectCodeRecord `json:"detail"`
	}

	OndemandSubjectCodeResponse struct {
		SUBJECT_CODE  string `json:"subject_code"`
		RECORD   	[]ondemandSubjectCodeRecord
	}

	ondemandSubjectCodeRecord struct {
		AUDIO_ID     string `json:"audio_id "`
		SUBJECT_CODE     string `json:"subject_code"`
		SUBJECT_ID	string `json:"subject_id"`
		AUDIO_SEC	string `json:"audio_sec"`
		SEM      string `json:"sem"`
		YEAR         string `json:"year"`
		AUDIO_CREATE         string `json:"audio_create"`
		AUDIO_STATUS         string `json:"audio_status"`
		AUDIO_TEACH         string `json:"audio_teach"`
		AUDIO_COMMENT         string `json:"audio_comment"`

	}

	OndemandServiceInterface interface {
		GetOndemandAll(ondemandRequest OndemandRequest) (*OndemandResponse,error)
		GetOndemandSubjectCode(ondemandSubjectCodeRequest OndemandSubjectCodeRequest) (*OndemandSubjectCodeResponse,error)
		// GradeYear(gradeRequest GradeRequest) (*GradeResponse, error)
		// GradeAll(std_code string) (*GradeResponse, error)
	}
)

func NewOndemandServices(ondemandRepo repositories.OndemandRepoInterface, redis_cache *redis.Client) OndemandServiceInterface {
	return &ondemandServices{
		ondemandRepo:   ondemandRepo,
		redis_cache: redis_cache,
	}
}
