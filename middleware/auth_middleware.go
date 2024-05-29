package middleware

import (
	"gin_social/errorhandler"
	"gin_social/helpers"

	"github.com/gin-gonic/gin"
)


func JWTMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			errorhandler.HandleError(c, &errorhandler.Unauthorized{Message: "UnauthorizedEmpty"})
			c.Abort()
			return
		}

		userId, err := helpers.ValidateToken(tokenString)

		if err != nil{
			errorhandler.HandleError(c, &errorhandler.Unauthorized{Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("UserID", *userId)
		c.Next()
	}
}