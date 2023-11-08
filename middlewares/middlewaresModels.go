package middlewares

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type (
	CacheStudent struct {
		StdCode string `json:"std_code"`
		Role    string `json:"std_role"`
	}

	CacheService struct {
		ServiceId string `json:"service_id"`
		Role      string `json:"service_role"`
	}

	TokenResponse struct {
		AccessToken     string `json:"accessToken"`
		RefreshToken    string `json:"refreshToken"`
		IsAuth          bool   `json:"isAuth"`
		AccessTokenKey  string `json:"access_token_key"`
		RefreshTokenKey string `json:"refresh_token_key"`
	}

	ClaimsToken struct {
		Issuer          string `json:"issuer"`
		Subject         string `json:"subject"`
		Role            string `json:"role"`
		StudentCode     string `json:"std_code"`
		ExpiresToken    string `json:"expires_token"`
		AccessTokenKey  string `json:"access_token_key"`
		RefreshTokenKey string `json:"refresh_token_key"`
	}

	RefreshAuthen struct {
		// StdCode      string `json:"std_code"`
		RefreshToken string `json:"refresh_token"`
	}
)

// ทำการแกะ header HTTP request
// Authorization: Bearer TOKEN
func GetHeaderAuthorization(c *gin.Context) (token string, err error) {

	const BEARER_SCHEMA = "Bearer "
	AUTH_HEADER := c.GetHeader("Authorization")

	if len(AUTH_HEADER) == 0 {
		return "", err
	}

	if strings.HasPrefix(AUTH_HEADER, BEARER_SCHEMA) {
		token = AUTH_HEADER[len(BEARER_SCHEMA):]
		return token, nil
	} else {
		return "", err
	}

}

func VerifyToken(preTokenKey string, token string, redis_cache *redis.Client) (bool, error) {

	claims, err := GetClaims(token)
	if err != nil {
		return false, err
	}

	if preTokenKey == "accessToken" {
		_, err = redis_cache.Get(ctx, claims.AccessTokenKey).Result()
	} else {
		_, err = redis_cache.Get(ctx, claims.RefreshTokenKey).Result()
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func CheckExistsToken(token string, redis_cache *redis.Client) (bool, error) {

	claims, err := GetClaims(token)
	if err != nil {
		return false, err
	}

	_, err = redis_cache.Get(ctx, claims.AccessTokenKey).Result()

	if err != nil {
		return false, err
	}

	return true, nil
}

func GetClaims(encodedToken string) (*ClaimsToken, error) {

	parseToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("token.secretKey")), nil
	})
	if err != nil {
		return nil, err
	}

	claimsToken := &ClaimsToken{}
	parseClaims := parseToken.Claims.(jwt.MapClaims)

	if parseClaims["issuer"] != nil {
		claimsToken.Issuer = parseClaims["issuer"].(string)
	}

	if parseClaims["subject"] != nil {
		claimsToken.Subject = parseClaims["subject"].(string)
	}

	if parseClaims["role"] != "" {
		claimsToken.Role = parseClaims["role"].(string)
	} else {
		claimsToken.Role = ""
	}

	if parseClaims["std_code"] != "" {
		claimsToken.StudentCode = parseClaims["std_code"].(string)
	} else {
		claimsToken.StudentCode = ""
	}

	if parseClaims["access_token_key"] != nil {
		claimsToken.AccessTokenKey = parseClaims["access_token_key"].(string)
	}

	if parseClaims["refresh_token_key"] != nil {
		claimsToken.RefreshTokenKey = parseClaims["refresh_token_key"].(string)
	}

	if parseClaims["expires_token"] != nil {
		claimsToken.ExpiresToken = fmt.Sprintf("%v", parseClaims["expires_token"])
	}

	return claimsToken, nil
}

func RevokeToken(token string, redis_cache *redis.Client) bool {

	claims, err := GetClaims(token)
	if err != nil {
		return false
	}

	redis_cache.Del(ctx, claims.AccessTokenKey).Result()
	redis_cache.Del(ctx, claims.RefreshTokenKey).Result()

	return true
}
