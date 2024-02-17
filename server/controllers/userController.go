package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
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

		c.JSON(
			http.StatusOK,
			gin.H{
				"status": true,
				"user":   user,
			})
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

		c.JSON(
			http.StatusCreated,
			gin.H{
				"relationship": relationShipObj,
			})
	}
}

// This function is used to get followers
func GetFollowers() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		perPage, err := strconv.Atoi(c.Query("limit"))
		if err != nil || perPage < 1 {
			perPage = 10
		}

		skip := (page - 1) * perPage

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		userId, ok := c.Get("userId")
		defer cancel()

		if !ok {
			helpers.ErrorResponse(c, http.StatusBadRequest, "User Id not found in the context")
			return
		}

		objectUserId, err := primitive.ObjectIDFromHex(fmt.Sprint(userId))
		defer cancel()

		if err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "Invalid User Id Format")
			return
		}

		totalCount, err := relationShipCollection.CountDocuments(ctx, bson.M{"follower": objectUserId})
		defer cancel()
		if err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		cursor, err := relationShipCollection.Find(ctx, bson.M{"follower": objectUserId}, options.Find().SetSkip(int64(skip)).SetLimit(int64(perPage)))
		defer cancel()
		if err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		var results []model.UserRelationShipModel
		if err := cursor.All(ctx, &results); err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		hasPrevPage := page > 1

		// Check if there is a next page
		hasNextPage := (page-1)*perPage+len(results) < int(totalCount)
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":      true,
				"data":        results,
				"total":       totalCount,
				"hasPrevPage": hasPrevPage,
				"hasNextPage": hasNextPage,
			})
	}
}

// This function is used to get following
func GetFollowing() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))

		if err != nil || page < 1 {
			page = 1
		}

		perPage, err := strconv.Atoi(c.Query("limit"))

		if err != nil || perPage < 1 {
			perPage = 10
		}

		skip := (page - 1) * perPage

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		userId, ok := c.Get("userId")

		defer cancel()

		if !ok {
			helpers.ErrorResponse(c, http.StatusBadRequest, "User ID not found in the context")
			return
		}

		objectUserId, err := primitive.ObjectIDFromHex(fmt.Sprint(userId))
		defer cancel()

		if err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "Invalid User Id Format")
			return
		}

		totalCount, err := relationShipCollection.CountDocuments(ctx, bson.M{"following": objectUserId})
		defer cancel()

		if err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		cursor, err := relationShipCollection.Find(ctx, bson.M{"following": objectUserId}, options.Find().SetSkip(int64(skip)).SetLimit(int64(perPage)))
		defer cancel()
		if err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		var results []model.UserRelationShipModel
		if err := cursor.All(ctx, &results); err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		hasPrevPage := page > 1

		// Check if there is a next page
		hasNextPage := (page-1)*perPage+len(results) < int(totalCount)
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":      true,
				"data":        results,
				"total":       totalCount,
				"hasPrevPage": hasPrevPage,
				"hasNextPage": hasNextPage,
			})
	}

}
