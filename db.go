package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection(context context.Context) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := "mongodb+srv://r9gz0adw:isek92yfZipKffwr@main.cpcsrgk.mongodb.net/?retryWrites=true&w=majority"
	fmt.Println(uri)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context, opts)
	if err != nil {
		panic(err)
	}
	return client, nil
}
