package masterservice

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"fmt"
	"net/http"
)

func (s *studentServices) Certificate(ID_TOKEN string) (*TokenCertificateResponse, error) {

	certificateTokenResponse := TokenCertificateResponse{
		CertificateToken: "",
		StartDate:        "",
		ExpireDate:       "",
		Certificate:      "",
		Message:          "",
		StatusCode:       422,
	}

	claimsToken, err := middlewares.GetClaims(ID_TOKEN)
	if err != nil {
		certificateTokenResponse.Message = "Don't claim token becourse Not valid."
		return &certificateTokenResponse, err
	}

	fmt.Println(claimsToken.StudentCode)

	generateToken, err := middlewares.GenerateTokenCertificate(ID_TOKEN, claimsToken.StudentCode, "egraduate", s.redis_cache)
	if err != nil {
		certificateTokenResponse.Message = "Certificate Generate Certificate fail."
		return &certificateTokenResponse, err
	}

	certificateTokenResponse.CertificateToken = generateToken.CertificateToken
	certificateTokenResponse.StartDate = generateToken.StartDate
	certificateTokenResponse.ExpireDate = generateToken.ExpireDate
	certificateTokenResponse.Certificate = generateToken.Certificate
	certificateTokenResponse.Message = "Generate Certificate success..."
	certificateTokenResponse.StatusCode = http.StatusOK

	return &certificateTokenResponse, nil
}
