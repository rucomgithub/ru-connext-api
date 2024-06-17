package repositories

import (
	"github.com/jmoiron/sqlx"
)

type (
	insuranceRepoDB struct {
		mysql_db *sqlx.DB
	}

	InsuranceRepo struct {
		StudentCode     string `db:"studentcode"`
		PersonCode      string `db:"idcard"`
		NameInsurance   string `db:"nameinsurance"`
		StartDate       string `db:"startdate"`
		EndDate         string `db:"enddate"`
		StatusInsurance string `db:"statusinsurance"`
		TypeInsurance   string `db:"typeinsurance"`
		YearMonth       string `db:"yearmonth"`
	}

	InsuranceRepoInterface interface {
		GetInsuranceListAll(studentcode string) (*[]InsuranceRepo, error)
	}
)

func NewInsuranceRepo(mysql_db *sqlx.DB) InsuranceRepoInterface {
	return &insuranceRepoDB{mysql_db}
}
