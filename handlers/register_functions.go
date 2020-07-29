package handlers

import (
	"MoviesDB/restapi/operations/add_comment"
	"MoviesDB/restapi/operations/add_movie"
	"MoviesDB/restapi/operations/movieinfo"

	"github.com/go-openapi/runtime/middleware"
)

type HandlersInterface interface {
	AddComments(params add_comment.PostcommentsParams) middleware.Responder
	AddMovies(params add_movie.PostmovieParams) middleware.Responder
	GetMovieInfo(params movieinfo.GetmoviesinfoParams) middleware.Responder
}
