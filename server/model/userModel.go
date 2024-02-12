package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName *string            `bson:"firstName" validate:"required,min=2,max=100"`
	LastName  *string            `bson:"lastName" validate:"required,min=2,max=100"`
	Password  *string            `bson:"password" validate:"required,min=6" json:"password,omitempty"`
	Email     *string            `bson:"email" validate:"email,required"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
