package service

import (
	"MoviesDB/restapi/operations/add_comment"
	"reflect"
	"testing"

	"github.com/go-openapi/runtime/middleware"
)

func TestAddComment(t *testing.T) {
	type args struct {
		params add_comment.PostcommentsParams
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
			if got := AddComment(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
