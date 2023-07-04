package repositories

import "github.com/jmoiron/sqlx"

type (
	ondemandRepoDB struct {
		mysql_db *sqlx.DB
	}

	OndemandRepo struct {
		SUBJECT_CODE     string `db:"subject_code"`
		SUBJECT_ID     string `db:"subject_id"`
		SUBJECT_NAME_ENG string `db:"subject_NameEng"`
		SEMESTER      string `db:"semester"`
		YEAR         string `db:"year"`
		//TOTAL	int `db:"total"`
		DETAIL []OndemandSubjectCodeRepo  `db:"detail"`
	}

	OndemandSubjectCodeRepo struct {
		AUDIO_ID     string `db:"audio_id"`
		SUBJECT_CODE     string `db:"subject_code"`
		SUBJECT_ID	string `db:"subject_id"`
		AUDIO_SEC	string `db:"audio_sec"`
		SEM      string `db:"sem"`
		YEAR         string `db:"year"`
		AUDIO_CREATE         string `db:"audio_create"`
		AUDIO_STATUS         string `db:"audio_status"`
		AUDIO_TEACH         string `db:"audio_teach"`
		AUDIO_COMMENT         string `db:"audio_comment"`

	}

	OndemandRepoInterface interface {
		//GetOndemandAll() (*[]OndemandRepo, error)
		GetOndemandAll(course_no, semester,year string) (*OndemandRepo, error)
		GetOndemandSubjectCode(subject_code string) (*[]OndemandSubjectCodeRepo, error)
		// GetGradeAll(std_code string) (*[]GradeRepo, error)
	}
)

func NewOndemandRepo(mysql_db *sqlx.DB) OndemandRepoInterface {
	return &ondemandRepoDB{mysql_db: mysql_db}
}