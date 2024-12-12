package middleware

import (
	"Master_Data/package/response"
	"Master_Data/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MasterDataMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.SendResponse(c, http.StatusUnauthorized, "Missing token", nil, "Missing token")
			return
		}

		claims, err := utils.ValidateJWT(token)
		if err != nil {
			response.SendResponse(c, http.StatusUnauthorized, "Invalid token", nil, "Invalid token")
			return
		}

		c.Set("client_id", claims.Claims.(jwt.MapClaims)["client_id"])
		c.Set("uuid_key", claims.Claims.(jwt.MapClaims)["uuid_key"])
		c.Set("role", claims.Claims.(jwt.MapClaims)["role"])
		c.Next()
	}
}
