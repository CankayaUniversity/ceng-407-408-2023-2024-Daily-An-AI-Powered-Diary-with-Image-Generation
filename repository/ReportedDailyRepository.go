package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Final-Projectors/daily-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReportedDailyRepository struct {
	reportedDailies *mongo.Collection
}

func NewReportedDailyRepository(_reportedDailies *mongo.Collection) *ReportedDailyRepository {
	return &ReportedDailyRepository{
		reportedDailies: _reportedDailies,
	}
}

func (r *ReportedDailyRepository) Create(reportedDaily *model.ReportedDaily) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.reportedDailies.InsertOne(ctx, reportedDaily)
	if err != nil {
		return errors.New("create_error")
	}
	return nil
}

func (r *ReportedDailyRepository) FindById(id primitive.ObjectID) (model.ReportedDaily, error) {
	var reportedDaily model.ReportedDaily
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.reportedDailies.FindOne(ctx, bson.M{"_id": id}).Decode(&reportedDaily)
	if err != nil {
		return reportedDaily, err
	}
	return reportedDaily, nil
}

func (r *ReportedDailyRepository) List(dailyId primitive.ObjectID) ([]model.ReportedDaily, error) {
	var reportedDailies []model.ReportedDaily
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.reportedDailies.Find(ctx, bson.M{"dailyId": dailyId})
	if err != nil {
		return reportedDailies, err
	}

	err = cursor.All(ctx, &reportedDailies)
	return reportedDailies, err
}

func (r *ReportedDailyRepository) Delete(_id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.reportedDailies.DeleteOne(ctx, bson.M{"_id": _id})
	return err
}

func (r *ReportedDailyRepository) Replace(id primitive.ObjectID, newReportedDaily *model.ReportedDaily) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.reportedDailies.ReplaceOne(ctx, bson.M{"_id": id}, newReportedDaily)
	return err
}

func (r *ReportedDailyRepository) IncrementReports(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.reportedDailies.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$inc": bson.M{"reports": 1}})
	return err
}
