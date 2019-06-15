package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/phamvinhdat/project_news/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_"github.com/phamvinhdat/project_news/api"
	"github.com/phamvinhdat/project_news/models"
	"golang.org/x/crypto/bcrypt"
)

type JWTAuthen struct {
	UserRepo repository.IUserRepo
}

type Token struct {
	jwt.StandardClaims
	Username string
}

func NewJWTAuthen(user repository.IUserRepo) *JWTAuthen {
	return &JWTAuthen{
		UserRepo: user,
	}
}

func (*JWTAuthen) ParseToken(tokenStr string) (*Token, error) {
	tk := &Token{}
	godotenv.Load()
	_, err := jwt.ParseWithClaims(tokenStr, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("token_password")), nil
	})
	if err != nil {
		return nil, err
	}

	return tk, nil
}

//JWTAuthentication check whether token client submissions are valid
func (*JWTAuthen) JWTAuthentication() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		//response := make(map[string]interface{})
		cookie, err := c.Request.Cookie("token")
		if err != nil { //token missing
			//response = api.Message(false, "Missing auth token")
			//c.JSON(http.StatusUnauthorized, response)
			c.Redirect(http.StatusSeeOther, "/")
			return
		}

		tokenPath := cookie.Value
		tk := &Token{}
		godotenv.Load()
		token, err := jwt.ParseWithClaims(tokenPath, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			//response = api.Message(false, "Malformed authentication token")
			//c.JSON(http.StatusForbidden, response)
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}

		if !token.Valid {
			//response = api.Message(false, "Token is not valid.")
			//c.JSON(http.StatusForbidden, response)
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), "username", tk.Username)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	})
}

func (*JWTAuthen) NewToken(username string, minute int) string {
	tk := &Token{
		Username: username,
	}
	tk.ExpiresAt = time.Now().Add(time.Minute * time.Duration(minute)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	godotenv.Load()

	tokenStr, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return ""
	}

	return tokenStr
}

func (j *JWTAuthen) Validate(user *models.User) error {
	if !strings.Contains(user.Email, "@") {
		return errors.New("Email is invalid")
	}

	if len(user.Password) < 6 {
		return errors.New("Password length have to greater than 6")
	}

	if len(user.Name) < 1 {
		return errors.New("Name is required")
	}

	userReturn, _ := j.UserRepo.FetchByEmail(user.Email)
	if userReturn != nil {
		errors.New("Email address already in use by another user")
	}

	userReturn, _ = j.UserRepo.FetchByUsername(user.Username)
	if userReturn != nil {
		errors.New("Username exists")
	}

	return nil
}

func (j *JWTAuthen) Auth(username, password string) error {
	findUser, err := j.UserRepo.FetchByUsername(username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
