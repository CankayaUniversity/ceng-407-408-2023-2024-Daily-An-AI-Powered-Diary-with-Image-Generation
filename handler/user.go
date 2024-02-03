package handler

import (
	"net/http"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteUser(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}

	res, err := database.Users.DeleteOne(c, bson.M{"_id": userId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}

	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no user was deleted"})
	}
	c.JSON(http.StatusOK, res)
}
