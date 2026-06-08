package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type WhitelistIP struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IPAddress string             `bson:"ip_address" json:"ip_address"`
	Enabled   bool               `bson:"enabled" json:"enabled"`
}
