package controllers

import (
	"github.com/adasarpan404/change/database"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var relationShipCollection *mongo.Collection = database.OpenCollection(database.Client, "relationShip")
