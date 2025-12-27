package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

var ctx = context.Background()

func GoogleAuth(c *gin.Context) {

	idToken, err := GetHeaderAuthorization(c)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"isAuth":  false,
			"message": "authorization header not found",
		})
		c.Abort()
		return
	}

	payload, err := verifyGoogleAuth(idToken)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"isAuth":  false,
			"message": "google token invalid",
		})
		c.Abort()
		return
	}

	// ✅ เก็บข้อมูล user ไว้ใช้ต่อใน handler
	c.Set("google_sub", payload.Subject)
	c.Set("email", payload.Claims["email"])
	c.Set("name", payload.Claims["name"])

	c.Next()
}

func verifyGoogleAuth(idToken string) (*idtoken.Payload, error) {

	if idToken == "" {
		return nil, errors.New("id token is empty")
	}

	ctx := context.Background()

	payload, err := idtoken.Validate(
		ctx,
		idToken,
		"505837308278-q7h4ihqtu9quajaotvtejhosb2pepmmt.apps.googleusercontent.com", // 👈 Web Client ID
	)
	if err != nil {
		return nil, err
	}

	// (optional) เช็คเพิ่ม
	if payload.Issuer != "https://accounts.google.com" &&
		payload.Issuer != "accounts.google.com" {
		return nil, errors.New("invalid issuer")
	}

	return payload, nil
}
