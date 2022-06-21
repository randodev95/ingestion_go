package controllers

import (
	"context"
	"ingestion_api/configs"
	"ingestion_api/models"
	"ingestion_api/responses"
	"ingestion_api/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.UserValidation
		defer cancel()

		//fmt.Println(c.Request)

		if ctxc, err := utils.ValidBindCheck(c, &user); err != nil {
			c = ctxc
			return
		}

		new_uuid := uuid.New().String()
		new_uuid = strings.Replace(new_uuid, "-", "", -1)

		newUser := models.User{
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			EmailId:    user.EmailId,
			UserId:     new_uuid,
			Address:    map[string]utils.Address{"primary_address": utils.GetRequestAdress(utils.GetIpFromRequest(c.Request))},
			DeviceInfo: utils.GetRequestDeviceInfo(c.Request),
		}

		result, err := userCollection.InsertOne(ctx, newUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.User_Response{Status: http.StatusInternalServerError, Message: "User Creation Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.User_Response{Status: http.StatusCreated, Message: "User Created SuccesSfully", Data: map[string]interface{}{"data": result}})
	}
}

func GetUserDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("UserId")

		var user models.User
		defer cancel()

		if err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, responses.User_Response{Status: http.StatusInternalServerError, Message: "Could find user with the parameter", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.User_Response{Status: http.StatusOK, Message: "user found", Data: map[string]interface{}{"data": user}})

	}
}

func EditUserDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("UserId")

		var user models.User
		defer cancel()

		if ctxc, err := utils.ValidBindCheck(c, &user); err != nil {
			c = ctxc
			return
		}

		update := bson.M{"first_name": user.FirstName, "last_name": user.LastName, "email_id": user.EmailId, "user_id": user.UserId, "address": user.Address, "device_info": user.DeviceInfo}
		result, err := userCollection.UpdateOne(ctx, bson.M{"user_id": userId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.User_Response{Status: http.StatusInternalServerError, Message: "Error Updating User", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"user_id": userId})
			if err != nil {
				c.JSON(http.StatusInternalServerError, "user_updation_unsuccessful")
			}
		}

		c.JSON(http.StatusOK, responses.User_Response{Status: http.StatusOK, Message: "User updation successful", Data: map[string]interface{}{"data": result}})

	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()

		result, err := userCollection.DeleteOne(ctx, bson.M{"user_id": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.User_Response{Status: http.StatusInternalServerError, Message: "unexpected error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, responses.User_Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User Not Found"}})
			return
		}

		c.JSON(http.StatusOK, responses.User_Response{Status: http.StatusOK, Message: "User Deleted", Data: map[string]interface{}{"data": result}})
	}
}
