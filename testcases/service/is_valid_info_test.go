package service

import (
	"testing"
)

func TestUserNameIsValid(t *testing.T) {
	type args struct {
		uname string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserNameIsValid(tt.args.uname); got != tt.want {
				t.Errorf("UserNameIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMovieNameIsValid(t *testing.T) {
	type args struct {
		moviename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MovieNameIsValid(tt.args.moviename); got != tt.want {
				t.Errorf("MovieNameIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
