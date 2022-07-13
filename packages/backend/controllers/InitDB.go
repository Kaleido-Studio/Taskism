package controllers

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 5 * time.Second}, "Taskism", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
}
