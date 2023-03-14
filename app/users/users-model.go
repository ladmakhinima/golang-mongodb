package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	FrontendTech = iota + 1
	BackendTech
	FullstackTech
)

type UserModel struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Firstname string             `json:"firstname" binding:"required;min:3;type: varchar(100)" bson:"firstname"`
	Lastname  string             `json:"lastname" binding:"required;min:3;type: varchar(100)" bson:"lastname"`
	Age       int                `json:"age" binding:"required;type: int;" bson:"age"`
	Tech      int                `json:"tech" binding:"required;type: int;" bson:"tech"`
}
