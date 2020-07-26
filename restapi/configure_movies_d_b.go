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

	if api.AddMovieAddmovieratingHandler == nil {
		api.AddMovieAddmovieratingHandler = add_movie.AddmovieratingHandlerFunc(func(params add_movie.AddmovieratingParams) middleware.Responder {
			return middleware.NotImplemented("operation add_movie.Addmovierating has not yet been implemented")
		})
	}
	if api.MovieinfoGetmoviesinfoHandler == nil {
		api.MovieinfoGetmoviesinfoHandler = movieinfo.GetmoviesinfoHandlerFunc(func(params movieinfo.GetmoviesinfoParams) middleware.Responder {
			return middleware.NotImplemented("operation movieinfo.Getmoviesinfo has not yet been implemented")
		})
	}
	if api.AddCommentPostcommentsHandler == nil {
		api.AddCommentPostcommentsHandler = add_comment.PostcommentsHandlerFunc(func(params add_comment.PostcommentsParams) middleware.Responder {
			return middleware.NotImplemented("operation add_comment.Postcomments has not yet been implemented")
		})
	}
	if api.AddMoviePostmovieHandler == nil {
		api.AddMoviePostmovieHandler = add_movie.PostmovieHandlerFunc(func(params add_movie.PostmovieParams) middleware.Responder {
			return middleware.NotImplemented("operation add_movie.Postmovie has not yet been implemented")
		})
	}
	if api.RatemoviesRatemoviesHandler == nil {
		api.RatemoviesRatemoviesHandler = ratemovies.RatemoviesHandlerFunc(func(params ratemovies.RatemoviesParams) middleware.Responder {
			return middleware.NotImplemented("operation ratemovies.Ratemovies has not yet been implemented")
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
