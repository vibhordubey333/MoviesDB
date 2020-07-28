package service

import (
	"MoviesDB/constants"
	connect "MoviesDB/dbconnection"
	"MoviesDB/entities"
	"context"
	"log"
)

//Hardcoded data , initialized from init method.
func DummyData() error {
	mongoObj := connect.GetMongoObject()
	if mongoObj == nil {
		log.Println("Error while conencting with DB .")
	}

	collection := mongoObj.Database(constants.DB_NAME).Collection(constants.COLLECTION_NAME)

	commentsObj1 := map[string]interface{}{
		"Alexi Murdoch":  "Phenomenal movie.",
		"Olafur Arnalds": "Awesome.",
	}
	commentsObj2 := map[string]interface{}{
		"Tom Hardy":       "Nice movie.",
		"Chris Hemsworth": "Good movie.",
		"Alexi Murdoch":   "Phenomenal movie.",
		"Olafur Arnalds":  "Awesome.",
	}
	insertObj1 := entities.IMDBRegistry{"Shawshank Redemption", "9.3", 2, commentsObj1}

	insertObj2 := entities.IMDBRegistry{"Hannibal", "9.0", 4, commentsObj2}

	multipleDBS := []interface{}{insertObj1, insertObj2}

	insertObj, err := collection.InsertMany(context.TODO(), multipleDBS)

	if err != nil {
		log.Fatalf("Error while inserting many documnets: Error :", err)
		return err
	}
	log.Println("Dummy data response :", insertObj)
	return nil
}
