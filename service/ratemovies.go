package service

import (
	"MoviesDB/constants"
	"MoviesDB/entities"
	"MoviesDB/models"
	"MoviesDB/restapi/operations/add_comment"
	"MoviesDB/restapi/operations/add_movie"
	"fmt"
	"log"

	"github.com/go-openapi/runtime/middleware"
)

//Adding Rating to a movie
func AddRating(params add_movie.AddmovieratingParams) middleware.Responder {
	fmt.Println("Request recieved to add rating.")

	userName := params.Body.UserName
	movieName := params.Body.MovieName
	movieRating := params.Body.MovieRating
	var userAlreadyExistCheck bool = false

	fmt.Println("Username : ", userName, " movieName: ", movieName, "movieComment: ", movieRating)
	if UserNameIsValid(userName) == userName && len(userName) > 0 {
		//Checking if provided moviename exist in DB.
		if MovieNameIsValid(movieName) == movieName && len(movieName) > 0 {

			searchResponse, err := ReadDocument(entities.IMDBRegistry{MovieName: movieName}, &entities.IMDBRegistry{})

			if err != nil {
				errMsg := "Details not found in DB."
				return add_movie.NewAddmovieratingInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
			}
			fmt.Println("SearchResponse : ", searchResponse)
			if searchResponse != nil {
				resultSet := searchResponse.(entities.IMDBRegistry)
				fmt.Println("Rating object : ", resultSet.Ratings)
				//TODO : We can remove this check as above we are already checking for the valid user.
				for _, v := range resultSet.Ratings {
					if v.UserName == userName {
						userAlreadyExistCheck = true
					}
				}
				_ = userAlreadyExistCheck
				filterObject := make(map[string]interface{})
				filterObject["MovieName"] = entities.IMDBRegistry{MovieName: movieName}
				updatedRating := make(map[string]interface{}, 0)

				newRating := entities.Ratings{movieName, userName, movieRating}
				resultSet.Ratings = append(resultSet.Ratings, newRating)
				//In below "ratings" case sensitivity matters otherwise an additional ratings field will be created.
				updatedRating["ratings"] = resultSet.Ratings
				updatedResponse, err := UpdateDocument(entities.IMDBRegistry{MovieName: movieName}, "$set", updatedRating)
				if err != nil {
					errMsg := constants.INVALID_MOVIENAME + constants.REQUEST_FAILED
					return add_comment.NewPostcommentsInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
				} else {
					log.Println("Rating saved to DB :", updatedResponse)
					successMsg := fmt.Sprintf("Rating successfully added to DB.")
					return add_comment.NewPostcommentsOK().WithPayload(&add_comment.PostcommentsOKBody{Code: constants.SUCCESS_CODE, Message: successMsg})
				}

			} else {
				errMsg := "Details not found in DB."
				return add_movie.NewAddmovieratingInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
			}
		} else {
			errMsg := "Invalid details."
			return add_movie.NewAddmovieratingInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
		}
	}

	return nil
}
