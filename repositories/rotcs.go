package repositories

import "github.com/jmoiron/sqlx"

type (
	rotcsRepoDB struct {
		mysql_db *sqlx.DB
	}

	RotcsRegisterRepo struct {
		StudentCode  string `db:"studentCode"`
		LocationArmy string `db:"locationArmy"`
		LayerArmy    string `db:"layerArmy"`
		YearReport   string `db:"yearReport"`
		LayerReport  string `db:"layerReport"`
		TypeReport   string `db:"typeReport"`
		Status       string `db:"status"`
	}

	RotcsExtendRepo struct {
		StudentCode string                  `db:"studentCode"`
		ExtendYear  string                  `db:"extendYear"`
		Code9       string                  `db:"code9"`
		Option1     string                  `db:"option1"`
		Option2     string                  `db:"option2"`
		Option3     string                  `db:"option3"`
		Option4     string                  `db:"option4"`
		Option5     string                  `db:"option5"`
		Option6     string                  `db:"option6"`
		Option7     string                  `db:"option7"`
		Option8     string                  `db:"option8"`
		Option9     string                  `db:"option9"`
		OptionOther string                  `db:"optionOther"`
		Detail      []RotcsExtendDetailRepo `db:"detail"`
	}

	RotcsExtendDetailRepo struct {
		Id               string `db:"id"`
		RegisterYear     string `db:"registerYear"`
		RegisterSemester string `db:"registerSemester"`
		Credit           string `db:"credit"`
		Created          string `db:"created"`
		Modified         string `db:"modified"`
	}

	RotcsRepoInterface interface {
		GetRotcsRegister(std_code string) (*[]RotcsRegisterRepo, error)
		GetRotcsExtend(std_code string) (*RotcsExtendRepo, error)
	}
)

func NewRotcsRepo(mysql_db *sqlx.DB) RotcsRepoInterface {
	return &rotcsRepoDB{mysql_db}
}
