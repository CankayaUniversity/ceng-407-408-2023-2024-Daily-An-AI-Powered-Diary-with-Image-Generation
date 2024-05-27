package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserPreference struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Author   primitive.ObjectID `json:"author" bson:"author" binding:"required"`
	Topics   []string           `json:"topics" bson:"topics"`
	Keywords []string           `json:"keywords" bson:"keywords"`
}
