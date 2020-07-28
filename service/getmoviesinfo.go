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
	//var results bson.M
	var store primitive.M

	if movieName != "" { // Proceed , if moviename is not empty.
		//results = movieName
		mongoObj := connect.GetMongoObject()
		tmp := bson.M{
			"moviename": movieName,
		}
		if mongoObj == nil {
			errMsg := "Error while establishing connection with DB."
			return movieinfo.NewGetmoviesinfoInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
		}

		collection := mongoObj.Database(constants.DB_NAME).Collection(constants.COLLECTION_NAME)

		//Searching for MovieName in DB.
		err := collection.FindOne(context.Background(), tmp).Decode(&store)
		if err != nil {
			errMsg := "Error while finding movie in DB. " + err.Error()
			return movieinfo.NewGetmoviesinfoInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
		}
		fmt.Printf("Result is : %+v\n", store)

		//Decoding object.

	} else { // If movie name is empty return with error.
		errMsg := constants.MOVIE_NAME_EMPTY
		movieinfo.NewGetmoviesinfoInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
	}
	//fmt.Println("Movie Name: ", results)
	//return nil
	movieName = fmt.Sprintf("%v", store["moviename"])
	ratingCount, _ := strconv.Atoi(fmt.Sprintf("%v", store["ratinggivencount"]))
	avgRating := fmt.Sprintf("%v", store["rating"])

	userComments := make([]string, 0)
	//userComments = append(userComments, store["comments"].(primitive.M))
	//fmt.Println("Comments: ", userComments)
	for k, v := range store["comments"].(primitive.M) {
		fmt.Println("K:", k, "V:", v)
		temp := k + ":" + v.(string)
		userComments = append(userComments, temp)
	}
	//return nil
	return movieinfo.NewGetmoviesinfoOK().WithPayload(&movieinfo.GetmoviesinfoOKBody{MovieName: movieName, PeopleRated: int64(ratingCount), AvgRating: avgRating, UserComments: userComments})
}

func MapToSlice(m map[string]interface{}) []string {

	slice := make([]string, 0)

	for k, _ := range m {
		slice = append(slice, k)
	}
	fmt.Println("Checking : ", slice)
	return slice
}
