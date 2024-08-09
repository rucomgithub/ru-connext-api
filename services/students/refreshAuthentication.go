package students

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"net/http"
	"fmt"
)

func (s *studentServices) RefreshAuthentication(refreshToken string) (*TokenResponse, error) {

	studentTokenResponse := TokenResponse{
		AccessToken:  "",
		RefreshToken: "",
		IsAuth:       false,
		Message:      "",
		StatusCode:   403,
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

	claimsToken, err := middlewares.GetClaims(refreshToken)
	if err != nil {
		studentTokenResponse.Message = "Don't claim token becourse Not valid."
		return &studentTokenResponse, err
	}

	fmt.Println(claimsToken.StudentCode)

	prepareToken, err := s.studentRepo.Authentication(claimsToken.StudentCode)
	if err != nil || prepareToken.STATUS != 1 {
		studentTokenResponse.Message = "Don't Authenticated token becourse Not found student code in database."
		return &studentTokenResponse, err
	}

	fmt.Println(prepareToken.ROLE)

	generateToken, err := middlewares.GenerateToken(prepareToken.STD_CODE,prepareToken.ROLE, s.redis_cache)
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
