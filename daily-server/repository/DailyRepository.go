package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/model"
	openai "github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	dailies         *mongo.Collection
	users           *UserRepository
	userPreferences *mongo.Collection
}

func NewDailyRepository(_userRepository *UserRepository) *DailyRepository {
	return &DailyRepository{
		dailies:         database.Dailies,
		users:           _userRepository,
		userPreferences: database.UserPreferences,
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

func (r *DailyRepository) GetExplore() ([]model.Daily, error) {
	var dailies []model.Daily
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{"$sample", bson.D{{"size", 5}}}},
	}

	cursor, err := r.dailies.Aggregate(ctx, pipeline)

	/*	filter := bson.M{
			"isShared": true,
			"image":    bson.M{"$exists": true},
		}
	*/
	// opts := options.Find().SetSort(bson.M{"favourites": -1}).SetLimit(5)

	// cursor, err := r.dailies.Find(ctx, filter, opts)
	if err != nil {
		return dailies, err
	}

	err = cursor.All(ctx, &dailies)
	return dailies, err
}

func (r *DailyRepository) GetSimilarDailiesUnviewed(userId primitive.ObjectID) ([]primitive.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"author": userId}
	var userPref model.UserPreference

	err := r.userPreferences.FindOne(context.Background(), filter).Decode(&userPref)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			newUser := bson.M{"author": userId, "keywords": []string{}, "topics": []string{}}
			r.userPreferences.InsertOne(context.Background(), newUser)
			return []primitive.M{}, errors.New("no user preference found")
		} else {
			return []primitive.M{}, err
		}
	}
	interests := ""
	for _, topics := range userPref.Topic {
		interests += topics
		interests += ", "
	}
	for _, keywords := range userPref.Keywords {
		interests += keywords
		interests += ", "
	}

	// EMBEDDINGS
	client := openai.NewClient(os.Getenv("OPEN_API_KEY"))

	queryReq := openai.EmbeddingRequest{
		Input: interests,
		Model: openai.LargeEmbedding3,
	}
	targetResponse, err := client.CreateEmbeddings(ctx, queryReq)
	if err != nil {
		fmt.Println(err.Error())
		return []primitive.M{}, errors.New("preferences could not be embedded")
	}
	embedding := targetResponse.Data[0].Embedding

	vs_aggregation := bson.A{
		bson.D{
			{"$vectorSearch",
				bson.D{
					{"queryVector", embedding},
					{"path", "embedding"},
					{"numCandidates", 10},
					{"index", "embeddings_index"},
					{"limit", 3},
				},
			},
		},
		bson.D{
			{"$match", bson.D{
				{"viewers", bson.D{
					{"$ne", userId}, // Replace userId with the actual user's ObjectId
				}},
			}},
		},
		bson.D{
			{"$project", bson.D{
				{"_id", 1},
				{"text", 1}, // Replace with actual fields you want to project
				{"image", 1},
				{"emotions", 1},
				{"topic", 1},
			}},
		},
	}

	cursor, err := r.dailies.Aggregate(ctx, vs_aggregation)
	if err != nil {
		return nil, err
	}
	// Iterate over the results
	var results []bson.M
	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	matchNotViewedStage := bson.D{
		{"$match", bson.D{
			{"viewers", bson.D{
				{"$ne", userId},
			}},
		}},
	}
	sortByFavoritesStage := bson.D{
		{"$sort", bson.D{
			{"favourites", -1},
		}},
	}
	limitStage := bson.D{
		{"$limit", 2},
	}
	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 1},
			{"text", 1}, // Replace with actual fields you want to project
			{"image", 1},
			{"emotions", 1},
			{"topic", 1},
		}},
	}

	randomDailiesPipeline := mongo.Pipeline{
		matchNotViewedStage,
		sortByFavoritesStage,
		limitStage,
		projectStage,
	}
	cursor, err = r.dailies.Aggregate(ctx, randomDailiesPipeline)
	if err != nil {
		return nil, err
	}

	var randomDailiesResults []bson.M
	if err := cursor.All(ctx, &randomDailiesResults); err != nil {
		return nil, err
	}
	combinedResults := append(results, randomDailiesResults...)

	return combinedResults, nil
}

func (r *DailyRepository) GetSimilarDailies(userId primitive.ObjectID) ([]primitive.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"author": userId}
	var userPref model.UserPreference

	err := r.userPreferences.FindOne(context.Background(), filter).Decode(&userPref)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []primitive.M{}, errors.New("no user preference found")
		} else {
			return []primitive.M{}, err
		}
	}
	var interests string
	for _, topics := range userPref.Topic {
		interests += topics
		interests += ", "
	}
	for _, keywords := range userPref.Keywords {
		interests += keywords
		interests += ", "
	}

	// EMBEDDINGS
	client := openai.NewClient(os.Getenv("OPEN_API_KEY"))

	queryReq := openai.EmbeddingRequest{
		Input: interests,
		Model: openai.LargeEmbedding3,
	}
	targetResponse, err := client.CreateEmbeddings(ctx, queryReq)
	if err != nil {
		return []primitive.M{}, errors.New("preferences could not be embedded")
	}
	embedding := targetResponse.Data[0].Embedding

	aggregation := bson.A{
		bson.D{
			{"$vectorSearch",
				bson.D{
					{"queryVector", embedding},
					{"path", "embedding"},
					{"numCandidates", 10},
					{"index", "embeddings_index"},
					{"limit", 10},
				},
			},
		},
		bson.D{
			{"$project", bson.D{
				{"_id", 1},
				{"text", 1}, // Replace with actual fields you want to project
				{"score", bson.D{{"$meta", "searchScore"}}},
			}},
		},
	}

	cursor, err := r.dailies.Aggregate(ctx, aggregation)
	if err != nil {
		return nil, err
	}
	// Iterate over the results
	var results []bson.M
	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *DailyRepository) List(author_id primitive.ObjectID, limit int) ([]model.Daily, error) {
	var dailies []model.Daily
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetLimit(int64(limit)).SetSort(bson.D{{"createdAt", -1}})

	cursor, err := r.dailies.Find(ctx, bson.M{"author": author_id}, opts)
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
	opts := options.Update().SetUpsert(true)

	if _, err := r.dailies.UpdateOne(ctx, bson.M{"_id": dailyID}, updateFavourites, opts); err != nil {
		return err
	}

	err := r.users.AddToFav(userID, dailyID) // Assuming AddToFav is implemented correctly
	if err != nil {
		return err
	}
	return nil
}

func (r *DailyRepository) UpdateUserPreferences(keywords []string, topic string, authorId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"author": authorId}
	update := bson.M{
		"$addToSet": bson.M{
			"keywords": bson.M{"$each": keywords},
			"topics":   topic,
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err := r.userPreferences.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
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
	result, err := r.dailies.UpdateOne(ctx, getDaily, dailyOperation)
	if result.MatchedCount == 0 && err != nil {
		return errors.New("not found")
	}
	if result.ModifiedCount == 0 && err != nil {
		return errors.New("same image")
	}
	return nil
}
