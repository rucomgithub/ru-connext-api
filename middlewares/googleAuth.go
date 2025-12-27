package middlewares

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"context"
	"errors"
	"net/http"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/oauth2/v1"
	"google.golang.org/api/option"

	"google.golang.org/api/idtoken"
)

var ctx = context.Background()

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
		return nil, errors.New("id token is empty")
	}

	log.Printf("idToken: %s", idToken)

	ctx := context.Background()

	payload, err := idtoken.Validate(
		ctx,
		idToken,
		"668594026369-pbki4bj9l02svr8412ahgnhu99vpi0k3.apps.googleusercontent.com", // Web Client ID
	)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
