package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/adasarpan404/change/helpers"
	"github.com/adasarpan404/change/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.TODO(), 100*time.Second)
		var post model.Post
		if err := c.BindJSON(&post); err != nil {
			helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
		}
		validationErr := validate.Struct(post)
		if validationErr != nil {
			helpers.ErrorResponse(c, http.StatusBadRequest, validationErr.Error())
		}

		userId, ok := c.Get("userId")
		defer cancel()
		if !ok {
			helpers.ErrorResponse(c, http.StatusBadRequest, "User ID not found in context")
			return
		}

		objectUserId, err := primitive.ObjectIDFromHex(fmt.Sprint(userId))
		defer cancel()
		if err != nil {
			helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID format")
			return
		}
		post.Author = objectUserId
		post.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		post.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		post.ID = primitive.NewObjectID()
		postObj, err := postCollection.InsertOne(ctx, post)
		if err != nil {
			msg := "Room item was not created"
			helpers.ErrorResponse(c, http.StatusInternalServerError, msg)
			return
		}
		defer cancel()
		c.JSON(http.StatusCreated, gin.H{"room": postObj})
	}
}

func GetPost() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdatePost() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeletePost() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetHomeFeed() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
