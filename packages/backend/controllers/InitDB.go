package controllers

import (
	"context"
	"time"

	"taskism/models"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 5 * time.Second}, "Taskism", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	user := new(models.User)
	mgm.Coll(user).Indexes()
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	mgm.Coll(user).Indexes().CreateOne(context.Background(), indexModel)
}
