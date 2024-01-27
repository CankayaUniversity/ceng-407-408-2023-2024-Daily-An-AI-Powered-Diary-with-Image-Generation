package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Daily struct {
	ID         primitive.ObjectID   `json:"id" bson:"_id"`
	Text       string               `json:"text" bson:"text"`
	Author     primitive.ObjectID   `json:"author" bson:"author"`
	Keywords   []string             `json:"keywords" bson:"keywords"`
	Emotions   Emotion              `json:"emotions" bson:"emotions"`
	Image      string               `json:"image" bson:"image"`
	Favourites int                  `json:"favourites" bson:"favourites"`
	CreatedAt  primitive.DateTime   `json:"createdAt" bson:"createdAt"`
	Viewers    []primitive.ObjectID `json:"viewers" bson:"viewers"`
	IsShared   bool                 `json:"isShared" bson:"isShared"`
}

type Emotion struct {
	anger int
	happy int
	sad   int
	shock int
}
