package service

import (
	"MoviesDB/constants"
	connect "MoviesDB/dbconnection"
	"context"

	//"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserNameIsValid checks if username is valid.
func UserNameIsValid(uname string) string {

	var store primitive.M

	mongoObj := connect.GetMongoObject()

	if mongoObj == nil {
		errMsg := "Error while establishing connection with DB."
		log.Fatalln(errMsg + "While checking UserNameIsValid.")
		return ""
	}

	srchObject := bson.M{
		"username": uname,
	}
	collection := mongoObj.Database(constants.DB_NAME).Collection(constants.COLLECTION_NAME_USER)

	//Searching username in DB.
	err := collection.FindOne(context.Background(), srchObject).Decode(&store)
	if err != nil {
		errMsg := "Error while finding username in DB. "
		log.Println(errMsg + err.Error())
		return ""
	}

	//Type fetching from interface
	userName := store["username"].(string)

	return userName
}

//MovieNameIsValid checks if username is valid.
func MovieNameIsValid(moviename string) string {

	var store primitive.M

	mongoObj := connect.GetMongoObject()

	if mongoObj == nil {
		errMsg := "Error while establishing connection with DB."
		log.Fatalln(errMsg + "While checking MovieNameIsValid.")
		return ""
	}

	srchObject := bson.M{
		"moviename": moviename,
	}
	collection := mongoObj.Database(constants.DB_NAME).Collection(constants.COLLECTION_NAME)

	//Searching username in DB.
	err := collection.FindOne(context.Background(), srchObject).Decode(&store)
	if err != nil {
		errMsg := "Error while finding moviename in DB. "
		log.Println(errMsg + err.Error())
		return ""
	}

	//Type fetching from interface
	movieName := store["moviename"].(string)

	return movieName
}
