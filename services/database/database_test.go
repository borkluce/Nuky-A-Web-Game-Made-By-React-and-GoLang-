package mongotest

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestConnectToMongoDB(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal("Error on loading .env file")
	}

	connString := os.Getenv("CONNECTION_STRING")
	if connString == "" {
		t.Fatal("CONNECTION_STRING is not set in .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		t.Fatalf("Connection failed: %s", err.Error())
	}

	if err = client.Ping(ctx, nil); err != nil {
		t.Fatalf("Connection failed: %s", err.Error())
	}
}
