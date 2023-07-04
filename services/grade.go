package services

import (
	"RU-Smart-Workspace/ru-smart-api/repositories"

	"github.com/go-redis/redis/v8"
)

type (
	gradeServices struct {
		gradeRepo   repositories.GradeRepoInterface
		redis_cache *redis.Client
	}

	GradeRequest struct {
		STD_CODE string `json:"std_code" validate:"min=9,max=10,regexp=^[0-9]"`
		YEAR     string `json:"year" validate:"min=4,max=4,regexp=^[0-9]"`
	}

	GradeResponse struct {
		STD_CODE string `json:"std_code"`
		YEAR     string `json:"year"`
		RECORD   []gradeRecord
	} 

	gradeRecord struct {
		REGIS_YEAR     string `json:"regis_year"`
		REGIS_SEMESTER string `json:"regis_semester"`
		COURSE_NO      string `json:"course_no"`
		CREDIT         string `json:"credit"`
		GRADE          string `json:"grade"`
	}

	GradeServiceInterface interface {
		GradeYear(gradeRequest GradeRequest) (*GradeResponse, error)
		GradeAll(std_code string) (*GradeResponse, error)
	}
)

func NewGradeServices(gradeRepo repositories.GradeRepoInterface, redis_cache *redis.Client) GradeServiceInterface {
	return &gradeServices{
		gradeRepo:   gradeRepo,
		redis_cache: redis_cache,
	}
}
