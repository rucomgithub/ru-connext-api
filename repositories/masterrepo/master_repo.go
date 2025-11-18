package masterrepo

import (
	"github.com/jmoiron/sqlx"
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
)

type (
	studentRepoDB struct {
		oracle_db     *sqlx.DB
		oracle_db_dbg *sqlx.DB
	}

	StudentRequestSuccessRepo struct {
		ENROLL_YEAR             string `db:"ENROLL_YEAR"`
		ENROLL_SEMESTER         string `db:"ENROLL_SEMESTER"`
		STD_CODE                string `db:"STD_CODE"`
		PRENAME_THAI_S          string `db:"PRENAME_THAI_S"`
		PRENAME_ENG_S           string `db:"PRENAME_ENG_S"`
		FIRST_NAME              string `db:"FIRST_NAME"`
		LAST_NAME               string `db:"LAST_NAME"`
		FIRST_NAME_ENG          string `db:"FIRST_NAME_ENG"`
		LAST_NAME_ENG           string `db:"LAST_NAME_ENG"`
		THAI_NAME               string `db:"THAI_NAME"`
		PLAN_NO                 string `db:"PLAN_NO"`
		SEX                     string `db:"SEX"`
		REGINAL_NAME            string `db:"REGINAL_NAME"`
		SUBSIDY_NAME            string `db:"SUBSIDY_NAME"`
		STATUS_NAME_THAI        string `db:"STATUS_NAME_THAI"`
		BIRTH_DATE              string `db:"BIRTH_DATE"`
		STD_ADDR                string `db:"STD_ADDR"`
		ADDR_TEL                string `db:"ADDR_TEL"`
		JOB_POSITION            string `db:"JOB_POSITION"`
		STD_OFFICE              string `db:"STD_OFFICE"`
		OFFICE_TEL              string `db:"OFFICE_TEL"`
		DEGREE_NAME             string `db:"DEGREE_NAME"`
		BSC_DEGREE_NO           string `db:"BSC_DEGREE_NO"`
		BSC_DEGREE_THAI_NAME    string `db:"BSC_DEGREE_THAI_NAME"`
		BSC_INSTITUTE_NO        string `db:"BSC_INSTITUTE_NO"`
		INSTITUTE_THAI_NAME     string `db:"INSTITUTE_THAI_NAME"`
		CK_CERT_NO              string `db:"CK_CERT_NO"`
		CHK_CERT_NAME_THAI      string `db:"CHK_CERT_NAME_THAI"`
		ID                      string `db:"ID"`
		SUCCESS_YEAR            string `db:"SUCCESS_YEAR"`
		SUCCESS_SEMESTER        string `db:"SUCCESS_SEMESTER"`
		NAME_THAI       		string `db:"NAME_THAI"`
		NAME_ENG        		string `db:"NAME_ENG"`
		THESIS_THAI     		string `db:"THESIS_THAI"`
		THESIS_ENG      		string `db:"THESIS_ENG"`
		DEGREE          		string `db:"DEGREE"`
		REGISTRATION           	string `db:"REGISTRATION"`
		GRADES                	string `db:"GRADES"`
		ADDRESS               	string `db:"ADDRESS"`
		CREATED                 string `db:"CREATED"`
		MODIFIED                string `db:"MODIFIED"`
		THESIS_THAI_TITLE       string `db:"THESIS_THAI_TITLE"`
		THESIS_ENG_TITLE        string `db:"THESIS_ENG_TITLE"`
		THESIS_THESIS_NAME      string `db:"THESIS_THESIS_NAME"`
		THESIS_THESIS_TYPE      string `db:"THESIS_THESIS_TYPE"`
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

	RegisterFeeRepo struct {
		STD_CODE      string `db:"STD_CODE"`
		YEAR  		  string `db:"YEAR"`
		SEMESTER  	  string `db:"SEMESTER"`
		TOTAL_AMOUNT  int    `db:"TOTAL_AMOUNT"`
		REGIS_TYPE    string    `db:"REGIS_TYPE"`
		REGIS_NAME    string    `db:"REGIS_NAME"`
	}

	GradeRepo struct {
		YEAR      string `db:"YEAR"`
		SEMESTER  string `db:"SEMESTER"`
		STD_CODE  string `db:"STD_CODE"`
		COURSE_NO string `db:"COURSE_NO"`
		CREDIT    int    `db:"CREDIT"`
		GRADE     string `db:"GRADE"`
		COURSE_TYPE_NO string `db:"COURSE_TYPE_NO"`
		THAI_DESCRIPTION string `db:"THAI_DESCRIPTION"`
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
		GetStudentRequestSuccess(studentCode string) (*StudentRequestSuccessRepo, error)
		AddRequestSuccess(row *entities.RequestSuccess) error

		GetRegisterByYear(std_code, year string) (*[]RegisterRepo, error)
		GetRegisterAll(std_code string) (*[]RegisterRepo, error)
		GetRegisterFeeAll(std_code,role string) (*[]RegisterFeeRepo, error)

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
