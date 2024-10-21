package middlewares

import (
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func VerifyTokenOfficer(preTokenKey string, token string, redis_cache *redis.Client) (bool, error) {

	claims, err := GetClaimsOfficer(token)
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

func RevokeTokenOfficer(token string, redis_cache *redis.Client) bool {

	claims, err := GetClaimsOfficer(token)
	if err != nil {
		return false
	}

	redis_cache.Del(ctx, claims.AccessTokenKey).Result()
	redis_cache.Del(ctx, claims.RefreshTokenKey).Result()

	return true
}

func GetClaimsOfficer(encodedToken string) (*ClaimsTokenOfficer, error) {

	parseToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("token.officerKey")), nil
	})
	if err != nil {
		return nil, err
	}

	claimsToken := &ClaimsTokenOfficer{}
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

	if parseClaims["officer"] != "" {
		claimsToken.Officer = parseClaims["officer"].(string)
	} else {
		claimsToken.Officer = ""
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
