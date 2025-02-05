package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
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

func disconn(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func GetCampaign()  {}
func SaveCampaign() {}
func RegisterUser(user UserRegister) bool {
	client := conn()
	defer disconn(client)
	db := os.Getenv("DB_NAME")
	coll := client.Database(db).Collection("users")

	err := coll.FindOne(context.TODO(), bson.D{})
	return false
}
