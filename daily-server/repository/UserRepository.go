package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	users *mongo.Collection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: database.Users,
	}
}

func (r *UserRepository) AddToFav(userId primitive.ObjectID, dailyId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userRecord model.User
	err := r.users.FindOne(ctx, bson.M{"_id": userId, "favouriteDailies": bson.M{"$exists": true}}).Decode(&userRecord)
	if err == mongo.ErrNoDocuments {
		_, err = r.users.UpdateOne(ctx, bson.M{"_id": userId}, bson.M{"$set": bson.M{"favouriteDailies": []primitive.ObjectID{}}})
		if err != nil {
			return err
		}
	}

	update := bson.M{"$addToSet": bson.M{"favouriteDailies": dailyId}}
	_, err = r.users.UpdateOne(ctx, bson.M{"_id": userId}, update)
	return err
}

func (r *UserRepository) AddToViewed(userId primitive.ObjectID, dailyId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userRecord model.User
	err := r.users.FindOne(ctx, bson.M{"_id": userId, "viewedDailies": bson.M{"$exists": true}}).Decode(&userRecord)
	if err == mongo.ErrNoDocuments {
		fmt.Println("Buraya girdim")
		// If viewedDailies doesn't exist, initialize it
		_, err = r.users.UpdateOne(ctx, bson.M{"_id": userId}, bson.M{"$set": bson.M{"viewedDailies": []primitive.ObjectID{}}})
		if err != nil {
			return err
		}
	}

	update := bson.M{"$addToSet": bson.M{"viewedDailies": dailyId}}
	_, err = r.users.UpdateOne(ctx, bson.M{"_id": userId}, update)
	return err
}

func (r *UserRepository) Create(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.users.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) FindById(id primitive.ObjectID) (model.User, error) {
	var user model.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.users.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}

func (r *UserRepository) List() ([]model.User, error) {
	var users []model.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.users.Find(ctx, bson.D{})
	if err != nil {
		return users, err
	}

	err = cursor.All(ctx, &users)
	return users, err
}

func (r *UserRepository) Delete(_id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.users.DeleteOne(ctx, bson.M{"_id": _id})
	return err
}

func (r *UserRepository) Replace(id primitive.ObjectID, newUser *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.users.ReplaceOne(ctx, bson.M{"_id": id}, newUser)
	return err
}
