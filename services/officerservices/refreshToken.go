package officerservices

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"fmt"
	"net/http"
)

func (s *officerServices) RefreshAuthenticationOfficer(refreshToken string) (*AuthenResponse, error) {

	authenTokenResponse := AuthenResponse{
		AccessToken:  "",
		RefreshToken: "RefreshToken",
		IsAuth:       false,
		Message:      "",
		StatusCode:   403,
	}

	// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
	isToken, err := middlewares.VerifyTokenOfficer("refreshToken", refreshToken, s.redis_cache)
	if err != nil && !isToken {
		authenTokenResponse.Message = "You are not officer authentication."
		return &authenTokenResponse, err
	}

	isRevokeToken := middlewares.RevokeTokenOfficer(refreshToken, s.redis_cache)
	if !isRevokeToken {
		authenTokenResponse.Message = "Don't revoke token officer becourse Not found."
		return &authenTokenResponse, err
	}

	claimsToken, err := middlewares.GetClaimsOfficer(refreshToken)
	if err != nil {
		authenTokenResponse.Message = "Don't claim token officer becourse Not valid."
		return &authenTokenResponse, err
	}

	fmt.Println(claimsToken.Officer)

	authenRepo, err := s.officerRepo.GetUserLogin(claimsToken.Officer)
	if err != nil || authenRepo.Status != 1 {
		authenTokenResponse.Message = "Don't Authenticated token becourse Not found email in database."
		return &authenTokenResponse, err
	}

	fmt.Println(authenRepo.Role)

	generateToken, err := middlewares.GenerateTokenOfficer(authenRepo.Username, authenRepo.Role, s.redis_cache)
	if err != nil {
		authenTokenResponse.Message = "Refresh and Generate token officer fail."
		return &authenTokenResponse, err
	}

	authenTokenResponse.AccessToken = generateToken.AccessToken
	authenTokenResponse.RefreshToken = generateToken.RefreshToken
	authenTokenResponse.IsAuth = generateToken.IsAuth
	authenTokenResponse.Message = "Refresh and Generate token officer success."
	authenTokenResponse.StatusCode = http.StatusOK

	return &authenTokenResponse, nil

}
