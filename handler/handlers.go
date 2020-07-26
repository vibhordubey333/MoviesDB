package handlers

import (
	"MoviesDB/restapi/operations/add_comment"
	"MoviesDB/restapi/operations/add_movie"
	"MoviesDB/restapi/operations/movieinfo"
	"log"

	"github.com/go-openapi/runtime/middleware"
	//"go.uber.org/zap"
)

func AddComments(params add_comment.AddCommentsParams) middleware.Responder {
	log.Println("Request recieved by AddComments handler.")

}

func AddMovies(params add_movie.AddMovieParams) middleware.Responder {

}

func AddRating(params add_movie.AddRatingBody) middleware.Responder {
	return nil
}

func GetMovieInfo(params movieinfo.GetMovieInfoParams) middleware.Responder {

}
