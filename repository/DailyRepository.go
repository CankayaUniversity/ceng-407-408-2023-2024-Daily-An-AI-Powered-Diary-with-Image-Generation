package repository

import (
	"context"
	"time"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
	HELPERS:
	D: An ordered representation of a BSON document (slice)
	M: An unordered representation of a BSON document (map)
	A: An ordered representation of a BSON array
	E: A single element inside a D type

	collection.FindOne(context.TODO(), filter).Decode(&result), collection.InsertOne(context.TODO(), object)
	collection.DeleteOne(ctx, filter), collection.ReplaceOne(ctx, filter, replacement)
	collection.UpdateOne(ctx, filter, update)

	Instead of UpdateOne, use ReplaceOne to change the fields of a resource, update is to change the structure

	ctx.WithTimeout allows us to cancel an operation after a specified time
*/

type DailyRepository struct {
	dailies *mongo.Collection
}

func NewDailyRepository(_dailies *mongo.Collection) *DailyRepository {
	return &DailyRepository{
		dailies: _dailies,
	}
}

func (r *DailyRepository) Create(daily *model.Daily) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.dailies.InsertOne(ctx, daily)
	return err
}

func (r *DailyRepository) FindById(id primitive.ObjectID) (model.Daily, error) {
	var daily model.Daily
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.dailies.FindOne(ctx, bson.M{"_id": id}).Decode(&daily)
	return daily, err
}

func (r *DailyRepository) List(author_id primitive.ObjectID) ([]model.Daily, error) {
	var dailies []model.Daily
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.Dailies.Find(ctx, bson.M{"author": author_id})
	if err != nil {
		return dailies, err
	}

	err = cursor.All(ctx, dailies)
	return dailies, err
}

func (r *DailyRepository) Delete(_id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.dailies.DeleteOne(ctx, bson.M{"_id": _id})
	return err
}

func (r *DailyRepository) Replace(id primitive.ObjectID, newDaily *model.Daily) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.dailies.ReplaceOne(ctx, bson.M{"_id": id}, newDaily)
	return err
}

func (r *DailyRepository) UpdateImage(id primitive.ObjectID, newImage string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.dailies.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"image": newImage}})
	return err
}

func (r *DailyRepository) AddViewer(dailyID primitive.ObjectID, viewerID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.dailies.UpdateOne(ctx, bson.M{"_id": dailyID}, bson.M{"$inc": bson.M{"favourites": 1}})
	if err != nil {
		return err
	}

	update := bson.D{{Key: "$addToSet", Value: bson.D{{Key: "viewers", Value: viewerID}}}}
	_, err = r.dailies.UpdateOne(ctx, bson.M{"_id": dailyID}, update)

	return err
}

func (r *DailyRepository) UpdateKeywords(dailyID primitive.ObjectID, keywords []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "keywords", Value: keywords}}}}
	_, err := r.dailies.UpdateOne(ctx, bson.M{"_id": dailyID}, update)

	return err
}
