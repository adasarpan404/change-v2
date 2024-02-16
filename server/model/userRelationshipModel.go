package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRelationShipModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Follower  primitive.ObjectID `bson:"follower"`
	Following primitive.ObjectID `bson:"following"`
}
