package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := "mongodb+srv://vafaill_master:jqSqKsjGHSa8hlK1@cluster0.uuzqws6.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer client.Disconnect(ctx)

	// Ping to check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Failed to ping MongoDB Atlas:", err)
		return
	}

	fmt.Println("Successfully connected to MongoDB Atlas!")
}
