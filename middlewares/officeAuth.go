package middlewares

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/oauth2/v1"
	"google.golang.org/api/option"
)

func OfficeAuth(c *gin.Context) {

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
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"accessToken": "", "isAuth": false, "message": "Office is not authorized"})
		c.Abort()
		return
	}
	c.Next()

}

func verifyOfficeAuth(id_token string) (*oauth2.Tokeninfo, error) {

	timeout := time.Duration(5 * time.Second)
	httpClient := &http.Client{Timeout: timeout}

	oauth2Service, err := oauth2.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err
	}

	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(id_token)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
