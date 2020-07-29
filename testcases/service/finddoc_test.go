package service

import (
	"reflect"
	"testing"
)

func TestReadDocument(t *testing.T) {
	type args struct {
		filterobj    interface{}
		outputObject interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDocument(tt.args.filterobj, tt.args.outputObject)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateDocument(t *testing.T) {
	type args struct {
		filterObject interface{}
		operation    string
		update       map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateDocument(tt.args.filterObject, tt.args.operation, tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}
