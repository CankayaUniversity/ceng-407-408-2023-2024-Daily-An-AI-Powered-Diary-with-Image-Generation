package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/model"
	"github.com/Final-Projectors/daily-server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized token validation"})
			c.Abort()
			return
		}
		userId, err := utils.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"message": "Unauthorizedfrom middleware"})
			c.Abort()
			return
		}
		objectId, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"message": "Unauthorized from middleware"})
			c.Abort()
			return
		}
		c.Set("user_id", objectId)
		c.Next()
	}
}

func JwtAuthMiddlewareRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		var result model.User
		err = database.Users.FindOne(c, bson.M{"_id": c.Keys["user_id"]}).Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		if strings.TrimSpace(strings.ToLower(result.Role)) != role {
			if strings.TrimSpace(strings.ToLower(result.Role)) != "admin" {
				c.JSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("Unauthorized, user is not %v", role)})
				c.Abort()
				return
			}
		}
	}
}
