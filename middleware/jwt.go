package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"time"
	"errors"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		if token == "" {
			token = c.Request.Header.Get("Authorization")
			if s := strings.Split(token, " "); len(s) == 2 {
				token = s[1]
			}
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				if token, err = j.RefreshToken(token); err == nil {
					c.Header("Authorization", "Bear "+token)
					c.JSON(200, gin.H{"code": 200, "msg": "token已刷新", "data": map[string]string{"token": token}})
					c.Abort()
				}
			}
			c.JSON(401, gin.H{"code": 401, "msg": "token认证失败", "data": struct{}{}})
			c.Abort()
		}
		c.Set("claims", claims)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

type CustomClaims struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	HeadImg  string `json:"headImg"`
	HashIP   string `json:"hashIP"`
	RoleId     int `json:"roleId"`
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{[]byte("zhoujianjun666"),}
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
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
