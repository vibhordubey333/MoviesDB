package service

import (
	"MoviesDB/constants"
	connect "MoviesDB/dbconnection"
	"MoviesDB/entities"
	"MoviesDB/models"
	"MoviesDB/restapi/operations/add_movie"
	"context"
	"log"

	"github.com/go-openapi/runtime/middleware"
)

func AddMovies(params add_movie.PostmovieParams) middleware.Responder {
	log.Println("Processing request AddMovies. ")

	userName := params.Body.UserName

	if UserNameIsValid(userName) == userName {
		log.Println(userName, "logged in. Welcome.")

		//Moviename should be provided.
		if params.Body.MovieName == "" { // Return with error in case of MovieName is left empty.
			log.Fatalf("Movie name not provided. %v", constants.REQUEST_FAILED)
			errMsg := "Please provide a movie name."
			return add_movie.NewPostmovieInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
		} else {

			//Getting mongoDB object.
			mongoObj := connect.GetMongoObject()
			collection := mongoObj.Database(constants.DB_NAME).Collection(constants.COLLECTION_NAME)

			//Storing comments in map[string]interface then feeding to DB.
			mapObj := make(map[string]interface{}, 0)
			mapObj[userName] = params.Body.MovieComments
			insertObj1 := entities.IMDBRegistry{params.Body.MovieName, params.Body.MovieRating, 0, mapObj}
			_, err := collection.InsertOne(context.Background(), insertObj1)
			if err != nil {
				errMsg := "Error while saving details to DB."
				log.Fatalln(errMsg)
				return add_movie.NewPostmovieInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
			}
		}
		msg := params.Body.MovieName + ", successfully saved."
		return add_movie.NewPostmovieOK().WithPayload(&add_movie.PostmovieOKBody{Code: constants.SUCCESS_CODE, Message: msg})
	}
	//Return with error in case of incorrect username.
	errMsg := constants.INVALID_USER
	log.Fatalln(errMsg, constants.REQUEST_FAILED)
	return add_movie.NewPostmovieInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
}
