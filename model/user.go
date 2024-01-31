package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID               primitive.ObjectID   `json:"id" bson:"_id"`
	Email            string               `json:"email" bson:"email" binding:"required"`
	Password         string               `json:"password" bson:"password" binding:"required"`
	FavouriteDailies []primitive.ObjectID `json:"favouriteDailies" bson:"favouriteDailies"`
	CreatedAt        primitive.DateTime   `json:"createdAt" bson:"createdAt"`
	ViewedDailies    []primitive.ObjectID `json:"viewedDailies" bson:"viewedDailies"`
	Role             string               `json:"role" bson:"role"`
}

type UserLoginRegisterRequest struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}
