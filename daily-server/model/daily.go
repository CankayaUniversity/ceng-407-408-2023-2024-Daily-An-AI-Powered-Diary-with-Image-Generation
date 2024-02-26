package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Daily struct {
	ID         primitive.ObjectID   `json:"id" bson:"_id"`
	Text       string               `json:"text" bson:"text" binding:"required"`
	Author     primitive.ObjectID   `json:"author" bson:"author"`
	Keywords   []string             `json:"keywords" bson:"keywords,omitempty"`
	Emotions   Emotion              `json:"emotions" bson:"emotions"`
	Image      string               `json:"image" bson:"image"`
	Favourites int                  `json:"favourites" bson:"favourites"`
	CreatedAt  primitive.DateTime   `json:"createdAt" bson:"createdAt"`
	Viewers    []primitive.ObjectID `json:"viewers" bson:"viewers,omitempty"`
	IsShared   bool                 `json:"isShared" bson:"isShared"`
}

type Emotion struct {
	Anger     int `json:"anger" bson:"anger"`
	Happiness int `json:"happiness" bson:"happiness"`
	Sadness   int `json:"sadness" bson:"sadness"`
	Shock     int `json:"shock" bson:"shock"`
}

type DailyRequestDTO struct {
	ID primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
}

type CreateDailyDTO struct {
	Text     string `json:"text" bson:"text" binding:"required"`
	Image    string `json:"image" bson:"image"`
	IsShared *bool  `json:"isShared" bson:"isShared" binding:"required"`
}

type DeleteDailyDTO struct {
	ID *primitive.ObjectID `json:"id" bson:"_id"`
}

type EditDailyImageDTO struct {
	ID    primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Image string             `json:"image" bson:"image"`
}
