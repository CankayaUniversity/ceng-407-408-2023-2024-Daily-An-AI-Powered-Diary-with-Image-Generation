package handler

import (
	"context"
	"net/http"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteUserAdmin(c *gin.Context) {
	var userRequest model.UserDeleteRequest
	var user model.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := database.Users.FindOne(c, bson.M{"email": userRequest.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if user.Role == "admin" || user.Role == "moderator" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user is admin"})
		return
	}

	res, err := database.Users.DeleteOne(c, bson.M{"email": userRequest.Email})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}
	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no user was deleted"})
	}
	c.JSON(http.StatusOK, res)
}

func GrantModRights(c *gin.Context) {
	var userRequest model.UserMakeAdminRequest
	var user model.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := database.Users.FindOne(c, bson.M{"email": userRequest.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}
	if user.Role == "moderator" || user.Role == "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user is already moderator or admin"})
		return
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "role", Value: "moderator"}}}}
	_, er := database.Users.UpdateOne(context.TODO(), bson.M{"email": userRequest.Email}, update)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "moderator role granted"})
}

func TakeModRights(c *gin.Context) {
	var userRequest model.UserMakeAdminRequest
	var user model.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := database.Users.FindOne(c, bson.M{"email": userRequest.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}
	if user.Role != "moderator" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user is not a moderator"})
		return
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "role", Value: "user"}}}}
	_, er := database.Users.UpdateOne(context.TODO(), bson.M{"email": userRequest.Email}, update)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user role took"})
}
