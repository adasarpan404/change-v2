package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Content      *string            `bson:"content"`
	Author       primitive.ObjectID `bson:"author"`
	Attachments  *[]string          `bson:"image"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
	LikeCount    int32              `bson:"likeCount"`
	CommentCount int32              `bson:"commentCount"`
}
