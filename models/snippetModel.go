package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Snippet struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username *string            `json:"username"`
	Title    *string            `json:"title"`
	Language *string            `json:"language"`
	Code     *string            `json:"code"`
}
