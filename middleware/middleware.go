package middleware

import (
	"log"
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

func JwtAuthMiddlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := utils.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"message": "Unauthorized"})
			c.Abort()
			log.Fatal(err)
			return
		}
		objectId, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"message": "Unauthorized"})
			c.Abort()
			log.Fatal(err)
			return
		}
		c.Set("user_id", objectId)
		var result model.User
		err = database.Users.FindOne(c, bson.M{"_id": c.Keys["user_id"]}).Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
			c.Abort()
			log.Fatal(err)
			return
		}
		if strings.TrimSpace(strings.ToLower(result.Role)) != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized, user is not an admin"})
			c.Abort()
			log.Fatal(err)
			return
		}
	}
}
