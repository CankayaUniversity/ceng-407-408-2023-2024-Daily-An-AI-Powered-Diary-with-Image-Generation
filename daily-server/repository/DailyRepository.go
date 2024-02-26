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
	users   *UserRepository
}

func NewDailyRepository(_userRepository *UserRepository) *DailyRepository {
	return &DailyRepository{
		dailies: database.Dailies,
		users:   _userRepository,
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

	cursor, err := r.dailies.Find(ctx, bson.M{"author": author_id})
	if err != nil {
		return dailies, err
	}

	err = cursor.All(ctx, &dailies)
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

func (r *DailyRepository) FavouriteDaily(dailyID primitive.ObjectID, userID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateFavourites := bson.M{"$inc": bson.M{"favourites": 1}}
	if _, err := r.dailies.UpdateOne(ctx, bson.M{"_id": dailyID}, updateFavourites); err != nil {
		return err
	}

	err := r.users.AddToFav(userID, dailyID) // Assuming AddToFav is implemented correctly
	return err
}

func (r *DailyRepository) View(dailyID primitive.ObjectID, viewerID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Add the viewer to the daily
	updateViewers := bson.M{"$addToSet": bson.M{"viewers": viewerID}}
	if _, err := r.dailies.UpdateOne(ctx, bson.M{"_id": dailyID}, updateViewers); err != nil {
		return err
	}

	// Add the daily to the user's viewed list
	err := r.users.AddToViewed(viewerID, dailyID) // Assuming AddToViewed is implemented correctly
	return err
}

func (r *DailyRepository) UpdateKeywords(dailyID primitive.ObjectID, keywords []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "keywords", Value: keywords}}}}
	_, err := r.dailies.UpdateOne(ctx, bson.M{"_id": dailyID}, update)

	return err
}

func (r *DailyRepository) EditDailyImage(dailyID primitive.ObjectID, image string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	getDaily := bson.M{"_id": dailyID}
	dailyOperation := bson.M{"$set": bson.M{"image": image}}
	_, err := r.dailies.UpdateOne(ctx, getDaily, dailyOperation)
	return err
}
