package students
 
import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"net/http"
	"log"
	"context"
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "strings"
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

	generateToken, err := middlewares.GenerateToken(prepareToken.STD_CODE,prepareToken.ROLE, s.redis_cache)
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


func (s *studentServices) VerifyToken(ctx context.Context, token string) (*GoogleTokenInfo, error) {
    url := fmt.Sprintf("https://oauth2.googleapis.com/tokeninfo?id_token=%s", token)

    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to verify token: %w", err)
    }
    defer resp.Body.Close()

	log.Println(resp)

    if resp.StatusCode != http.StatusOK {
        body, _ := ioutil.ReadAll(resp.Body)
        return nil, fmt.Errorf("token verification failed: %s", string(body))
    }

    var tokenInfo GoogleTokenInfo
    if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
        return nil, fmt.Errorf("failed to decode token info: %w", err)
    }

	log.Println(tokenInfo)

    // Verify email domain
    if !strings.HasSuffix(tokenInfo.Email, "@rumail.ru.ac.th") {
        return nil, errors.New("invalid email domain")
    }

    // Verify email is verified
    if tokenInfo.EmailVerified == "true" {
        return nil, errors.New("email not verified")
    }

    return &tokenInfo, nil
}

// ExtractStudentID extracts student ID from email
func (s *studentServices) ExtractStudentID(email string) string { 
    parts := strings.Split(email, "@")
    if len(parts) > 0 {
        return parts[0]
    }
    return ""
}