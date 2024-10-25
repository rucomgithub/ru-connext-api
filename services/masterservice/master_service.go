package masterservice

import (
	"RU-Smart-Workspace/ru-smart-api/repositories/masterrepo"
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type (
	studentServices struct {
		studentRepo masterrepo.StudentRepoInterface
		redis_cache *redis.Client
	}

	AuthenPlayload struct {
		Std_code      string `json:"std_code"`
		Refresh_token string `json:"refresh_token"`
	}

	RegisterPlayload struct {
		Std_code        string `json:"std_code"`
		Course_year     string `json:"course_year"`
		Course_semester string `json:"course_semester"`
	}

	AuthenPlayloadRedirect struct {
		Std_code     string `json:"std_code"`
		Access_token string `json:"access_token"`
	}

	AuthenTestPlayload struct {
		Std_code string `json:"std_code"`
	}

	AuthenServicePlayload struct {
		ServiveId string `json:"service_id"`
	}

	TokenResponse struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
		IsAuth       bool   `json:"isAuth"`
		Message      string `json:"message"`
		StatusCode   int    `json:"status_code"`
	}

	TokenRedirectResponse struct {
		IsAuth     bool   `json:"isAuth"`
		Message    string `json:"message"`
		StdCode    string `json:"std_code"`
		StatusCode int    `json:"status_code"`
	}

	// claims คือข้อมูลที่อยู่ในส่วน Payload ของ Token
	// -iss (issuer) : เว็บหรือบริษัทเจ้าของ token
	// -sub (subject) : subject ของ token
	// -aud (audience) : ผู้รับ token
	// -exp (expiration time) : เวลาหมดอายุของ token
	// -nbf (not before) : เป็นเวลาที่บอกว่า token จะเริ่มใช้งานได้เมื่อไหร่
	// -iat (issued at) : ใช้เก็บเวลาที่ token นี้เกิดปัญหา
	// -jti (JWT id) : เอาไว้เก็บไอดีของ JWT แต่ละตัวนะครับ
	// -name (Full name) : เอาไว้เก็บชื่อ
	ClaimsToken struct {
		Issuer              string `json:"issuer"`
		Subject             string `json:"subject"`
		Role                string `json:"role"`
		ExpiresAccessToken  string `json:"expires_access_token"`
		ExpiresRefreshToken string `json:"expiration_refresh_token"`
	}

	StudentProfileService struct {
		STD_CODE             string `json:"STD_CODE"`
		NAME_THAI            string `json:"NAME_THAI"`
		NAME_ENG             string `json:"NAME_ENG"`
		BIRTH_DATE           string `json:"BIRTH_DATE"`
		STD_STATUS_DESC_THAI string `json:"STD_STATUS_DESC_THAI"`
		CITIZEN_ID           string `json:"CITIZEN_ID"`
		REGIONAL_NAME_THAI   string `json:"REGIONAL_NAME_THAI"`
		STD_TYPE_DESC_THAI   string `json:"STD_TYPE_DESC_THAI"`
		FACULTY_NAME_THAI    string `json:"FACULTY_NAME_THAI"`
		MAJOR_NAME_THAI      string `json:"MAJOR_NAME_THAI"`
		WAIVED_NO            string `json:"WAIVED_NO"`
		WAIVED_PAID          string `json:"WAIVED_PAID"`
		WAIVED_TOTAL_CREDIT  int    `json:"WAIVED_TOTAL_CREDIT"`
		CHK_CERT_NAME_THAI   string `json:"CHK_CERT_NAME_THAI"`
		PENAL_NAME_THAI      string `json:"PENAL_NAME_THAI"`
		MOBILE_TELEPHONE     string `json:"MOBILE_TELEPHONE"`
		EMAIL_ADDRESS        string `json:"EMAIL_ADDRESS"`
	}

	StudentSuccessService struct {
		STD_CODE        string `json:"STD_CODE"`
		NAME_THAI       string `json:"NAME_THAI"`
		NAME_ENG        string `json:"NAME_ENG"`
		YEAR            string `json:"YEAR"`
		SEMESTER        string `json:"SEMESTER"`
		CURR_NAME       string `json:"CURR_NAME"`
		MAJOR_NAME_THAI string `json:"MAJOR_NAME_THAI"`
		MAJOR_NAME      string `json:"MAJOR_NAME"`
		PLAN            string `json:"PLAN"`
		CONFERENCE_NO   string `db:"CONFERENCE_NO"`
		SERIAL_NO       string `json:"SERIAL_NO"`
		CONFERENCE_DATE string `json:"CONFERENCE_DATE"`
		GRADUATED_DATE  string `json:"GRADUATED_DATE"`
		CONFIRM_DATE    string `json:"CONFIRM_DATE"`
	}

	RegisterResponse struct {
		STD_CODE string                 `json:"std_code"`
		YEAR     string                 `json:"year"`
		REGISTER []RegisterResponseRepo `json:"register"`
	}

	GradeResponse struct {
		STD_CODE       string              `json:"STD_CODE"`
		YEAR           string              `json:"YEAR"`
		SUMMARY_CREDIT int                 `json:"SUMMARY_CREDIT"`
		GPA            float32             `json:"GPA"`
		GRADE          []GradeResponseRepo `json:"GRADEDATA"`
	}

	StudentResponse struct {
		STUDENT_CODE string `json:"std_code"`
	}

	RegisterResponseRepo struct {
		YEAR      string `json:"YEAR"`
		SEMESTER  string `json:"SEMESTER"`
		STD_CODE  string `json:"STD_CODE"`
		COURSE_NO string `json:"COURSE_NO"`
		CREDIT    int    `json:"CREDIT"`
	}

	GradeResponseRepo struct {
		YEAR      string `json:"YEAR"`
		SEMESTER  string `json:"SEMESTER"`
		STD_CODE  string `json:"STD_CODE"`
		COURSE_NO string `json:"COURSE_NO"`
		CREDIT    int    `json:"CREDIT"`
		GRADE     string `json:"GRADE"`
	}

	RegisterRequest struct {
		STD_CODE string `json:"std_code" validate:"min=9,max=10,regexp=^[0-9]"`
		YEAR     string `json:"year" validate:"min=4,max=4,regexp=^[0-9]"`
	}

	PrivacyPolicyResponse struct {
		STD_CODE string `db:"STD_CODE"`
		VERSION  string `db:"VERSION"`
		CREATED  string `db:"CREATED"`
		MODIFIED string `db:"MODIFIED"`
	}

	PrivacyPolicyRequest struct {
		STD_CODE string `json:"std_code" validate:"min=9,max=10,regexp=^[0-9]"`
		VERSION  string `json:"version" validate:"regexp=^[0-9]"`
	}

	StudentServicesInterface interface {
		GetStudentProfile(stdCode string) (*StudentProfileService, error)
		GetStudentSuccess(stdCode string) (*StudentSuccessService, error)
		GetStudentSuccessCheck(stdCode string) (*StudentSuccessService, error)

		GetRegisterAll(stdCode string) (*RegisterResponse, error)
		GetRegisterByYear(stdCode, year string) (*RegisterResponse, error)

		GetGradeAll(stdCode string) (*GradeResponse, error)
		GetGradeByYear(stdCode, year string) (*GradeResponse, error)

		SetPrivacyPolicy(privacyPolicyRequest PrivacyPolicyRequest) (*PrivacyPolicyResponse, error)
	}
)

func NewStudentServices(studentRepo masterrepo.StudentRepoInterface, redis_cache *redis.Client) StudentServicesInterface {
	return &studentServices{
		studentRepo: studentRepo,
		redis_cache: redis_cache,
	}
}
