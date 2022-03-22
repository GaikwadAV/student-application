package utility

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
// connection with mongodb
func Connection() (*mongo.Client, context.Context) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)

	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	return client, ctx
}
