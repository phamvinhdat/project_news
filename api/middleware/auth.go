package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/phamvinhdat/project_news/api"
)

type JWTAuthen struct{}

type Token struct {
	jwt.StandardClaims
	Username string
}

func New()*JWTAuthen{
	return &JWTAuthen{}
}

//JWTAuthentication check whether token client submissions are valid
func (*JWTAuthen)JWTAuthentication() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		response := make(map[string]interface{})
		tokenHeader := c.Request.Header.Get("Authorization") //grap the token from the header

		if tokenHeader == "" { //token missing
			response = api.Message(false, "Missing auth token")
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		seperate := " "
		splitted := strings.Split(tokenHeader, seperate)
		if len(splitted) != 2 {
			response = api.Message(false, "Invalid/malfomed auth token")
			c.JSON(http.StatusForbidden, response)
			return
		}

		tokenPath := splitted[1]
		tk := &Token{}
		godotenv.Load()
		token, err := jwt.ParseWithClaims(tokenPath, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = api.Message(false, "Malformed authentication token")
			c.JSON(http.StatusForbidden, response)
			return
		}

		if !token.Valid {
			response = api.Message(false, "Token is not valid.")
			c.JSON(http.StatusForbidden, response)
			return
		}

		ctx := context.WithValue(c.Request.Context(), "username", tk.Username)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	})
}

func (*JWTAuthen)NewToken(username string, minute int) string {
	tk := &Token{
		Username: username,
	}
	tk.ExpiresAt = time.Now().Add(time.Minute * time.Duration(minute)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	godotenv.Load()

	tokenStr, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil{ 
		return ""
	}

	return tokenStr
}
