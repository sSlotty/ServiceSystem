package middleware

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"sSlotty/authentication-service/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "data": bson.M{}})
			c.Abort()
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.JWTAuthService().ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "data": bson.M{}})
			c.Abort()
			return
		}

		if token.Valid {
			_, err := token.Claims.(jwt.MapClaims)
			if err {
				c.Abort()
				c.AbortWithStatus(http.StatusUnauthorized)

			}
		} else {
			fmt.Println("testing")
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
