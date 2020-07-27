package service

import (
	"MoviesDB/constants"
	connect "MoviesDB/dbconnection"
	"MoviesDB/models"
	"MoviesDB/restapi/operations/movieinfo"
	"context"
	"fmt"
	"log"
	"encoding/json"
	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/bson"
)

func GetMoviesInfo(params movieinfo.GetmoviesinfoParams) middleware.Responder {
	log.Println("Processing request GetMoviesInfo. ")

	movieName := params.Body.MovieName

	if movieName != "" { // Proceed , if moviename is not empty.

		mongoObj := connect.GetMongoObject()

		if mongoObj == nil {
			errMsg := "Error while establishing connection with DB."
			return movieinfo.NewGetmoviesinfoInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
		}

		collection := mongoObj.Database(constants.DB_NAME).Collection(constants.COLLECTION_NAME)

		var results bson.M

		//Searching for MovieName in DB.
		err := collection.FindOne(context.Background(), results).Decode(&results)
		if err != nil {
			errMsg := "Error while finding movie in DB. " + err.Error()
			return movieinfo.NewGetmoviesinfoInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
		}
		fmt.Printf("Result is : %+v\n", results)

	} else { // If movie name is empty return with error.
		errMsg := constants.MOVIE_NAME_EMPTY
		movieinfo.NewGetmoviesinfoInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
	}
	
	return movieinfo.NewGetmoviesinfoOK().WithPayload(&movieinfo.GetmoviesinfoOKBody{MovieName:results.})
}
