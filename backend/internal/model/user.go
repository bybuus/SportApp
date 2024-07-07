package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `bson:"_id"`
	Login string             `bson:"login"`
	Phone string             `bson:"phone"`
}

type UserAuthorized struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}
