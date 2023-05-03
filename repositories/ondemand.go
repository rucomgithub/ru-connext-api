package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type (
	onDemandRepoDB struct {
		mysql_db *sql.DB
	}

	OnDemandRepo struct {
		STUDY_SEMESTER string `db:"STUDY_SEMESTER"`
		STUDY_YEAR     string `db:"STUDY_YEAR"`
		COURSE_NO      string `db:"COURSE_NO"`
		DAY_CODE       string `db:"DAY_CODE"`
		TIME_CODE      string `db:"TIME_CODE"`
		BUILDING_CODE  string `db:"BUILDING_CODE"`
		ROOM_CODE      string `db:"ROOM_CODE"`
	}

	OnDemandRepoInterface interface {
		GetOnDemand(std_code, year string) (*[]OnDemandRepo, error)
		GetOnDemandAll(course_no, year, semester string) (*[]OnDemandRepo, error)
	}
)

func NewOnDemandRepo(mysql_db *sql.DB) OnDemandRepoInterface {
	return &onDemandRepoDB{mysql_db: mysql_db}
}
