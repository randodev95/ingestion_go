package models

import (
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"ingestion_api/utils"
)

type User struct {
	FirstName  string                   `bson:"first_name" validate:"required"`
	LastName   string                   `bson:"last_name" validate:"required"`
	EmailId    string                   `bson:"email_id" validate:"required"`
	UserId     string                   `bson:"user_id, inline , omitempty" validate:"required"`
	Address    map[string]utils.Address `bson:"adress,inline,omitempty" validate:"required"`
	DeviceInfo utils.UserDeviceInfo     `bson:"device_info,inline,omitempty" validate:"required"`
}

type UserValidation struct {
	FirstName  string                   `bson:"first_name"`
	LastName   string                   `bson:"last_name"`
	EmailId    string                   `bson:"email_id"`
	UserId     string                   `bson:"user_id,omitempty"`
	Address    map[string]utils.Address `bson:"adress,omitempty,inline"`
	DeviceInfo utils.UserDeviceInfo     `bson:"device_info,omitempty"`
}
