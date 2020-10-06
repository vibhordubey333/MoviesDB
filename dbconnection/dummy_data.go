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

	ratingObj1 := entities.Ratings{"Shawshank Redemption", "Alexi Murdoch", 5.5}
	ratingObj2 := entities.Ratings{"Shawshank Redemption", "Alexi Murdoch_2", 4.5}
	ratingObj3 := entities.Ratings{"Shawshank Redemption", "Alexi Murdoch_3", 7.8}
	ratingObj4 := entities.Ratings{"Hannibal", "Tom Hardy", 7.9}
	ratingObj5 := entities.Ratings{"Hannibal", "Tom Hardy_1", 9.1}
	ratingObj6 := entities.Ratings{"Hannibal", "Tom Hardy_2", 2.3}

	commentsObj1 := entities.Comments{"Shawshank Redemption", "Alexi Murdoch", "Phenomenal movie."}
	commentsObj2 := entities.Comments{"Shawshank Redemption", "Alexi Murdoch_2", "Phenomenal movie_2."}
	commentsObj3 := entities.Comments{"Shawshank Redemption", "Alexi Murdoch_3", "Phenomenal movie_3."}
	commentsObj4 := entities.Comments{"Hannibal", "Tom Hardy", "Nice Movie"}
	commentsObj5 := entities.Comments{"Hannibal", "Tom Hardy_1", "Nice Movie_1"}
	commentsObj6 := entities.Comments{"Hannibal", "Tom Hardy_2", "Nice Movie_2"}

	ratingsCollection1 := []entities.Ratings{ratingObj1, ratingObj2, ratingObj3}
	ratingsCollection2 := []entities.Ratings{ratingObj4, ratingObj5, ratingObj6}

	commentsCollection1 := []entities.Comments{commentsObj1, commentsObj2, commentsObj3}
	commentsCollection2 := []entities.Comments{commentsObj4, commentsObj5, commentsObj6}

	insertObj1 := entities.IMDBRegistry{"Shawshank Redemption", ratingsCollection1, 2, commentsCollection1}

	insertObj2 := entities.IMDBRegistry{"Hannibal", ratingsCollection2, 3, commentsCollection2}

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
