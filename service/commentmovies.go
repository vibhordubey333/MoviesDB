package service

import (
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

	userName := params.Body.UserName
	movieName := params.Body.MovieName
	movieComment := params.Body.MovieComment
	fmt.Println("Username : ", userName, " movieName: ", movieName, "movieComment: ", movieComment)
	if UserNameIsValid(userName) == userName && len(userName) > 0 {
		//Checking if provided moviename exist in DB.
		if MovieNameIsValid(movieName) == movieName && len(movieName) > 0 && len(movieComment) > 0 {
			log.Println("User is valid in AddComment.")

			searchResult, err := ReadDocument(entities.IMDBRegistry{MovieName: movieName}, &entities.IMDBRegistry{})
			if err != nil {
				log.Fatalln("Error while reading DB in AddComment: ", err.Error())
				errMsg := err.Error()
				return add_comment.NewPostcommentsInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
			}
			if err != nil {
				log.Fatalln("Error while reading DB in AddComment: ", err.Error())
				errMsg := err.Error()
				return add_comment.NewPostcommentsInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
			}
			if searchResult != nil {
				result := searchResult.(entities.IMDBRegistry)
				filterObject := make(map[string]interface{})
				filterObject["MovieName"] = entities.IMDBRegistry{MovieName: movieName}
				updatedComment := make(map[string]interface{}, 0)
				newComment := entities.Comments{movieName, userName, movieComment}
				result.Comments = append(result.Comments, newComment)
				updatedComment["comments"] = result.Comments
				updatedResponse, err := UpdateDocument(entities.IMDBRegistry{MovieName: movieName}, "$set", updatedComment)
				if err != nil {
					errMsg := constants.INVALID_MOVIENAME + constants.REQUEST_FAILED
					return add_comment.NewPostcommentsInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
				} else {
					log.Println("Comment saved to DB :", updatedResponse)
					successMsg := fmt.Sprintf("Comment successfully added to DB.")
					return add_comment.NewPostcommentsOK().WithPayload(&add_comment.PostcommentsOKBody{Code: constants.SUCCESS_CODE, Message: successMsg})
				}

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
