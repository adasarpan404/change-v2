package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/adasarpan404/change/helpers"
	"github.com/adasarpan404/change/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.TODO(), 100*time.Second)
		var user model.User
		if err := c.BindJSON(&user); err != nil {
			helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			helpers.ErrorResponse(c, http.StatusBadRequest, validationErr.Error())
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if count > 0 {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "this email already exists")
			return
		}
		password := HashPassword(*user.Password)
		user.Password = &password
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		token, err := helpers.GenerateToken(*user.Email, *user.FirstName, *user.LastName, user.ID.Hex())
		if err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		userObj, err := userCollection.InsertOne(ctx, user)
		if err != nil {
			msg := "User item was not created"
			helpers.ErrorResponse(c, http.StatusInternalServerError, msg)
			return
		}
		defer cancel()

		c.JSON(http.StatusCreated, gin.H{"user": userObj, "token": token})
	}
}