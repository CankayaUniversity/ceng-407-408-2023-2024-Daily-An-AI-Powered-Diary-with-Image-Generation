package middleware

import (
	"net/http"

	"github.com/Final-Projectors/daily-server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		userId, err := utils.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		objectId, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user_id", objectId)
		c.Next()
	}
}
