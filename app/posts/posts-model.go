package posts

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostModel struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title" bson:"title" binding:"required;"`
	Description string             `json:"description" bson:"description" binding:"required;"`
	Category    string             `json:"category" bson:"category" binding:"required"`
	Creator     primitive.ObjectID `json:"creator" bson:"creator"`
}
