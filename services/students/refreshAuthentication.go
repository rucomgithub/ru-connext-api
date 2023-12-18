package students

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"net/http"
)

func (s *studentServices) RefreshAuthentication(refreshToken, stdCode string) (*TokenResponse, error) {

	studentTokenResponse := TokenResponse{
		AccessToken:  "",
		RefreshToken: "",
		IsAuth:       false,
		Message:      "",
		StatusCode:   422,
	}

	// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
	isToken, err := middlewares.VerifyToken("refreshToken", refreshToken, s.redis_cache)
	if err != nil && !isToken {
		studentTokenResponse.Message = "You are not authentication."
		return &studentTokenResponse, err
	}

	isRevokeToken := middlewares.RevokeToken(refreshToken, s.redis_cache)
	if !isRevokeToken {
		studentTokenResponse.Message = "Don't revoke token becourse Not found."
		return &studentTokenResponse, err
	}

	prepareToken, err := s.studentRepo.Authentication(stdCode)
	if err != nil || prepareToken.STATUS != 1 {
		studentTokenResponse.Message = "Don't Authenticated token becourse Not found student code in database."
		return &studentTokenResponse, err
	}

	generateToken, err := middlewares.GenerateToken(prepareToken.STD_CODE, s.redis_cache)
	if err != nil {
		studentTokenResponse.Message = "Refresh and Generate Token fail."
		return &studentTokenResponse, err
	}

	studentTokenResponse.AccessToken = generateToken.AccessToken
	studentTokenResponse.RefreshToken = generateToken.RefreshToken
	studentTokenResponse.IsAuth = generateToken.IsAuth
	studentTokenResponse.Message = "Refresh and Generate Token success."
	studentTokenResponse.StatusCode = http.StatusOK

	return &studentTokenResponse, nil

}
