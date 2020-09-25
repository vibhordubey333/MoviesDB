package service

import (
	"MoviesDB/constants"
	connect "MoviesDB/dbconnection"
	"MoviesDB/entities"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReadDocument(filterobj interface{}, outputObject interface{}) (interface{}, error) {

	jsonData, err := json.Marshal(filterobj)
	if err != nil {
		return nil, err
	}
	var result interface{}
	var m interface{}
	json.Unmarshal(jsonData, &m)
	filter := m.(map[string]interface{})

	//Establishing connection with DB.
	mongoObj := connect.GetMongoObject()
	if mongoObj == nil {
		log.Fatalln(constants.ERR_DB, "In Find.")
		return nil, errors.New("Error while connecting with DB")
	}

	collection := mongoObj.Database(constants.DB_NAME).Collection(constants.COLLECTION_NAME)

	singleResult := collection.FindOne(context.TODO(), filter)

	if singleResult.Err() != nil {
		if singleResult.Err().Error() == "mongo: no documents in result" {
			return nil, errors.New(constants.NO_RECORD_FOUND)
		} else {
			return nil, singleResult.Err()
		}
	}
	err = singleResult.Decode(&result)
	if err != nil {
		return nil, err
	}
	bsonBytes, err := bson.Marshal(result)
	if err != nil {
		return nil, err
	}
	bson.Unmarshal(bsonBytes, outputObject)
	docValTmp := reflect.ValueOf(outputObject).Elem()
	return docValTmp.Interface(), nil
}

func UpdateDocument(filterObject interface{}, operation string, update map[string]interface{}) (int64, error) {

	mongoObj := connect.GetMongoObject()

	if mongoObj == nil {
		log.Fatalln(constants.ERR_DB)
		return 0, nil // TODO
	}
	collection := mongoObj.Database(constants.DB_NAME).Collection(constants.COLLECTION_NAME_USER)

	jsonData, err := json.Marshal(filterObject)
	if err != nil {
		return 0, err
	}

	var m interface{}
	json.Unmarshal(jsonData, &m)
	filter := m.(map[string]interface{})

	jsonDataUpdate, err := json.Marshal(update)
	if err != nil {
		return 0, err
	}
	var m1 interface{}
	json.Unmarshal(jsonDataUpdate, &m1)

	updateString := bson.M{operation: update}

	result, err := collection.UpdateOne(context.TODO(), filter, updateString)
	if err != nil {
		return 0, err
	} else {
		return result.ModifiedCount, nil
	}

}
func ReturnAllRecords(client *mongo.Client, filter bson.M) []*entities.IMDBRegistry {
	//When filter is empty it returns all the objects
	var poc_object []*entities.IMDBRegistry
	collection := client.Database(constants.DB_NAME).Collection("IMDBRegistry")

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error in ReturnAllRecords: ", err.Error())
	}

	for cur.Next(context.TODO()) {
		var records entities.IMDBRegistry
		err = cur.Decode(&records)
		if err != nil {
			fmt.Println("Error while iterating records :")
		}
		poc_object = append(poc_object, &records)
	}
	return poc_object
}
