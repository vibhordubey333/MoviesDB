package dbconnection

import (
	"MoviesDB/constants"

	"MoviesDB/entities"
	"context"
	"log"
)

//DummyData hardcoded data , initialized from init method.
func DummyData() error {
	mongoObj := GetMongoObject()
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

	_, err := collection.InsertMany(context.TODO(), multipleDBS)

	if err != nil {
		log.Fatalf("Error while inserting many documnets: Error :%v", err)
		return err
	}
	log.Println("Movie dummy data is loaded successfully.")
	return nil
}
func AddUsers() error {

	mongoObj := GetMongoObject()
	if mongoObj == nil {
		log.Println("Error while conencting with DB .")
	}

	collection := mongoObj.Database(constants.DB_NAME).Collection(constants.COLLECTION_NAME_USER)

	insertObj1 := entities.UserRegistry{"admin", "1-7-2020"}
	insertObj2 := entities.UserRegistry{"chief", "1-1-2020"}

	multipleDB := []interface{}{insertObj1, insertObj2}
	_, err := collection.InsertMany(context.TODO(), multipleDB)
	if err != nil {
		log.Fatalf("Error while inserting many  user documents. Error :%v", err)
		return err
	}
	log.Println("User dummy data is loaded successfully.")
	return nil
}
