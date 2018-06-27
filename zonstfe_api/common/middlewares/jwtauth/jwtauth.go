package jwtauth

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"errors"
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

func JWTAuth(SigningKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		token := c.Request.Header.Get("Authorization")
		if s := strings.Split(token, " "); len(s) == 2 {
			token = s[1]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
		// 解析token
		jwt := New(SigningKey)
		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
		c.Set("current_user", claims)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}
type CustomClaims struct {
	Id       int    `json:"id"`
	AppKey   string `json:"app_key"`
	Email    string `json:"email"`
	DealType string `json:"deal_type"`
	jwt.StandardClaims
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

//解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

//创建token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//刷新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

func New(SigningKey string) *JWT {
	return &JWT{
		[]byte(SigningKey),
	}
}
