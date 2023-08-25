package mr30s

import (
	mr30r "RU-Smart-Workspace/ru-smart-api/repositories/public/mr30r"
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type (
	mr30Services struct {
		mr30Repo    mr30r.Mr30RepoInterface
		redis_cache *redis.Client
	}

	Mr30Request struct {
		Course_year     string `json:"course_year"`
		Course_semester string `json:"course_semester"`
	}

	Mr30Response struct {
		COURSE_YEAR     string `json:"course_year"`
		COURSE_SEMESTER string `json:"course_semester"`
		RECORD          []mr30Record
	}

	Mr30YearResponse struct {
		RECORDYEAR []mr30YearRecord
	}

	mr30YearRecord struct {
		COURSE_YEAR     string `json:"course_year"`
		COURSE_SEMESTER string `json:"course_semester"`
	}

	mr30Record struct {
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
		COURSE_CREDIT        string `json:"course_credit"`
		COURSE_METHOD_DETAIL string `json:"course_method_detail"`
		DAY_NAME_S           string `json:"day_name_s"`
		TIME_PERIOD          string `json:"time_period"`
		COURSE_ROOM          string `json:"course_room"`
		COURSE_INSTRUCTOR    string `json:"course_instructor"`
		SHOW_RU30            string `json:"show_ru30"`
		COURSE_PR            string `json:"course_pr"`
		COURSE_COMMENT       string `json:"course_comment"`
		COURSE_EXAMDATE      string `json:"course_examdate"`
	}

	Mr30ServiceInterface interface {
		GetMr30(course_year, course_semester string) (*Mr30Response, error)
		GetMr30Year() (*Mr30YearResponse, error)
		GetMr30Searching(course_year, course_semester, course_no string) (*Mr30Response, error)
		GetMr30Pagination(course_year, course_semester, limit, offset string) (*Mr30Response, error)
	}
)

func NewMr30Services(mr30Repo mr30r.Mr30RepoInterface, redis_cache *redis.Client) Mr30ServiceInterface {
	return &mr30Services{
		mr30Repo:    mr30Repo,
		redis_cache: redis_cache,
	}
}
