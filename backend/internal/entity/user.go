package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	Id              primitive.ObjectID `bson:"_id"`
	Phone           string             `bson:"phone"`
	Login           string             `bson:"login"`
	OAuthByClientId []TokenByClientId  `bson:"tokenByClient"`
}

type TokenByClientId struct {
	Token    string `bson:"token"`
	ClientId string `bson:"clientId"`
}
