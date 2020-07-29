package service

import (
	"MoviesDB/restapi/operations/add_movie"
	"reflect"
	"testing"

	"github.com/go-openapi/runtime/middleware"
)

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
			if got := AddMovies(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMovies() = %v, want %v", got, tt.want)
			}
		})
	}
}
