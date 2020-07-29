package service

import (
	"MoviesDB/constants"
	"MoviesDB/entities"
	"MoviesDB/models"
	"MoviesDB/restapi/operations/add_comment"
	"log"

	"github.com/go-openapi/runtime/middleware"
)

func AddComment(params add_comment.PostcommentsParams) middleware.Responder {
	log.Println("Processing request to AddComment.")

	userName := params.Body.UserName
	movieName := params.Body.MovieName
	movieComment := make(map[string]interface{}, 0)
	movieComment["comments"] = params.Body.MovieComment
	//movieComment := params.Body.MovieComment
	if UserNameIsValid(userName) == userName && len(userName) > 0 {

		//Checking if provided moviename exist in DB.
		if MovieNameIsValid(movieName) == movieName && len(movieName) > 0 && len(movieComment) > 0 {

			searchResult, err := ReadDocument(entities.IMDBRegistry{Comments: movieComment}, &entities.IMDBRegistry{})
			if err != nil {

			}
			if searchResult != nil {
				result := searchResult.(entities.IMDBRegistry)
				commentUp := result.Comments
				_ = commentUp
				_, err := UpdateDocument(entities.IMDBRegistry{Comments: result.Comments}, "$push", nil)
				if err != nil {

				}
			}

			//UpdateDocument(entities.IMDBRegistry{Comments: movieComment})
		} else { // If moviename is invalid.
			errMsg := constants.INVALID_MOVIENAME + constants.REQUEST_FAILED
			return add_comment.NewPostcommentsInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
		}
	}

	//Return error if above conditions are not satisfied.
	errMsg := constants.INVALID_USER + constants.REQUEST_FAILED
	return add_comment.NewPostcommentsInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
}
