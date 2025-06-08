package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/bapakfadil/fastcampus/internal/configs"
	"github.com/bapakfadil/fastcampus/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		
		// Response if user hit API without token
		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		// JWT Token checker
		userID, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}

func AuthRefreshMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		
		// Response if user hit API without token
		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		// JWT Token checker
		userID, username, err := jwt.ValidateTokenWithoutExpiry(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}