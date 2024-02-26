package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReportedDaily struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	DailyID    primitive.ObjectID `json:"dailyId" bson:"dailyId" binding:"required"`
	ReportedAt primitive.DateTime `json:"reportedAt" bson:"reportedAt"`
	Reports    int                `json:"reports" bson:"reports"`
	Title      string             `json:"title" bson:"title" binding:"required"`
	Content    string             `json:"content" bson:"content"`
}
