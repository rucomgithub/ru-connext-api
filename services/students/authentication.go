package students

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"net/http"
)

func (s *studentServices) Authentication(stdCode string) (*TokenResponse, error) {

	studentTokenResponse := TokenResponse{
		AccessToken:  "",
		RefreshToken: "",
		IsAuth:       false,
		Message:      "",
		StatusCode:   422,
	}

	prepareToken, err := s.studentRepo.Authentication(stdCode)
	if err != nil || prepareToken.STATUS != 1 {
		studentTokenResponse.Message = "สถานะภาพการเป็นนักศึกษาของท่าน (จบการศึกษา,หมดสถานภาพ,ขาดการลงทะเบียนเรียน 2 ภาคการศึกษาขึ้นไป)."
		return &studentTokenResponse, err
	}

	generateToken, err := middlewares.GenerateToken(prepareToken.STD_CODE, s.redis_cache)
	if err != nil {
		studentTokenResponse.Message = "Authentication Generate Token fail."
		return &studentTokenResponse, err
	}

	studentTokenResponse.AccessToken = generateToken.AccessToken
	studentTokenResponse.RefreshToken = generateToken.RefreshToken
	studentTokenResponse.IsAuth = generateToken.IsAuth
	studentTokenResponse.Message = "Generate Token success..."
	studentTokenResponse.StatusCode = http.StatusOK

	return &studentTokenResponse, nil
}

func (s *studentServices) AuthenticationService(service_id string) (*TokenResponse, error) {

	studentTokenResponse := TokenResponse{
		AccessToken:  "",
		RefreshToken: "",
		IsAuth:       false,
		Message:      "",
		StatusCode:   422,
	}

	// prepareToken, err := s.studentRepo.Authentication(service_id)
	// if err != nil || prepareToken.STATUS != 1 {
	// 	studentTokenResponse.Message = "สถานะภาพการเป็นนักศึกษาของท่าน (จบการศึกษา,หมดสถานภาพ,ขาดการลงทะเบียนเรียน 2 ภาคการศึกษาขึ้นไป)."
	// 	return &studentTokenResponse, err
	// }

	generateToken, err := middlewares.GenerateServiceToken(service_id, s.redis_cache)
	if err != nil {
		studentTokenResponse.Message = "Authentication Generate Token fail."
		return &studentTokenResponse, err
	}

	studentTokenResponse.AccessToken = generateToken.AccessToken
	studentTokenResponse.RefreshToken = generateToken.RefreshToken
	studentTokenResponse.IsAuth = generateToken.IsAuth
	studentTokenResponse.Message = "Generate Service Token success..."
	studentTokenResponse.StatusCode = http.StatusOK

	return &studentTokenResponse, nil
}
