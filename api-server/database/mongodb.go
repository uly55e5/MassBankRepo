package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

const dbTimeout = 10 * time.Second

var mongoClient *mongo.Client
var mongoDB *mongo.Database

func InitMongoDB() {
	var err error
	dbhost := os.Getenv("MONGO_URI")
	dbname := os.Getenv("MONGO_DB_NAME")
	if mongoClient, err = mongo.NewClient(options.Client().ApplyURI(dbhost)); err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), dbTimeout)
	if err = mongoClient.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	mongoDB = mongoClient.Database(dbname)
	if mongoDB == nil {
		log.Fatal("Database not found")
	}
	options := options.IndexOptions{}
	_, err = mongoDB.Collection("massbank").Indexes().CreateOne(ctx, mongo.IndexModel{bson.D{{"accession", 1}}, options.SetName("accession_1").SetUnique(true)})
	if err != nil {
		log.Println("Error while index creation: ", err.Error())
	}
}

func CloseMongDB() {
	ctx, _ := context.WithTimeout(context.Background(), dbTimeout)
	if err := mongoClient.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
}
