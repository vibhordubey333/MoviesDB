package service

import (
	"MoviesDB/constants"
	"MoviesDB/models"
	"MoviesDB/restapi/operations/add_comment"
	"log"

	"github.com/go-openapi/runtime/middleware"
)

func AddComment(params add_comment.PostcommentsParams) middleware.Responder {
	log.Println("Processing request to AddComment.")

	//Return error if above conditions are not satisfied.
	errMsg := constants.INVALID_USER + constants.REQUEST_FAILED
	return add_comment.NewPostcommentsInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
}
