package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Daily struct {
	ID         primitive.ObjectID   `json:"id" bson:"_id"`
	Text       string               `json:"text" bson:"text" binding:"required"`
	Author     primitive.ObjectID   `json:"author" bson:"author"`
	Topics     []string             `json"topics" bson:"topics"`
	Keywords   []string             `json:"keywords" bson:"keywords,omitempty"`
	Emotions   Emotion              `json:"emotions" bson:"emotions"`
	Image      string               `json:"image" bson:"image"`
	Favourites int                  `json:"favourites" bson:"favourites"`
	CreatedAt  primitive.DateTime   `json:"createdAt" bson:"createdAt"`
	Viewers    []primitive.ObjectID `json:"viewers" bson:"viewers,omitempty"`
	IsShared   bool                 `json:"isShared" bson:"isShared"`
	Embedding  []float32            `json:"embedding" bson:"embedding"`
}

type Emotion struct {
	Surprise float64 `json:"surprise"`
	Love     float64 `json:"love"`
	Anger    float64 `json:"anger"`
	Joy      float64 `json:"joy"`
	Sadness  float64 `json:"sadness"`
	Fear     float64 `json:"fear"`
}

type Prediction struct {
	Topics   []string `json:"topics"`
	Emotions Emotion  `json:"emotions"`
	Keywords []string `json:"keywords"`
}

type PredictionResponse struct {
	Predictions       []Prediction `json:"predictions"`
	DeployedModelID   string       `json:"deployed_model_id"`
	ModelVersionID    string       `json:"model_version_id"`
	ModelResourceName string       `json:"model_resource_name"`
}

type DailyRequestDTO struct {
	ID primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
}

type Prompt2ImgDTO struct {
	Text string `json:"text" bson:"text" binding:"required"`
}

type CreateDailyDTO struct {
	Text     string `json:"text" bson:"text" binding:"required"`
	Image    string `json:"image" bson:"image"`
	IsShared *bool  `json:"isShared" bson:"isShared" binding:"required"`
	Prompt   string `json:"prompt" bson:"prompt"`
}

type DeleteDailyDTO struct {
	ID *primitive.ObjectID `json:"id" bson:"_id"`
}

type EditDailyImageDTO struct {
	ID    primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Image string             `json:"image" bson:"image"`
}

type ExploreDailyDTO struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Text     string             `json:"text" bson:"text" binding:"required"`
	Topics   []string           `json"topics" bson:"topics"`
	Emotions Emotion            `json:"emotions" bson:"emotions"`
	Image    string             `json:"image" bson:"image"`
}
