package controllers

import (
	"github.com/adasarpan404/roomies-be/db"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

var userCollection *mongo.Collection = db.OpenCollection(db.Client, "user")

var roomCollection *mongo.Collection = db.OpenCollection(db.Client, "room")
