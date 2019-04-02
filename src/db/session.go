package db

import (
	"context"
	"time"
	"utils"

	"apps"
	"settings"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

func timeOut(secondToTimeout int) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(secondToTimeout)*time.Second)
	return ctx
}

func NewConnection() (*mongo.Client, error) {
	client, err := mongo.Connect(timeOut(10), settings.ConnectionString)
	utils.CheckError(err)

	err = client.Ping(timeOut(10), readpref.Primary())

	return client, err
}

func GetCollection(collectionName string) *mongo.Collection {
	client, err := NewConnection()
	utils.CheckError(err)

	return client.Database(settings.DatabaseName).Collection(collectionName)
}

func GetAllInCollection(collectionName string) []bson.M {
	collection := GetCollection(collectionName)
	cur, err := collection.Find(timeOut(10), nil)
	var result []bson.M

	if err == mongo.ErrNilDocument {
		// if err.Error() == "document is nil" {
		return result
	}
	utils.CheckError(err)

	defer cur.Close(timeOut(10))

	for cur.Next(timeOut(10)) {
		var _result bson.M
		err := cur.Decode(&_result)
		utils.CheckError(err)

	}
	return result
}

func InsertIntoCollection(collectionName string, instance apps.GenericModel) error {
	collection := GetCollection(collectionName)
	_, err := collection.InsertOne(context.Background(), instance)

	return err
}
