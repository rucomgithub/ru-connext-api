package repositories

import (
	"github.com/jmoiron/sqlx"
)

type (
	eventRepoDB struct {
		mysql_db *sqlx.DB
	}

	EventRepo struct {
		StdID    string `db:"std_id"`
		Title    string `db:"event_title"`
		Time     string `db:"event_time"`
		TypeName string `db:"type_name"`
		Club     string `db:"event_club"`
		Semester string `db:"event_semester"`
		Year     string `db:"evnet_year"`
	}

	EventRepoInterface interface {
		GetEventListAll(studentcode string) (*[]EventRepo, error)
	}
)

func NewEventRepo(mysql_db *sqlx.DB) EventRepoInterface {
	return &eventRepoDB{mysql_db}
}
