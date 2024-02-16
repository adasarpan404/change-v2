package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/adasarpan404/change/helpers"
	"github.com/adasarpan404/change/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// This function is used to profile information
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Get("userId")
		if !ok {
			helpers.ErrorResponse(c, http.StatusBadRequest, "User ID not found in context")
			return
		}
		objectUserId, err := primitive.ObjectIDFromHex(fmt.Sprint(userId))
		if err != nil {
			helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID format")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user model.User
		projection := bson.M{"password": 0}
		err = userCollection.FindOne(ctx, bson.M{"_id": objectUserId}, options.FindOne().SetProjection(projection)).Decode(&user)
		if err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// This function is used to follow
func Follow() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		userId, ok := c.Get("userId")
		givenUserId := c.Param("id")
		defer cancel()

		if !ok {
			helpers.ErrorResponse(c, http.StatusBadRequest, "User Id not found in the context")
			return
		}

		objectUserId, err := primitive.ObjectIDFromHex(fmt.Sprint(userId))
		defer cancel()

		if err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID Format")
		}

		objectGiveUserId, err := primitive.ObjectIDFromHex(givenUserId)
		defer cancel()

		if err != nil {
			helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid Room ID Format")
			return
		}

		relationShip := model.UserRelationShipModel{
			Follower:  objectGiveUserId,
			Following: objectUserId,
		}
		relationShipObj, err := relationShipCollection.InsertOne(ctx, relationShip)

		if err != nil {
			msg := "You are not following the user"
			helpers.ErrorResponse(c, http.StatusInternalServerError, msg)
			return
		}

		defer cancel()

		c.JSON(http.StatusCreated, gin.H{"relationship": relationShipObj})
	}
}

// This function is used to get followers
func GetFollowers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// This function is used to get following
func GetFollowing() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
