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
	var userAlreadyExistCheck bool = false

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

			if searchResult != nil {
				result := searchResult.(entities.IMDBRegistry)
				//Creating a filter object to search and update the comment.

				fmt.Println("Comments Object: ", result.Comments)

				for _, v := range result.Comments {
					if v.UserName == userName {
						userAlreadyExistCheck = true
					}
				}
				filterObject := make(map[string]interface{})
				filterObject["MovieName"] = entities.IMDBRegistry{MovieName: movieName}
				updatedComment := make(map[string]interface{}, 0)

				if userAlreadyExistCheck == false {
					// If code comes here it means user doesn't exist in database.
					//Creating update object to update the comments
					log.Println("User is new .")
					newComment := entities.Comments{movieName, userName, movieComment}
					result.Comments = append(result.Comments, newComment)
					//In below "comments" case sensitivity matters otherwise an additional comments field will be created.
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
				} else {
					log.Println("User already exist.")

					searchResult, err = ReadDocument(entities.IMDBRegistry{MovieName: movieName}, &entities.IMDBRegistry{})

					if err != nil {
						log.Fatalln("Error while reading DB in AddComment: ", err.Error())
						errMsg := err.Error()
						return add_comment.NewPostcommentsInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
					}
					result = searchResult.(entities.IMDBRegistry)

					var imdbRegistryObj entities.IMDBRegistry
					updatedCommentLoc := imdbRegistryObj.Comments

					for _, v := range result.Comments {
						if v.UserName == userName {
							tempObj := entities.Comments{v.MovieNameCom, v.UserName, movieComment}
							updatedCommentLoc = append(updatedCommentLoc, tempObj)
						} else {
							tempObj := entities.Comments{v.MovieNameCom, v.UserName, v.Comment}
							updatedCommentLoc = append(updatedCommentLoc, tempObj)
						}
					}
					updatedComment["comments"] = updatedCommentLoc
					fmt.Println("Updated Comments: ", updatedComment)
					//Delting old document.
					/*deleteResponse, err := DeleteDocument(entities.IMDBRegistry{MovieName: movieName,Comments: })
					fmt.Println("DeleteResponse : ", deleteResponse)
					if err != nil {
						fmt.Println("User already exist error while deleting document.")
					}
					*/
					updatedResponse, err := UpdateDocument(entities.IMDBRegistry{MovieName: movieName}, "$set", updatedComment)
					fmt.Println("Updated Response := ", updatedResponse)
					if err != nil {
						errMsg := constants.INVALID_MOVIENAME + constants.REQUEST_FAILED
						return add_comment.NewPostcommentsInternalServerError().WithPayload(&models.Error{Code: constants.INTERNAL_ERROR_CODE, Message: &errMsg})
					}

					//It means user already exist so we will replace there comment with new one.
					for _, v := range result.Comments {
						if v.UserName == userName {
							v.Comment = ""
							v.Comment = movieComment
							newComment := entities.Comments{movieName, userName, v.Comment}
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
					}
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
