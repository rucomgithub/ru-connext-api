package studentr

import (
	"github.com/jmoiron/sqlx"
)

type (
	studentRepoDB struct {
		oracle_db *sqlx.DB
	}

	StudentProfileRepo struct {
		STD_CODE             string `db:"STD_CODE"`
		NAME_THAI            string `db:"NAME_THAI"`
		NAME_ENG             string `db:"NAME_ENG"`
		BIRTH_DATE           string `db:"BIRTH_DATE"`
		STD_STATUS_DESC_THAI string `db:"STD_STATUS_DESC_THAI"`
		CITIZEN_ID           string `db:"CITIZEN_ID"`
		REGIONAL_NAME_THAI   string `db:"REGIONAL_NAME_THAI"`
		STD_TYPE_DESC_THAI   string `db:"STD_TYPE_DESC_THAI"`
		FACULTY_NAME_THAI    string `db:"FACULTY_NAME_THAI"`
		MAJOR_NAME_THAI      string `db:"MAJOR_NAME_THAI"`
		WAIVED_NO            string `db:"WAIVED_NO"`
		WAIVED_PAID          string `db:"WAIVED_PAID"`
		WAIVED_TOTAL_CREDIT  int    `db:"WAIVED_TOTAL_CREDIT"`
		CHK_CERT_NAME_THAI   string `db:"CHK_CERT_NAME_THAI"`
		PENAL_NAME_THAI      string `db:"PENAL_NAME_THAI"`
		MOBILE_TELEPHONE     string `db:"MOBILE_TELEPHONE"`
		EMAIL_ADDRESS        string `db:"EMAIL_ADDRESS"`
	}

	RegisterRepo struct {
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

	RegisterAllRepo struct {
		YEAR      string `db:"YEAR"`
		SEMESTER  string `db:"SEMESTER"`
		COURSE_NO string `db:"COURSE_NO"`
		STD_CODE  string `db:"STD_CODE"`
		CREDIT    string `db:"CREDIT"`
	}

	PrepareTokenRepo struct {
		STD_CODE string `db:"STD_CODE"`
		STATUS   int    `db:"STATUS"`
	}

	StudentRepoInterface interface {
		GetStudentProfile(studentCode string) (*StudentProfileRepo, error)
		GetRegisterAll(studentCode, courseYear string) (*[]RegisterAllRepo, error)
		GetRegister(studentCode, courseYear, courseSemester string) (*[]RegisterRepo, error)
		Authentication(studentCode string) (*PrepareTokenRepo, error)
	}
)

func NewStudentRepo(oracle_db *sqlx.DB) StudentRepoInterface {
	return &studentRepoDB{oracle_db: oracle_db}
}
