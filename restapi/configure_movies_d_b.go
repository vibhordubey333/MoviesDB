// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"MoviesDB/restapi/operations"
	"MoviesDB/restapi/operations/add_comment"
	"MoviesDB/restapi/operations/add_movie"
	"MoviesDB/restapi/operations/movieinfo"
	"MoviesDB/restapi/operations/ratemovies"
)

//go:generate swagger generate server --target ..\..\MoviesDB --name MoviesDB --spec ..\api\movies_swagger.yml

func configureFlags(api *operations.MoviesDBAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MoviesDBAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AddCommentAddCommentsHandler == nil {
		api.AddCommentAddCommentsHandler = add_comment.AddCommentsHandlerFunc(func(params add_comment.AddCommentsParams) middleware.Responder {
			return middleware.NotImplemented("operation add_comment.AddComments has not yet been implemented")
		})
	}
	if api.AddMovieAddMovieHandler == nil {
		api.AddMovieAddMovieHandler = add_movie.AddMovieHandlerFunc(func(params add_movie.AddMovieParams) middleware.Responder {
			return middleware.NotImplemented("operation add_movie.AddMovie has not yet been implemented")
		})
	}
	if api.AddMovieAddRatingHandler == nil {
		api.AddMovieAddRatingHandler = add_movie.AddRatingHandlerFunc(func(params add_movie.AddRatingParams) middleware.Responder {
			return middleware.NotImplemented("operation add_movie.AddRating has not yet been implemented")
		})
	}
	if api.MovieinfoGetMovieInfoHandler == nil {
		api.MovieinfoGetMovieInfoHandler = movieinfo.GetMovieInfoHandlerFunc(func(params movieinfo.GetMovieInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation movieinfo.GetMovieInfo has not yet been implemented")
		})
	}
	if api.RatemoviesMoviesratingHandler == nil {
		api.RatemoviesMoviesratingHandler = ratemovies.MoviesratingHandlerFunc(func(params ratemovies.MoviesratingParams) middleware.Responder {
			return middleware.NotImplemented("operation ratemovies.Moviesrating has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
