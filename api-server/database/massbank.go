package database

import (
	"context"
	"log"

	"github.com/uly55e5/MassBankRepo/api-server/massbank"
	"github.com/uly55e5/MassBankRepo/api-server/mberror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func Count() (int64, error) {
	return mongoDB.Collection("massbank").EstimatedDocumentCount(context.Background())
}

func GetAllSpectra(skip int64, limit int64) (interface{}, error) {
	var filter = bson.D{}
	var projection = bson.D{}
	return GetSpectra(skip, limit, filter, projection)
}

func GetSpectraInfo(skip int64, limit int64) (interface{}, error) {
	var filter = bson.D{}
	var projection = bson.D{{"accession", 1}}
	return GetSpectra(skip, limit, filter, projection)
}

func GetSpectra(skip int64, limit int64, filter bson.D, projection bson.D) (interface{}, error) {
	cursor, err := mongoDB.Collection("massbank").Find(context.Background(), filter, options.Find().SetLimit(limit).SetSkip(skip).SetProjection(projection))
	if mberror.Check(err) {
		return nil, err
	}
	var result = []bson.M{}
	err = cursor.All(context.Background(), &result)
	return result, err
}

func GetSpectrum(acc string) (interface{}, error) {
	var result bson.M
	err := mongoDB.Collection("massbank").FindOne(context.Background(), bson.D{{"accession", acc}}).Decode(&result)
	if mberror.Check(err) {
		return nil, err
	}
	return result, nil
}
