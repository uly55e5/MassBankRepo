package database

import (
	"context"
	"github.com/uly55e5/MassBankRepo/api-server/massbank"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func InsertMassbank(mb *massbank.Massbank) (string, error) {
	one, err := mongoDB.Collection("massbank").InsertOne(context.Background(), mb)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return one.InsertedID.(primitive.ObjectID).String(), nil
}

func ClearMassbankCollection() error {
	_, err := mongoDB.Collection("massbank").DeleteMany(context.Background(), bson.D{})
	return err
}
