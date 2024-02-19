package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommonFields struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	User      primitive.ObjectID `bson:"user"`
	Post      primitive.ObjectID `bson:"post"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type Like struct {
	CommonFields
}

type Comment struct {
	CommonFields
	Message *string `bson:"message"`
}
