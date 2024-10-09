package students

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"net/http"
)

func (s *studentServices) Certificate(stdCode, certificate string) (*TokenCertificateResponse, error) {

	certificateTokenResponse := TokenCertificateResponse{
		AccessToken: "",
		StartDate:   "",
		ExpireDate:  "",
		Certificate: "",
		Message:     "",
		StatusCode:  422,
	}

	generateToken, err := middlewares.GenerateTokenCertificate(stdCode, certificate, s.redis_cache)
	if err != nil {
		certificateTokenResponse.Message = "Certificate Generate Certificate fail."
		return &certificateTokenResponse, err
	}

	certificateTokenResponse.AccessToken = generateToken.AccessToken
	certificateTokenResponse.StartDate = generateToken.StartDate
	certificateTokenResponse.ExpireDate = generateToken.ExpireDate
	certificateTokenResponse.Certificate = generateToken.Certificate
	certificateTokenResponse.Message = "Generate Certificate success..."
	certificateTokenResponse.StatusCode = http.StatusOK

	return &certificateTokenResponse, nil
}
