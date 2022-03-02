package database

import (
	"context"
	"github.com/uly55e5/MassBankRepo/api-server/massbank"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetSpectra(skip int64, limit int64) (interface{}, error) {
	cursor, err := mongoDB.Collection("massbank").Find(context.Background(), bson.D{}, options.Find().SetLimit(limit).SetSkip(skip))
	var result = []bson.M{}
	err = cursor.All(context.Background(), &result)
	return result, err
}
