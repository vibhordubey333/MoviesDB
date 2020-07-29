package dbconnection

import (
	connect "MoviesDB/dbconnection"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestGetMongoObject(t *testing.T) {
	tests := []struct {
		name string
		want *mongo.Client
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect.GetMongoObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMongoObject() = %v, want %v", got, tt.want)
			}
		})
	}
}
