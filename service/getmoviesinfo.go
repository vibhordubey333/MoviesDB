package service

import (
	"MoviesDB/constants"
	connect "MoviesDB/dbconnection"
	"MoviesDB/models"
	"MoviesDB/restapi/operations/movieinfo"
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMoviesInfo(params movieinfo.GetmoviesinfoParams) middleware.Responder {

	log.Println("Processing request GetMoviesInfo. ")

	movieName := params.Body.MovieName
	var store primitive.M

	if movieName != "" { // Proceed , if moviename is not empty.
		mongoObj := connect.GetMongoObject()
		if mongoObj == nil {
			errMsg := "Error while establishing connection with DB."
			return movieinfo.NewGetmoviesinfoInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
		}

		srchObject := bson.M{
			"moviename": movieName,
		}

		collection := mongoObj.Database(constants.DB_NAME).Collection(constants.COLLECTION_NAME)

		//Searching for MovieName in DB.
		err := collection.FindOne(context.Background(), srchObject).Decode(&store)
		if err != nil {
			errMsg := "Error while finding movie in DB. "
			log.Println(errMsg + err.Error())
			return movieinfo.NewGetmoviesinfoInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
		}
		movieName = fmt.Sprintf("%v", store["moviename"])
		ratingCount, _ := strconv.Atoi(fmt.Sprintf("%v", store["ratinggivencount"]))
		avgRating := fmt.Sprintf("%v", store["rating"])

		//Storing commnents in string slice.
		userComments := make([]string, 0)
		for k, v := range store["comments"].(primitive.M) {
			temp := k + ":" + v.(string)
			userComments = append(userComments, temp)
		}
		log.Println(constants.REQUEST_SUCCESS)
		return movieinfo.NewGetmoviesinfoOK().WithPayload(&movieinfo.GetmoviesinfoOKBody{MovieName: movieName, PeopleRated: int64(ratingCount), AvgRating: avgRating, UserComments: userComments})
	}

	//In case empty movie name is provided return with error.
	errMsg := constants.MOVIE_NAME_EMPTY
	log.Println(constants.REQUEST_FAILED + errMsg)
	return movieinfo.NewGetmoviesinfoInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
}
