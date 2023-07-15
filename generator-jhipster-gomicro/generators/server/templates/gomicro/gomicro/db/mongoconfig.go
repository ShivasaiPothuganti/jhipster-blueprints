package config

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/micro/micro/v3/service/logger"
	app "<%= packageName %>/config"
)

var client *mongo.Client

func GetInstance() *mongo.Client {
	db_url :=app.GetVal("GO_MICRO_MONGODB_URL")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(db_url)
	client, _ = mongo.Connect(ctx, clientOptions)
	err := client.Ping(ctx, nil)
	if err != nil {
		logger.Errorf(err.Error())
    	return nil
	}

	return client
}