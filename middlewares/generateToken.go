package middlewares

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func GenerateToken(stdCode string, redis_cache *redis.Client) (*TokenResponse, error) {

	generateToken := &TokenResponse{}
	//expirationAccessToken := time.Now().AddDate(0, 0, 1).Unix()
	expirationAccessToken := time.Now().Add(time.Minute * 60).Unix()
	//expirationRefreshToken := time.Now().AddDate(0, 1, 0).Unix()
	expirationRefreshToken := time.Now().Add(time.Minute * 60 * 6).Unix()

	generateToken.IsAuth = true
	generateToken.AccessTokenKey = stdCode + "::access::" + uuid.New().String()
	generateToken.RefreshTokenKey = stdCode + "::refresh::" + uuid.New().String()

	// ---------------------  Create Access Token  ----------------------------------------- //
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["issuer"] = viper.GetString("token.issuer")
	accessTokenClaims["subject"] = "Ru-Connext" + stdCode
	accessTokenClaims["role"] = ""
	accessTokenClaims["expires_token"] = expirationAccessToken
	accessTokenClaims["access_token_key"] = generateToken.AccessTokenKey
	accessTokenClaims["refresh_token_key"] = generateToken.RefreshTokenKey

	accessTokenHeader := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	NEW_ACCESS_TOKEN, err := accessTokenHeader.SignedString([]byte(viper.GetString("token.secretKey")))
	if err != nil {
		return nil, err
	}

	generateToken.AccessToken = NEW_ACCESS_TOKEN

	// ---------------------  Create Refresh Token  ----------------------------------------- //
	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["issuer"] = viper.GetString("token.issuer")
	refreshTokenClaims["subject"] = "Ru-Connext::" + stdCode
	refreshTokenClaims["role"] = ""
	refreshTokenClaims["expires_token"] = expirationRefreshToken
	refreshTokenClaims["access_token_key"] = generateToken.AccessTokenKey
	refreshTokenClaims["refresh_token_key"] = generateToken.RefreshTokenKey

	refreshTokenHeader := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	NEW_REFRESH_TOKEN, err := refreshTokenHeader.SignedString([]byte(viper.GetString("token.secretKey")))
	if err != nil {
		return nil, err
	}

	generateToken.RefreshToken = NEW_REFRESH_TOKEN

	// ---------------------------  redis cache database  ------------------------------------ //
	// เริ่มนับเวลา ณ ตอนนี้
	timeNow := time.Now()

	cacheStudent := CacheStudent{
		StdCode: stdCode,
		Role:    "something message",
	}

	cacheDataJson, _ := json.Marshal(cacheStudent)

	redisCacheExpiresAccessToken := time.Unix(expirationAccessToken, 0)
	err = redis_cache.Set(ctx, fmt.Sprint(generateToken.AccessTokenKey), cacheDataJson, redisCacheExpiresAccessToken.Sub(timeNow)).Err()
	if err != nil {
		return nil, err
	}

	redisCacheExpiresRefreshToken := time.Unix(expirationRefreshToken, 0)
	err = redis_cache.Set(ctx, fmt.Sprint(generateToken.RefreshTokenKey), cacheDataJson, redisCacheExpiresRefreshToken.Sub(timeNow)).Err()
	if err != nil {
		return nil, err
	}

	return generateToken, nil
}
