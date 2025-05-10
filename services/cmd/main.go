package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"services/internal/auth/repo"
	"services/internal/province/service"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kahleryasla/pkg/go/log/util"
)

// Clients
var (
	mongoClient *mongo.Client
)

// Services
var (
	authService     *service.AuthService
	provinceService *service.ProvinceService
)

// Repos
var (
	userRepo     *repo.UserRepo
	provinceRepo *repo.ProvinceRepo
)

func init() {
	envType := os.Getenv("ENV")
	if envType != "prod" {
		if err := godotenv.Load("../.env"); err != nil { // If environment is dev
			util.LogError("error loading .env"+err.Error(), "main.init()", "")
		}
	} else { // If the environment is prod
		util.LogSuccess(
			"Environment variables:"+
				"\n"+
				os.Getenv("DB")+
				"\n"+
				os.Getenv("PORT")+
				"\n"+
				os.Getenv("S3_BUCKET_NAME"),
			"main.init()",
			"",
		)
	}

	initClients()
	initRepos()
	initServices()

	// Check if the program can reach the working directory.
	dir, err := os.Getwd()
	if err != nil {
		util.LogError("failed to get working directory: "+err.Error(), "main.init()", "")
	} else {
		util.LogSuccess("working directory can be reached:", "main.init()", "")
		fmt.Println(dir)
	}
}

func main() {

}

// Inits: --------------------------------------------------------------------

func initClients() {
	setupDBConnection()
}

func initRepos() {

}

func initServices() {

}

// Setups: --------------------------------------------------------------------

func setupDBConnection() {
	var err error
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DB")))
	if err != nil {
		panic(err)
	}
}

func setupRoutes(mux *http.ServeMux) {

}

func setupMiddlewares() {}
