package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserPreference struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Author   primitive.ObjectID `json:"author" bson:"author" binding:"required"`
	Topic    []string           `json:"topic" bson:"topic"`
	Keywords []string           `json:"keywords" bson:"keywords"`
}
