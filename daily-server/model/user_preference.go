package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserPreference struct {
	ID       primitive.ObjectID   `json:"id" bson:"_id"`
	AuthorId primitive.ObjectID   `json:"author_id" bson:"author_id" binding:"required"`
	Topic    []string             `json:"topic" bson:"topic"`
	Keywords []string             `json:"keywords" bson:"keywords"`
	Liked    []primitive.ObjectID `json:"liked", bson:"liked`
}
