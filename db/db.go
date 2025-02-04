package db

import (
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func conn() *mongo.Client {
	uri := os.Getenv("DB_STRING")

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}

func GetCampaign()  {}
func SaveCampaign() {}
