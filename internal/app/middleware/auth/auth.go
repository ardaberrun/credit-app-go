package middleware

import (
	"github/ardaberrun/credit-app-go/internal/app/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
		tokenString := utils.GetToken(c.GetHeader("Authorization"));
		if tokenString != "" {
			claims, err := utils.ValidateJWT(tokenString);
			if err == nil {
				c.Set("claims", claims);
			}
		}

		c.Next();
	}
}