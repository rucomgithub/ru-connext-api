package repositories

import "github.com/jmoiron/sqlx"

type (
	registerRepoDB struct {
		oracle_db *sqlx.DB
	}

	RegisterRepo struct {
		YEAR      string `db:"YEAR"`
		SEMESTER  string `db:"SEMESTER"`
		COURSE_NO string `db:"COURSE_NO"`
		STD_CODE  string `db:"STD_CODE"`
		CREDIT    string `db:"CREDIT"`
	}

	YearRepo struct {
		YEAR string `db:"YEAR"`
	}

	YearSemesterRepo struct {
		YEAR     string `db:"YEAR"`
		SEMESTER string `db:"SEMESTER"`
	}

	ScheduleRepo struct {
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
		COURSE_METHOD_DETAIL string `db:"COURSE_METHOD_DETAIL"`
		DAY_NAME_S           string `db:"DAY_NAME_S"`
		TIME_PERIOD          string `db:"TIME_PERIOD"`
		COURSE_ROOM          string `db:"COURSE_ROOM"`
		COURSE_INSTRUCTOR    string `db:"COURSE_INSTRUCTOR"`
		SHOW_RU30            string `db:"SHOW_RU30"`
		COURSE_CREDIT        string `db:"COURSE_CREDIT"`
		COURSE_PR            string `db:"COURSE_PR"`
		COURSE_COMMENT       string `db:"COURSE_COMMENT"`
		COURSE_EXAMDATE      string `db:"COURSE_EXAMDATE"`
	}

	RegisterRepoInterface interface {
		GetRegisterAll(std_code, year string) (*[]RegisterRepo, error)
		GetListYearAll(std_code string) (*[]YearRepo, error)
		GetListYearSemesterAll(std_code string) (*[]YearSemesterRepo, error)
		GetScheduleAll(year, semester, studentCode string) (*[]ScheduleRepo, error)
		GetYearSemesterLatest() (*YearSemesterRepo, error)
	}
)

func NewRegisterRepo(oracle_db *sqlx.DB) RegisterRepoInterface {
	return &registerRepoDB{oracle_db: oracle_db}
}
