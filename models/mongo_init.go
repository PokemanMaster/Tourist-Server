package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Mongo = InitMongo()

// InitMongo 初始化MongoDB
func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "123456",
	}).ApplyURI("mongodb://192.168.127.163:27017"))
	if err != nil {
		log.Println("Connection MongoDB Error:", err)
		return nil
	}
	return client.Database("qi")
}
