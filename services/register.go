package services

import (
	"RU-Smart-Workspace/ru-smart-api/repositories"
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type (
	registerServices struct {
		registerRepo repositories.RegisterRepoInterface
		redis_cache  *redis.Client
	}

	RegisterRequest struct {
		STD_CODE string `json:"std_code" validate:"min=9,max=10,regexp=^[0-9]"`
		YEAR     string `json:"year" validate:"min=4,max=4,regexp=^[0-9]"`
	}

	RegisterScheduleRequest struct {
		YEAR     string `json:"year" validate:"min=4,max=4,regexp=^[0-9]"`
		SEMESTER string `json:"semester" validate:"min=1,max=1,regexp=^[1-5]"`
	}

	RegisterResponse struct {
		STD_CODE string `json:"std_code"`
		YEAR     string `json:"year"`
		RECORD   []RegisterRecord
	}

	RegisterScheduleResponse struct {
		YEAR     string `json:"course_year"`
		SEMESTER string `json:"course_semester"`
		RECORD   []RegisterMr30Record
	}

	RegisterYearResponse struct {
		STD_CODE string `json:"std_code"`
		RECORD   []registerYearRecord
	}

	registerYearRecord struct {
		YEAR string `json:"year"`
	}

	RegisterYearSemesterResponse struct {
		STD_CODE string `json:"std_code"`
		RECORD   []registerYearSemesterRecord
	}

	registerYearSemesterRecord struct {
		YEAR     string `json:"year"`
		SEMESTER string `json:"semester"`
	}

	RegisterMr30Record struct {
		ID                   string `json:"id"`
		COURSE_YEAR          string `json:"course_year"`
		COURSE_SEMESTER      string `json:"course_semester"`
		COURSE_NO            string `json:"course_no"`
		COURSE_METHOD        string `json:"course_method"`
		COURSE_METHOD_NUMBER string `json:"course_method_number"`
		DAY_CODE             string `json:"day_code"`
		TIME_CODE            string `json:"time_code"`
		ROOM_GROUP           string `json:"room_group"`
		INSTR_GROUP          string `json:"instr_group"`
		COURSE_METHOD_DETAIL string `json:"course_method_detail"`
		DAY_NAME_S           string `json:"day_name_s"`
		TIME_PERIOD          string `json:"time_period"`
		COURSE_ROOM          string `json:"course_room"`
		COURSE_INSTRUCTOR    string `json:"course_instructor"`
		SHOW_RU30            string `json:"show_ru30"`
		COURSE_CREDIT        string `json:"course_credit"`
		COURSE_PR            string `json:"course_pr"`
		COURSE_COMMENT       string `json:"course_comment"`
		COURSE_EXAMDATE      string `json:"course_examdate"`
	}

	RegisterRecord struct {
		YEAR      string `json:"year"`
		SEMESTER  string `json:"semester"`
		COURSE_NO string `json:"course_no"`
		CREDIT    string `json:"credit"`
	}

	RegisterServiceInterface interface {
		GetRegister(registerRequest RegisterRequest) (*RegisterResponse, error)
		GetListYear(std_code string) (*RegisterYearResponse, error)
		GetListYearSemester(std_code string) (*RegisterYearSemesterResponse, error)
		GetScheduleYearSemester(std_code string, registerScheduleRequest RegisterScheduleRequest) (*RegisterScheduleResponse, error)
		GetSchedule(std_code string) (*RegisterScheduleResponse, error)
	}
)

func NewRegisterServices(registerRepo repositories.RegisterRepoInterface, redis_cache *redis.Client) RegisterServiceInterface {
	return &registerServices{
		registerRepo: registerRepo,
		redis_cache:  redis_cache,
	}
}
