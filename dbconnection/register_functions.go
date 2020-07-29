package dbconnection

import "go.mongodb.org/mongo-driver/mongo"

type DBConnectionInterface interface {
	GetMongoObject() *mongo.Client
}
