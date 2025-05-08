package main

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

}

// Clients
const (
    mongoClient *mongo.Client
)

// Services
const (

)

// Repositories
const (

)

// Inits: --------------------------------------------------------------------

func init() {
    envType := os.Getenv("ENV")
    if envType != "prod" {
        
    }
}

func initClients() {

}
func initRepos() {

}

func initServices() {

}

// Setups: --------------------------------------------------------------------

func setupDBConnection() *mongo.Client {

}

func setupRoutes 