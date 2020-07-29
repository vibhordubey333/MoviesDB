package constants

//Contains constants used throughout the application for easy maintenance.
const (
	URI                  = "mongodb://localhost:27017"
	INTERNAL_ERROR_CODE  = "500"
	SUCCESS_CODE         = "200"
	DB_NAME              = "IMDB"
	COLLECTION_NAME      = "IMDBRegistry"
	MOVIE_NAME_EMPTY     = "Movie name is empty."
	REQUEST_FAILED       = "Request is failed."
	REQUEST_SUCCESS      = "Request is successful."
	COLLECTION_NAME_USER = "UserRegistry"
	INVALID_USER         = "Invalid user."
	INVALID_MOVIENAME    = "Invalid movie name"
	ERR_DB               = "Error while establishing connection with DB."
	NO_RECORD_FOUND      = "No record found."
)
