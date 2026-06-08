package services

import (
	"context"
	"time"
	"webhook-project/config"
	"webhook-project/models"
)

func SaveLog(ip, endpoint, status, errMsg string) {

	log := models.AccessLog{
		IPAddress: ip,
		Endpoint:  endpoint,
		Status:    status,
		Error:     errMsg,
		Time:      time.Now(),
	}

	config.DB.Collection("access_logs").InsertOne(context.TODO(), log)
}