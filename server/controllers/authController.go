package controllers

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/adasarpan404/change/helpers"
	"github.com/adasarpan404/change/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// This function is used to hash password
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

// This function is used to verify password
func verifyPassword(providedPassword string, userPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		check = false
		msg = "password is incorrect"
	}
	return check, msg
}

// This function is used to Signup
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

		count, err := userCollection.CountDocuments(
			ctx,
			bson.M{
				"email": user.Email,
			})
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
		token, err := helpers.GenerateToken(
			*user.Email,
			*user.FirstName,
			*user.LastName,
			user.ID.Hex(),
		)

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

		c.JSON(
			http.StatusCreated,
			gin.H{
				"status": true,
				"user":   userObj,
				"token":  token,
			})
	}
}

// This function is used to login
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.TODO(), 100*time.Second)
		var user, foundUser model.User

		if err := c.BindJSON(&user); err != nil {
			helpers.ErrorResponse(c, http.StatusBadRequest, err.Error())
		}

		err := userCollection.FindOne(
			ctx,
			bson.M{"email": user.Email},
		).Decode(&foundUser)
		defer cancel()

		if err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "Email is Incorrect")
			return
		}

		passwordIsValid, msg := verifyPassword(*foundUser.Password, *user.Password)
		defer cancel()

		if !passwordIsValid {
			helpers.ErrorResponse(c, http.StatusInternalServerError, msg)
			return
		}

		if foundUser.Email == nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "user not found")
			return
		}

		token, err := helpers.GenerateToken(
			*foundUser.Email,
			*foundUser.FirstName,
			*foundUser.LastName,
			foundUser.ID.Hex(),
		)
		if err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": true,
				"user":   foundUser,
				"token":  token,
			})
	}
}

// This function is used to authentice the routes
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		res := strings.Split(bearerToken, " ")
		clientToken := res[1]
		if clientToken == "" {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "No Authentication Header Provided")
			c.Abort()
			return
		}
		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			helpers.ErrorResponse(c, http.StatusInternalServerError, err)
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("firstName", claims.FirstName)
		c.Set("lastName", claims.LastName)
		c.Set("userId", claims.ID)
		c.Next()
	}
}
