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

func GenerateTokenCertificate(ID_TOKEN, stdCode, certificate string, redis_cache *redis.Client) (*TokenCertificateResponse, error) {

	generateToken := &TokenCertificateResponse{}
	//expirationAccessToken := time.Now().AddDate(0, 0, 1).Unix()
	timeStart := time.Now()
	timeExpire := timeStart.Add(time.Minute * 60)
	expirationAccessToken := timeExpire.Unix()

	generateToken.AccessTokenKey = stdCode + "::certificate::" + uuid.New().String()
	generateToken.StartDate = timeStart.String()
	generateToken.ExpireDate = timeExpire.String()
	generateToken.Certificate = certificate + "::certificate::" + uuid.New().String()

	// ---------------------  Create Access Token  ----------------------------------------- //
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["issuer"] = viper.GetString("token.issuer") + certificate
	accessTokenClaims["subject"] = "Certificate::" + stdCode
	accessTokenClaims["certificate"] = certificate
	accessTokenClaims["start_date"] = timeStart
	accessTokenClaims["expire_date"] = timeExpire
	accessTokenClaims["std_code"] = stdCode
	accessTokenClaims["access_token"] = ID_TOKEN
	accessTokenClaims["expires_token"] = expirationAccessToken
	accessTokenClaims["access_token_key"] = generateToken.AccessTokenKey

	accessTokenHeader := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	NEW_ACCESS_TOKEN, err := accessTokenHeader.SignedString([]byte(viper.GetString("token.certificateKey")))
	if err != nil {
		return nil, err
	}

	generateToken.CertificateToken = NEW_ACCESS_TOKEN

	// ---------------------------  redis cache database  ------------------------------------ //
	// เริ่มนับเวลา ณ ตอนนี้
	timeNow := time.Now()

	cacheCertificate := CacheCertificate{
		StdCode:     stdCode,
		Certificate: certificate,
	}

	cacheDataJson, _ := json.Marshal(cacheCertificate)

	redisCacheExpiresAccessToken := time.Unix(expirationAccessToken, 0)
	err = redis_cache.Set(ctx, fmt.Sprint(generateToken.AccessTokenKey), cacheDataJson, redisCacheExpiresAccessToken.Sub(timeNow)).Err()
	if err != nil {
		return nil, err
	}

	return generateToken, nil
}
