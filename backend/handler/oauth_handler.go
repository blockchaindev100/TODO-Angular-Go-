package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/blockchaindev100/todo/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/idtoken"
)

type TokenRequest struct {
	Token string `json:"token"`
}

func GoogleLogin(c *gin.Context) {
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")
	c.Redirect(http.StatusSeeOther, url)
}

// func GoogleCallback(c *gin.Context) {
// 	state := c.Query("state")
// 	if state != "randomstate" {
// 		c.String(http.StatusUnauthorized, "Invalid state parameter")
// 		return
// 	}

// 	code := c.Query("code")
// 	fmt.Println("Code", code)
// 	token, err := config.AppConfig.GoogleLoginConfig.Exchange(context.Background(), code)
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Failed to exchange token")
// 		return
// 	}
// 	fmt.Println("Access token", token.AccessToken)

// 	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Failed to get user info")
// 		return
// 	}
// 	defer resp.Body.Close()

// 	userData, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Failed to read response body")
// 		return
// 	}
// 	fmt.Println(string(userData))
// 	c.IndentedJSON(http.StatusOK, token)
// }

func GoogleLoginHandler(c *gin.Context) {
	var tokenRequest TokenRequest
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}
	if err := c.BindJSON(&tokenRequest); err != nil {
		c.String(http.StatusBadRequest, "Invalid request payload")
		return
	}

	payload, err := idtoken.Validate(context.Background(), tokenRequest.Token, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		c.String(http.StatusUnauthorized, "Invalid token")
		return
	}

	// Directly use the payload.Claims map
	userInfo := payload.Claims

	fmt.Println("User Info", userInfo)
	c.JSON(http.StatusOK, userInfo)
}
