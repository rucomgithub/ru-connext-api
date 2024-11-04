package officerservices

import (
	"RU-Smart-Workspace/ru-smart-api/repositories/officerrepos"

	"github.com/go-redis/redis/v8"
)

type (
	officerServices struct {
		officerRepo officerrepos.OfficerRepoInterface
		redis_cache *redis.Client
	}

	AuthenRequest struct {
		Username string `json:"Username" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}

	AuthenResponse struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
		IsAuth       bool   `json:"isAuth"`
		Message      string `json:"message"`
		StatusCode   int    `json:"status_code"`
	}

	TokenOffice struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    string `json:"expires_in"`
		ExpiresOn    string `json:"expires_on"`
		ExtExpiresIn string `json:"ext_expires_in"`
		NotBefore    string `json:"not_before"`
		RefreshToken string `json:"refresh_token"`
		Resource     string `json:"resource"`
		Scope        string `json:"scope"`
		TokenType    string `json:"token_type"`
	}

	TokenOfficeError struct {
		CorrelationID    string  `json:"correlation_id"`
		Error            string  `json:"error"`
		ErrorCodes       []int64 `json:"error_codes"`
		ErrorDescription string  `json:"error_description"`
		ErrorURI         string  `json:"error_uri"`
		Timestamp        string  `json:"timestamp"`
		TraceID          string  `json:"trace_id"`
	}

	User struct {
		Odata_context     string      `json:"@odata.context"`
		BusinessPhones    []string    `json:"businessPhones"`
		DisplayName       string      `json:"displayName"`
		GivenName         string      `json:"givenName"`
		ID                string      `json:"id"`
		JobTitle          string      `json:"jobTitle"`
		Mail              string      `json:"mail"`
		MobilePhone       string      `json:"mobilePhone"`
		OfficeLocation    string      `json:"officeLocation"`
		PreferredLanguage interface{} `json:"preferredLanguage"`
		Surname           string      `json:"surname"`
		UserPrincipalName string      `json:"userPrincipalName"`
		AccessToken       string      `json:"access_token"`
		RefreshToken      string      `json:"refresh_token"`
		RoleToken         string      `json:"role_token"`
		ExpiresTime       int64       `json:"expiretime"`
	}

	UserLogin struct {
		Username string `json:"username"`
		Role     string `json:"role"`
		Key      string `json:"key"`
		Created  string `json:"created"`
		Modified string `json:"modified"`
		Status   string `json:"status"`
	}

	QualificationRequest struct {
		STATUS      string `json:"status" validate:"required"`
		DESCRIPTION string `json:"description"`
	}

	QualificationResponse struct {
		STD_CODE     string `json:"std_code"`
		REQUEST_DATE string `json:"request_date"`
		OPERATE_DATE string `json:"operate_date"`
		CONFIRM_DATE string `json:"confirm_date"`
		CANCEL_DATE  string `json:"cancel_date"`
		STATUS       string `json:"status"`
		CREATED      string `json:"created"`
		MODIFIED     string `json:"modified"`
		DESCRIPTION  string `json:"description"`
	}

	OfficerServiceInterface interface {
		AuthenticationOfficer(authenRequest AuthenRequest) (*AuthenResponse, error)
		RefreshAuthenticationOfficer(refreshToken string) (*AuthenResponse, error)
		VerifyAuthentication(Username string, Password string) (TokenOffice, error)

		GetQualificationAll() (*[]QualificationResponse, error)
		GetQualification(std_code string) (*QualificationResponse, error)
		UpdateQualification(std_code, status, description string) (*QualificationResponse, int64, error)
	}
)

func NewOfficerServices(officerRepo officerrepos.OfficerRepoInterface, redis_cache *redis.Client) OfficerServiceInterface {
	return &officerServices{
		officerRepo: officerRepo,
		redis_cache: redis_cache,
	}
}
