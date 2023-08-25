package repositories

import "github.com/jmoiron/sqlx"

type (
	gradeRepoDB struct {
		oracle_db *sqlx.DB
	}

	GradeRepo struct {
		REGIS_YEAR     string `db:"REGIS_YEAR"`
		REGIS_SEMESTER string `db:"REGIS_SEMESTER"`
		COURSE_NO      string `db:"COURSE_NO"`
		CREDIT         string `db:"CREDIT"`
		GRADE          string `db:"GRADE"`
	}

	GradeRepoInterface interface {
		GetGradeYear(std_code, year string) (*[]GradeRepo, error)
		GetGradeAll(std_code string) (*[]GradeRepo, error)
	}
)

func NewGradeRepo(oracle_db *sqlx.DB) GradeRepoInterface {
	return &gradeRepoDB{oracle_db: oracle_db}
}
