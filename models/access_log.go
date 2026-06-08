package models

import (
	"time"

)

type AccessLog struct {
	IPAddress string `bson:"ip_address" json:"ip_address"`
	Endpoint  string `bson:"endpoint" json:"endpoint"`
	Status    string `bson:"status" json:"status"`
	Error     string `bson:"error" json:"error"`
	Time      time.Time `bson:"time" json:"time"`
}