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

	OfficerRepoInterface interface {
		GetUserLogin(username string) (*UserLoginRepo, error)
	}
)

func NewOfficerRepo(oracle_db_dbg *sqlx.DB) OfficerRepoInterface {
	return &officerRepoDB{oracle_db_dbg: oracle_db_dbg}
}
