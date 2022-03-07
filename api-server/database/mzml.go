package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func InsertMzML(mz interface{}) (string, error) {
	one, err := mongoDB.Collection("mzml").InsertOne(context.Background(), mz)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return one.InsertedID.(primitive.ObjectID).String(), nil
}
