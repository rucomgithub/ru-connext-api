package masterrepo

import (
	"github.com/jmoiron/sqlx"
)

type (
	studentRepoDB struct {
		oracle_db     *sqlx.DB
		oracle_db_dbg *sqlx.DB
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
		THAI_NAME            string  `db:"THAI_NAME"`
		ENG_NAME             string  `db:"ENG_NAME"`
		THAI_DEGREE          string  `db:"THAI_DEGREE"`
		ENG_DEGREE           string  `db:"ENG_DEGREE"`
		THAI_MAJOR           string  `db:"THAI_MAJOR"`
	}

	StudentSuccessRepo struct {
		STD_CODE        string  `db:"STD_CODE"`
		NAME_THAI       string  `db:"NAME_THAI"`
		NAME_ENG        string  `db:"NAME_ENG"`
		YEAR            string  `db:"YEAR"`
		SEMESTER        string  `db:"SEMESTER"`
		CURR_NAME       string  `db:"CURR_NAME"`
		CURR_ENG        string  `db:"CURR_ENG"`
		THAI_NAME       string  `db:"THAI_NAME"`
		ENG_NAME        string  `db:"ENG_NAME"`
		MAJOR_NAME      string  `db:"MAJOR_NAME"`
		MAJOR_ENG       string  `db:"MAJOR_ENG"`
		MAIN_MAJOR_THAI string  `db:"MAIN_MAJOR_THAI"`
		MAIN_MAJOR_ENG  string  `db:"MAIN_MAJOR_ENG"`
		PLAN            string  `db:"PLAN"`
		GPA             float32 `db:"GPA"`
		CONFERENCE_NO   string  `db:"CONFERENCE_NO"`
		SERIAL_NO       string  `db:"SERIAL_NO"`
		CONFERENCE_DATE string  `db:"CONFERENCE_DATE"`
		ADMIT_DATE      string  `db:"ADMIT_DATE"`
		ADMIT_DATE_EN      string  `db:"ADMIT_DATE_EN"`
		GRADUATED_DATE  string  `db:"GRADUATED_DATE"`
		GRADUATED_DATE_EN  string  `db:"GRADUATED_DATE_EN"`
		CONFIRM_DATE    string  `db:"CONFIRM_DATE"`
		MOBILE     		string `db:"MOBILE"`
		EMAIL  			string `db:"EMAIL"`
	}

	RegisterRepo struct {
		YEAR      string `db:"YEAR"`
		SEMESTER  string `db:"SEMESTER"`
		STD_CODE  string `db:"STD_CODE"`
		COURSE_NO string `db:"COURSE_NO"`
		CREDIT    int    `db:"CREDIT"`
	}

	GradeRepo struct {
		YEAR      string `db:"YEAR"`
		SEMESTER  string `db:"SEMESTER"`
		STD_CODE  string `db:"STD_CODE"`
		COURSE_NO string `db:"COURSE_NO"`
		CREDIT    int    `db:"CREDIT"`
		GRADE     string `db:"GRADE"`
	}

	StudentRepo struct {
		STD_CODE string `db:"STD_CODE"`
	}

	GPARepo struct {
		SUMMARY_CREDIT int     `db:"CREDIT"`
		GPA            float32 `db:"GPA"`
	}

	PrivacyPolicy struct {
		STD_CODE string `db:"STD_CODE"`
		VERSION  string `db:"VERSION"`
		STATUS   string `db:"STATUS"`
		CREATED  string `db:"CREATED"`
		MODIFIED string `db:"MODIFIED"`
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

	Company struct {
		STD_CODE string `db:"STD_CODE"`
		EMAIL    string `db:"EMAIL"`
		FULLNAME string `db:"FULLNAME"`
		COMPANY  string `db:"COMPANY"`
		CREATED  string `db:"CREATED"`
		MODIFIED string `db:"MODIFIED"`
	}

	StudentRepoInterface interface {
		GetStudentProfile(studentCode string) (*StudentProfileRepo, error)

		GetStudentSuccess(studentCode string) (*StudentSuccessRepo, error)

		GetRegisterByYear(std_code, year string) (*[]RegisterRepo, error)
		GetRegisterAll(std_code string) (*[]RegisterRepo, error)

		GetGradeByYear(std_code, year string) (*[]GradeRepo, error)
		GetGradeAll(std_code string) (*[]GradeRepo, error)

		GetGpaAll(std_code string) (*GPARepo, error)
		GetGpaByYear(std_code, year string) (*GPARepo, error)

		AddPrivacyPolicy(std_code, version string) error
		UpdatePrivacyPolicy(std_code, version, status string) error
		GetPrivacyPolicy(std_code, version string) (*PrivacyPolicy, error)

		GetQualification(std_code string) (*Qualification, error)
		AddQualification(std_code string) error

		GetCommpanyByEmail(email string) (*Company, error)
		GetCommpany(std_code, email string) (*Company, error)
		AddCommpany(std_code, email, fullname, company string) error
	}
)

func NewStudentRepo(oracle_db_dbg *sqlx.DB) StudentRepoInterface {
	return &studentRepoDB{oracle_db_dbg: oracle_db_dbg}
}
