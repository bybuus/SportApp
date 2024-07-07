package repository

import (
	"context"
	"github.com/kendoow/SportApp/backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repo) FindUserByPhone(ctx context.Context, phone string) (*entity.UserEntity, error) {
	authCollection := repo.collections.Auth
	filter := bson.M{"phone": phone}
	var result entity.UserEntity

	// Find the user with the given phone number
	err := authCollection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// If the error is ErrNoDocuments, it means no user was found
			return nil, nil
		}
		// Handle other potential errors TODO: add logger
		return nil, err
	}

	return &result, nil //TODO returns UserEntity because use ObjectId in update
}

func (repo *Repo) UpsertUserByPhone(
	phone string, login string, token string, clientId string,
) error {
	authCollection := repo.collections.Auth

	filter := bson.M{"phone": phone}
	updates := bson.M{
		"$set": bson.D{
			{"login", login},
			{"phone", phone},
		},
		"$addToSet": bson.M{
			"tokenByClient": bson.M{
				"token":    token,
				"clientId": clientId,
			},
		},
	}
	_, err := authCollection.UpdateOne(
		context.TODO(),
		filter,
		updates,
		options.Update().SetUpsert(true))

	return err
}
