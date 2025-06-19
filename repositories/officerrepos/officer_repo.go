package officerrepos

import (
	"github.com/jmoiron/sqlx"
)

type (
	officerRepoDB struct {
		oracle_db_dbg *sqlx.DB
	}

	UserLoginRepo struct {
		Username string `db:"USERNAME"`
		Role     string `db:"ROLE"`
		Key      string `db:"KEY"`
		Created  string `db:"CREATED"`
		Modified string `db:"MODIFIED"`
		Status   int    `db:"STATUS"`
	}

	Qualification struct {
		STD_CODE     string `db:"STD_CODE"`
		REQUEST_DATE string `db:"REQUEST_DATE"`
		OPERATE_DATE string `db:"OPERATE_DATE"`
		CONFIRM_DATE string `db:"CONFIRM_DATE"`
		CANCEL_DATE  string `db:"CANCEL_DATE"`
		STATUS       string `db:"STATUS"`
		CREATED      string `db:"CREATED"`
		MODIFIED     string `db:"MODIFIED"`
		DESCRIPTION  string `db:"DESCRIPTION"`
	}

	OfficerRepoInterface interface {
		GetUserLogin(username string) (*UserLoginRepo, error)

		GetQualificationAll() (*[]Qualification, error)
		GetQualification(std_code string) (*Qualification, error)
		AddQualification(std_code string) error
		UpdateQualification(std_code, status, description string) (int64, error)

		FindReport(startdate,enddate string) ([]map[string]interface{},error)
	}
)

func NewOfficerRepo(oracle_db_dbg *sqlx.DB) OfficerRepoInterface {
	return &officerRepoDB{oracle_db_dbg: oracle_db_dbg}
}
