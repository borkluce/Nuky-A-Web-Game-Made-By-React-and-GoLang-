package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	province_repo "services/internal/province/repo"
	province_service "services/internal/province/service"
)

var (
	mongoClient     *mongo.Client
	provinceService *province_service.ProvinceService
	provinceRepo    *province_repo.ProvinceRepo
)

func init() {
	// Load environment variables
	envType := os.Getenv("ENV")
	if envType != "prod" {
		if err := godotenv.Load("../../.env"); err != nil {
			log.Printf("Error loading .env: %v", err)
		}
	}

	initClients()
	initRepos()
	initServices()
}

func main() {
	defer mongoClient.Disconnect(context.TODO())

	// Setting up to work with UTC
	c := cron.New(cron.WithLocation(time.UTC))

	_, err := c.AddFunc("36 15 * * *", func() {
		log.Println("Daily Nuke Time!")

		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Execute the destroyment round
		roundCount, err := provinceService.ExecuteDestroymentRound(ctx)
		if err != nil {
			log.Printf("Nuke updating error: %v", err)
		} else {
			log.Printf("Nuke operation successful! Round count: %d", roundCount)
		}
	})

	if err != nil {
		log.Fatalf("Adding Cron job error: %v", err)
	}

	c.Start()
	log.Println("Nuke timer started. Daily nuke at 14:00 UTC")

	// Keep the program running
	select {}
}

func initClients() {
	setupDBConnection()
}

func initRepos() {
	db := mongoClient.Database("nuky_db")
	provinceRepo = province_repo.NewProvinceRepo(db.Collection("provinces"))
	log.Println("Repository initialized")
}

func initServices() {
	startDateStr := os.Getenv("GAME_START_DATE")
	if startDateStr == "" {
		log.Fatal("GAME_START_DATE environment variable is required")
	}

	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		log.Fatalf("Invalid GAME_START_DATE format: %v", err)
	}

	provinceService = province_service.NewProvinceService(provinceRepo, startDate)

	log.Println("Service initialized")
}

func setupDBConnection() {
	var err error
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DB")))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")
}
