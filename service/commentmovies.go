package service

import (

	//"MoviesDB/entities"
	//"MoviesDB/models"
	"MoviesDB/constants"
	"MoviesDB/entities"
	"MoviesDB/models"
	"MoviesDB/restapi/operations/add_comment"
	"fmt"

	//"context"
	//"fmt"
	"log"

	"github.com/go-openapi/runtime/middleware"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

func AddComment(params add_comment.PostcommentsParams) middleware.Responder {
	log.Println("Processing request to AddComment.")
	//mongoObj := connect.GetMongoObject()

	userName := params.Body.UserName
	movieName := params.Body.MovieName
	movieComment := params.Body.MovieComment
	fmt.Println("Username : ", userName, " movieName: ", movieName, "movieComment: ", movieComment)
	if UserNameIsValid(userName) == userName && len(userName) > 0 {
		//Checking if provided moviename exist in DB.
		if MovieNameIsValid(movieName) == movieName && len(movieName) > 0 && len(movieComment) > 0 {
			log.Println("User is valid in AddComment.")

			searchResult, err := ReadDocument(entities.IMDBRegistry{MovieName: movieName}, &entities.IMDBRegistry{})
			//searchResult := ReturnAllRecords(mongoObj, bson.M{"MovieName": movieName})
			log.Println("SearchResult is :", searchResult)
			fmt.Println("Err: ", err)

			if err != nil {
				log.Fatalln("Error while reading DB in AddComment: ", err.Error())
				errMsg := err.Error()
				return add_comment.NewPostcommentsInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
			}
			if searchResult != nil {
				log.Println("UserInfo is found.")
				result := searchResult.(entities.IMDBRegistry)
				fmt.Println("Result is :", result)

			}
		} else { // If moviename is invalid.
			errMsg := constants.INVALID_MOVIENAME + constants.REQUEST_FAILED
			return add_comment.NewPostcommentsInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
		}
	}
	//Return error if above conditions are not satisfied.
	errMsg := constants.INVALID_USER + constants.REQUEST_FAILED
	return add_comment.NewPostcommentsInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
}
