package students

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"net/http"
)

func (s *studentServices) AuthenticationRedirect(stdCode string, accessToken string) (*TokenRedirectResponse, error) {

	studentTokenResponse := TokenRedirectResponse{
		IsAuth:     false,
		Message:    "",
		StdCode:    "",
		StatusCode: 422,
	}

	isToken, err := middlewares.VerifyToken("accessToken", accessToken, s.redis_cache)
	if err != nil {
		studentTokenResponse.Message = "Authorization fail..."
		return &studentTokenResponse, err
	}

	studentTokenResponse.IsAuth = isToken
	studentTokenResponse.Message = "Authorization success..."
	studentTokenResponse.StdCode = stdCode
	studentTokenResponse.StatusCode = http.StatusOK

	return &studentTokenResponse, nil
}
