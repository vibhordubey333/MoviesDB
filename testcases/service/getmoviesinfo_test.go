package service

import (
	"MoviesDB/restapi/operations/movieinfo"
	"reflect"
	"testing"

	"github.com/go-openapi/runtime/middleware"
)

func TestGetMoviesInfo(t *testing.T) {
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
			if got := GetMoviesInfo(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMoviesInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
