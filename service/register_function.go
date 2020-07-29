package service

import (
	"MoviesDB/restapi/operations/add_comment"
	"MoviesDB/restapi/operations/add_movie"
	"MoviesDB/restapi/operations/movieinfo"

	"github.com/go-openapi/runtime/middleware"
)

type ServiceInterface interface {
	AddComment(params add_comment.PostcommentsParams) middleware.Responder
	ReadDocument(filterobj interface{}, outputObject interface{}) (interface{}, error)
	UpdateDocument(filterObject interface{}, operation string, update map[string]interface{}) (int64, error)
	GetMoviesInfo(params movieinfo.GetmoviesinfoParams) middleware.Responder
	UserNameIsValid(uname string) string
	MovieNameIsValid(moviename string) string
	AddMovies(params add_movie.PostmovieParams) middleware.Responder
}
