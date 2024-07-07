package repository

import (
	"context"
	"fmt"
	"github.com/kendoow/SportApp/backend/config"
	"github.com/kendoow/SportApp/backend/db"
	"github.com/kendoow/SportApp/backend/internal/entity"
	"github.com/kendoow/SportApp/backend/test/common/integration/base"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

type TestAuthRepo struct {
	base.MongoTestBase
	Repo           *Repo
	CollectionName string
}

func (suite *TestAuthRepo) SetupSuite() {
	suite.BeforeMongo()

	configuration := config.GetAppConfig()
	database, err := db.InitDB(configuration)
	suite.Require().NoError(err, "Unable to create db")
	suite.Repo = CreateRepo(database)
	suite.CollectionName = config.GetAppConfig().Mongo.Auth.CollectionName
}

func (suite *TestAuthRepo) TearDownSuite() {
	suite.AfterMongo()
}

func TestAuthRepo_Run(t *testing.T) {
	suite.Run(t, new(TestAuthRepo))
}

func authValue() bson.D {
	return bson.D{
		{"phone", "+79991112233"},
		{"login", "sportman"},
		{"tokenByClient", bson.A{
			bson.M{"token": "token", "clientId": "clientId"},
		}},
	}
}

func (suite *TestAuthRepo) TestRepo_FindUserByPhone() {
	//given
	value := authValue()
	suite.CreateSomething(
		context.Background(),
		suite.CollectionName,
		&value)

	//when
	dbEntity, err := suite.Repo.FindUserByPhone(context.Background(), fmt.Sprint(value[0].Value))
	suite.Require().NoError(err, "Err then calling method FindUserByPhone")

	//then
	suite.Require().Equal(value[0].Value, dbEntity.Phone)
	suite.Require().Equal(value[1].Value, dbEntity.Login)

	//over
	suite.CleanCollection(context.Background(), suite.CollectionName)
}

func (suite *TestAuthRepo) TestRepo_UpsertUserByPhone_NewUser() {
	//given
	login := "sportman"
	phone := "+79991112233"
	token := "token"
	clientId := "client"

	//when
	err := suite.Repo.UpsertUserByPhone(
		phone,
		login,
		token,
		clientId)
	suite.Require().NoError(err, "Err then calling method InsertUser")

	result := suite.FindAll(context.Background(), suite.CollectionName, func(cursor *mongo.Cursor) interface{} {
		dbEntity := entity.UserEntity{}
		err := cursor.Decode(&dbEntity)
		suite.Require().NoError(err, "Can not decode entity")

		return dbEntity
	})

	suite.Require().NotEmptyf(result, "Collection should not be empty")

	actualEntity := result[0].(entity.UserEntity)
	suite.Require().Equal(entity.UserEntity{
		Id:    actualEntity.Id,
		Phone: phone,
		Login: login,
		OAuthByClientId: []entity.TokenByClientId{
			{token, clientId},
		},
	}, result[0])

	//over
	suite.CleanCollection(context.Background(), suite.CollectionName)
}

func (suite *TestAuthRepo) TestRepo_UpsertUserByPhone_NewClientForUser() {
	//given
	value := authValue()
	suite.CreateSomething(
		context.Background(),
		suite.CollectionName,
		&value)

	login := "sportman"
	phone := "+79991112233"
	token := "token1"
	clientId := "client1"

	//when
	err := suite.Repo.UpsertUserByPhone(
		phone,
		login,
		token,
		clientId)
	suite.Require().NoError(err, "Err then calling method InsertUser")

	result := suite.FindAll(context.Background(), suite.CollectionName, func(cursor *mongo.Cursor) interface{} {
		dbEntity := entity.UserEntity{}
		err := cursor.Decode(&dbEntity)
		suite.Require().NoError(err, "Can not decode entity")

		return dbEntity
	})

	suite.Require().NotEmptyf(result, "Collection should not be empty")

	actualEntity := result[0].(entity.UserEntity)
	suite.Require().Len(actualEntity.OAuthByClientId, 2)
	suite.Require().Contains(actualEntity.OAuthByClientId, entity.TokenByClientId{Token: token, ClientId: clientId})

	//over
	suite.CleanCollection(context.Background(), suite.CollectionName)
}

func (suite *TestAuthRepo) TestRepo_UpsertUserByPhone_ReauthorizedWithNoExpiredToken() {
	//given
	value := authValue()
	suite.CreateSomething(
		context.Background(),
		suite.CollectionName,
		&value)

	login := "sportman"
	phone := "+79991112233"
	token := "token"
	clientId := "clientId"

	//when
	err := suite.Repo.UpsertUserByPhone(
		phone,
		login,
		token,
		clientId)
	suite.Require().NoError(err, "Err then calling method InsertUser")

	result := suite.FindAll(context.Background(), suite.CollectionName, func(cursor *mongo.Cursor) interface{} {
		dbEntity := entity.UserEntity{}
		err := cursor.Decode(&dbEntity)
		suite.Require().NoError(err, "Can not decode entity")

		return dbEntity
	})

	suite.Require().NotEmptyf(result, "Collection should not be empty")

	actualEntity := result[0].(entity.UserEntity)
	suite.Require().Len(actualEntity.OAuthByClientId, 1)
	suite.Require().Contains(actualEntity.OAuthByClientId, entity.TokenByClientId{Token: token, ClientId: clientId})

	//over
	suite.CleanCollection(context.Background(), suite.CollectionName)
}
