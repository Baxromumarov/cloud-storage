package middleware

import (
	"net/http"
	"strings"

	"github.com/baxromumarov/cloud-storage/config"
	"github.com/baxromumarov/cloud-storage/internal/models"
	"github.com/baxromumarov/cloud-storage/internal/pkg/jwt"


	"github.com/gin-gonic/gin"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		} else {
			splitedToken := strings.Split(token, " ")
			if len(splitedToken) != 2 {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid/Malformed auth token"})
				c.Abort()
				return
			}
			
			token = splitedToken[1]

			tokenInfo, err := jwt.JWTExtract(token, cfg.JWTSigningKey)
			if err != nil {
				c.JSON(401, models.Response{
					Code:    http.StatusUnauthorized,
					Message: err.Error(),
				})
				c.Abort()
				return
			}

			for key, value := range tokenInfo {
				c.Request.Header.Set(key, value)
				key1 := key
				value1 := value
				c.Set(key1, value1)
			}
		}
		c.Next()
	}
}
