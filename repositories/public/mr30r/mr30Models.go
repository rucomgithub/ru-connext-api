package mr30r

import "github.com/jmoiron/sqlx"

type (
	mr30RepoDB struct {
		oracle_db *sqlx.DB
	}

	Mr30Repo struct {
		ID                   string `db:"ID"`
		COURSE_YEAR          string `db:"COURSE_YEAR"`
		COURSE_SEMESTER      string `db:"COURSE_SEMESTER"`
		COURSE_NO            string `db:"COURSE_NO"`
		COURSE_METHOD        string `db:"COURSE_METHOD"`
		COURSE_METHOD_NUMBER string `db:"COURSE_METHOD_NUMBER"`
		DAY_CODE             string `db:"DAY_CODE"`
		TIME_CODE            string `db:"TIME_CODE"`
		ROOM_GROUP           string `db:"ROOM_GROUP"`
		INSTR_GROUP          string `db:"INSTR_GROUP"`
		COURSE_CREDIT        string `db:"COURSE_CREDIT"`
		COURSE_METHOD_DETAIL string `db:"COURSE_METHOD_DETAIL"`
		DAY_NAME_S           string `db:"DAY_NAME_S"`
		TIME_PERIOD          string `db:"TIME_PERIOD"`
		COURSE_ROOM          string `db:"COURSE_ROOM"`
		COURSE_INSTRUCTOR    string `db:"COURSE_INSTRUCTOR"`
		SHOW_RU30            string `db:"SHOW_RU30"`
		COURSE_PR            string `db:"COURSE_PR"`
		COURSE_COMMENT       string `db:"COURSE_COMMENT"`
		COURSE_EXAMDATE      string `db:"COURSE_EXAMDATE"`
	}

	Mr30YearRepo struct {
		COURSE_YEAR     string `db:"COURSE_YEAR"`
		COURSE_SEMESTER string `db:"COURSE_SEMESTER"`
	}

	Mr30RepoInterface interface {
		GetMr30(course_year, course_semester string) (*[]Mr30Repo, error)
		GetMr30Year() (*[]Mr30YearRepo, error)
	}
)

func NewMr30Repo(oracle_db *sqlx.DB) Mr30RepoInterface {
	return &mr30RepoDB{oracle_db: oracle_db}
}
