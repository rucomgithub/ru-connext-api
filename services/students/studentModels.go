package students

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"RU-Smart-Workspace/ru-smart-api/repositories/studentr"
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type (
	studentServices struct {
		studentRepo studentr.StudentRepoInterface
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

	TokenCertificateResponse struct {
		AccessToken      string `json:"accessToken"`
		CertificateToken string `json:"certificateToken"`
		StartDate        string `json:"startDate"`
		ExpireDate       string `json:"expireDate"`
		Certificate      string `json:"certificate"`
		Message          string `json:"message"`
		StatusCode       int    `json:"status_code"`
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

	RegisterResponse struct {
		STUDENT_CODE    string                   `json:"std_code"`
		COURSE_YEAR     string                   `json:"course_year"`
		COURSE_SEMESTER string                   `json:"course_semester"`
		REGISTER        []RegisterResponseFromDB `json:"register"`
	}

	StudentResponse struct {
		STUDENT_CODE string `json:"std_code"`
	}

	RegisterResponseFromDB struct {
		ID                   string `json:"id"`
		COURSE_YEAR          string `json:"course_year"`
		COURSE_SEMESTER      string `json:"course_semester"`
		COURSE_NO            string `json:"course_no"`
		COURSE_METHOD        string `json:"course_method"`
		COURSE_METHOD_NUMBER string `json:"course_method_number"`
		DAY_CODE             string `json:"day_code"`
		TIME_CODE            string `json:"time_code"`
		ROOM_GROUP           string `json:"room_group"`
		INSTR_GROUP          string `json:"instr_group"`
		COURSE_METHOD_DETAIL string `json:"course_method_detail"`
		DAY_NAME_S           string `json:"day_name_s"`
		TIME_PERIOD          string `json:"time_period"`
		COURSE_ROOM          string `json:"course_room"`
		COURSE_INSTRUCTOR    string `json:"course_instructor"`
		SHOW_RU30            string `json:"show_ru30"`
		COURSE_CREDIT        string `json:"course_credit"`
		COURSE_PR            string `json:"course_pr"`
		COURSE_COMMENT       string `json:"course_comment"`
		COURSE_EXAMDATE      string `json:"course_examdate"`
	}

	StudentServicesInterface interface {
		Certificate(token string) (*TokenCertificateResponse, error)
		Authentication(stdCode string) (*TokenResponse, error)
		AuthenticationService(service_id string) (*TokenResponse, error)
		AuthenticationRedirect(stdCode, accessToken string) (*TokenRedirectResponse, error)
		RefreshAuthentication(refreshToken string) (*TokenResponse, error)
		Unauthorization(token string) bool
		CheckExistsToken(token string) bool
		CheckToken(token string) (*middlewares.ClaimsToken, error)
		GetStudentProfile(stdCode string) (*StudentProfileService, error)
		GetRegister(studentCode, courseYear, courseSemester string) (*RegisterResponse, error)
		GetRegisterAll(studentCode, courseYear string) (*RegisterAllResponse, error)

		GetStudentAll() (*[]StudentResponse, error)
	}
)

func NewStudentServices(studentRepo studentr.StudentRepoInterface, redis_cache *redis.Client) StudentServicesInterface {
	return &studentServices{
		studentRepo: studentRepo,
		redis_cache: redis_cache,
	}
}
