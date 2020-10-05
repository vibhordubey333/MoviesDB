package service

import (
	"MoviesDB/restapi/operations/add_movie"
	"fmt"

	"github.com/go-openapi/runtime/middleware"
)

func AddRating(params add_movie.AddmovieratingParams) middleware.Responder {
	fmt.Println("Request recieved to add rating.")

	userName := params.Body.UserName
	movieName := params.Body.MovieName
	movieRating := params.Body.MovieRating
	var userAlreadyExistCheck bool = false

	fmt.Println("Username : ", userName, " movieName: ", movieName, "movieComment: ", movieRating)
	if UserNameIsValid(userName) == userName && len(userName) > 0 {
		_ = userAlreadyExistCheck
	}

	return nil
}
