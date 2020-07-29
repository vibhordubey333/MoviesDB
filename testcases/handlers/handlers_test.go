package handlers

import (
	"MovieDB/service"
	"MoviesDB/restapi/operations/add_comment"
	"MoviesDB/restapi/operations/add_movie"
	"MoviesDB/restapi/operations/movieinfo"
	"reflect"
	"testing"

	"github.com/go-openapi/runtime/middleware"
)

func TestAddComments(t *testing.T) {
	type args struct {
		params add_comment.PostcommentsParams
	}
	tests := []struct {
		name string
		args args
		want middleware.Responder
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.AddComments(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddMovies(t *testing.T) {
	type args struct {
		params add_movie.PostmovieParams
	}
	tests := []struct {
		name string
		args args
		want middleware.Responder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.AddMovies(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMovies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMovieInfo(t *testing.T) {
	type args struct {
		params movieinfo.GetmoviesinfoParams
	}
	tests := []struct {
		name string
		args args
		want middleware.Responder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.GetMovieInfo(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMovieInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
