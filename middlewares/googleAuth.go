package middlewares

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/oauth2/v1"
	"google.golang.org/api/option"

	"google.golang.org/api/idtoken"
)

var ctx = context.Background()

func GetHeaderAuthorizationToken(c *gin.Context) (string, error) {
	const bearerSchema = "Bearer "

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is empty")
	}

	// Trim all whitespace
	authHeader = strings.TrimSpace(authHeader)

	if !strings.HasPrefix(authHeader, bearerSchema) {
		return "", errors.New("authorization header must start with Bearer")
	}

	// Extract token
	token := authHeader[len(bearerSchema):]

	// 🧹 Clean the token thoroughly
	token = strings.TrimSpace(token)
	token = strings.ReplaceAll(token, "\n", "")
	token = strings.ReplaceAll(token, "\r", "")
	token = strings.ReplaceAll(token, "\t", "")
	token = strings.ReplaceAll(token, " ", "")

	// Validate JWT structure
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		log.Printf("❌ Invalid JWT structure: expected 3 parts, got %d", len(parts))
		return "", errors.New("invalid JWT format: must have 3 parts separated by dots")
	}

	// Validate each part is not empty
	for i, part := range parts {
		if part == "" {
			log.Printf("❌ JWT part %d is empty", i)
			return "", errors.New("invalid JWT format: empty part detected")
		}
	}

	log.Printf("✅ Valid JWT structure: %d.%d.%d characters",
		len(parts[0]), len(parts[1]), len(parts[2]))

	return token, nil
}

func GoogleAuth(c *gin.Context) {

	ID_TOKEN, err := GetHeaderAuthorization(c)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"accessToken": "", "isAuth": false, "message": "authorization key in header not found"})
		c.Abort()
		return
	}

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	log.Println("ID TOKEN:", ID_TOKEN)

	_, err = verifyGoogleAuth(ID_TOKEN)
	if err != nil {
		c.Error(err)
		c.Set("line", handlers.GetLineNumber())
		c.Set("file", handlers.GetFileName())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"accessToken": "", "isAuth": false, "message": "Google is not authorized" + err.Error()})
		c.Abort()
		return
	}
	c.Next()

}

func verifyGoogleAuth(id_token string) (*oauth2.Tokeninfo, error) {

	timeout := time.Duration(5 * time.Second)
	httpClient := &http.Client{Timeout: timeout}

	oauth2Service, err := oauth2.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err
	}

	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.AccessToken(id_token)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

func verifyGoogleAuthIDToken(idToken string) (*idtoken.Payload, error) {
	if idToken == "" {
		return nil, errors.New("ID token is empty")
	}

	log.Printf("JWT TOKEN: %s", idToken)

	ctx := context.Background()

	// 🔍 FIRST: Validate WITHOUT audience to see what's in the token
	payload, err := idtoken.Validate(ctx, idToken, "")
	if err != nil {
		log.Printf("Error validating ID token: %v", err)
		return nil, err
	}

	// 📝 Log the actual audience claim
	log.Printf("🔍 Token AUD claim: %v", payload.Audience)
	log.Printf("🔍 Token ISS claim: %v", payload.Issuer)
	log.Printf("🔍 Token SUB claim: %v", payload.Subject)
	log.Printf("🔍 Token EMAIL: %v", payload.Claims["email"])

	return payload, nil
}
