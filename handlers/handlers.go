package handlers

import (
	"MoviesDB/restapi/operations/add_movie"
	"MoviesDB/restapi/operations/movieinfo"
	"MoviesDB/service"
	"log"

	"github.com/go-openapi/runtime/middleware"
	//"go.uber.org/zap"
)

// func AddComments(params add_comment.AddCommentsParams) middleware.Responder {
// 	log.Println("Request recieved by AddComments handler.")

// }

func AddMovies(params add_movie.PostmovieParams) middleware.Responder {
	log.Println("Request recieved by AddMovies handler.")
	return service.AddMovies(params)
}

// func AddRating(params add_movie.AddRatingBody) middleware.Responder {
// 	return nil
// }

func GetMovieInfo(params movieinfo.GetmoviesinfoParams) middleware.Responder {
	log.Println("Request recieved by GetMoviesInfo handler.")
	return service.GetMoviesInfo(params)
}
