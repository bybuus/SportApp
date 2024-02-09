package workout

import (
	"context"
	"github.com/kendoow/SportApp/backend/config"
	"github.com/kendoow/SportApp/backend/internal/model"
	"github.com/kendoow/SportApp/backend/internal/utils/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	workoutCollection = mongoDB.Database(config.GetAppConfig().Mongo.DBName).Collection(config.GetAppConfig().Mongo.Workout.CollectionName)
)

func GetAllWorkout(ctx context.Context, filter *bson.M) (*mongo.Cursor, error) {
	cursor, err := workoutCollection.Find(ctx, filter)
	if err != nil {
		logging.Error.Println(err.Error())
		return nil, err
	}

	return cursor, nil
}

//TODO refactor to dao

func CreateWorkout(ctx context.Context, workout *model.Workout) (*mongo.InsertOneResult, error) {
	result, err := workoutCollection.InsertOne(ctx, workout)
	if err != nil {
		logging.Error.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func GetWorkoutById(ctx context.Context, filter bson.D) *mongo.SingleResult {
	result := workoutCollection.FindOne(ctx, filter)
	return result
}

func DeleteWorkouts(ctx context.Context, filter *bson.M) (*model.BulkWorkoutIds, error) {
	result, err := workoutCollection.DeleteMany(ctx, filter)
	if err != nil {
		logging.Err(err, "Failed deletion of workouts in repo, with reason: ")
		return nil, err
	}

	return result.DeletedCount
}
