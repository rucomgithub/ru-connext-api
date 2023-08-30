package repositories

import "github.com/jmoiron/sqlx"

type (
	scholarshipRepoDB struct {
		oracle_db *sqlx.DB
	}
	ScholarshipRepo struct {
		ID                  int     `db:"ID"`
		StudentCode         string  `db:"STD_CODE"`
		ScholarshipYear     int     `db:"SCHOLARSHIP_YEAR"`
		ScholarshipSemester string  `db:"SCHOLARSHIP_SEMESTER"`
		ScholarshipType     string  `db:"SCHOLARSHIP_TYPE"`
		Amount              float64 `db:"AMOUNT"`
		Created             string  `db:"CREATED"`
		Modified            string  `db:"MODIFIED"`
		Username            string  `db:"USERNAME"`
		SubsidyNumber       string  `db:"SUBSIDY_NO"`
		SubsidyNameThai     string  `db:"SUBSIDY_NAME_THAI"`
	}

	ScholarshipRepoInterface interface {
		GetScholarshipAll(std_code string) (*[]ScholarshipRepo, error)
	}
)

func NewScholarshipRepo(oracle_db *sqlx.DB) ScholarshipRepoInterface {
	return &scholarshipRepoDB{oracle_db: oracle_db}
}
