package officerservices

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

var ctx = context.Background()

func (s *officerServices) AuthenticationOfficer(authenRequest AuthenRequest) (*AuthenResponse, error) {

	authenTokenResponse := AuthenResponse{
		AccessToken:  "",
		RefreshToken: "",
		IsAuth:       false,
		Message:      "",
		StatusCode:   422,
	}

	_, err := s.VerifyAuthentication(authenRequest.Username, authenRequest.Password)

	if err != nil {
		return nil, err
	}

	userrole, err := s.officerRepo.GetUserLogin(authenRequest.Username)
	if err != nil {
		return nil, err
	}

	generateToken, err := middlewares.GenerateTokenOfficer(userrole.Username, userrole.Role, s.redis_cache)

	if err != nil {
		return nil, err
	}

	if err != nil {
		authenTokenResponse.Message = "Authentication Generate token officer fail."
		return &authenTokenResponse, err
	}

	authenTokenResponse.AccessToken = generateToken.AccessToken
	authenTokenResponse.RefreshToken = generateToken.RefreshToken
	authenTokenResponse.IsAuth = generateToken.IsAuth
	authenTokenResponse.Message = "Authentication Generate token officer success."
	authenTokenResponse.StatusCode = http.StatusOK

	return &authenTokenResponse, nil
}

func (s *officerServices) VerifyUser(token TokenOffice) (User, error) {

	var user User

	timeout := time.Duration(5 * time.Second)

	client := &http.Client{
		Timeout: timeout,
	}

	surl := "https://graph.microsoft.com/v1.0/me"

	req, err := http.NewRequest("GET", surl, nil)
	req.Header.Set("Authorization", token.AccessToken)

	if err != nil {
		return user, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return user, err
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}

	resp.Body.Close()
	if err != nil {
		return user, err
	}

	if resp.StatusCode != 200 {
		var tokenerror TokenOfficeError
		err = json.Unmarshal(f, &tokenerror)
		if err != nil {
			return user, err
		}
		return user, err
	}

	err = json.Unmarshal(f, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *officerServices) VerifyAuthentication(Username string, Password string) (TokenOffice, error) {
	timeout := time.Duration(22 * time.Second)

	token := TokenOffice{}

	client := &http.Client{
		Timeout: timeout,
	}

	data := url.Values{}
	data.Add("client_secret", viper.GetString("office.client_secret"))
	data.Add("client_id", viper.GetString("office.client_id"))
	data.Add("grant_type", "password")
	data.Add("resource", "https://graph.microsoft.com")
	data.Add("username", Username)
	data.Add("password", Password)

	surl := "https://login.microsoftonline.com/" + viper.GetString("office.tenant_id") + "/oauth2/token"

	req, err := http.NewRequest("POST", surl, bytes.NewBufferString(data.Encode()))
	//req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") // This makes it work
	if err != nil {
		return token, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return token, err
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}

	resp.Body.Close()

	if resp.StatusCode != 200 {
		var tokenerror TokenOfficeError
		err = json.Unmarshal(f, &tokenerror)

		if err != nil {
			return token, err
		}

		return token, errors.New("Error validating credentials due to invalid username or password." + tokenerror.ErrorDescription)
	}

	err = json.Unmarshal(f, &token)

	if err != nil {
		return token, err
	}

	return token, nil
}
