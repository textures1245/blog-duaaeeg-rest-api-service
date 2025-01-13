package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		log.Infof("accessToken: %+v", accessToken)
		if accessToken == "" {
			log.Info("error, authorization header is empty.")
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":      "Unauthorized",
				"status_code": http.StatusUnauthorized,
				"message":     "unauthorized access",
				"result":      nil,
			})
			c.Abort()
			return
		}

		secretKey := os.Getenv("JWT_SECRET_TOKEN")
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("error, unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(secretKey), nil
		})
		if err != nil {
			log.Info(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":      http.StatusText(http.StatusUnauthorized),
				"status_code": http.StatusUnauthorized,
				"message":     "error, unauthorized",
				"result":      nil,
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Infof("claims %+v", claims)
			c.Keys = make(map[string]interface{})
			c.Keys["user_uuid"] = claims["user_uuid"]
			c.Keys["email"] = claims["email"]
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":      http.StatusText(http.StatusUnauthorized),
				"status_code": http.StatusUnauthorized,
				"message":     "error, unauthorized",
				"result":      nil,
			})
			c.Abort()
			return
		}
	}
}
