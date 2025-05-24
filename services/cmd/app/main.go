package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	auth_repo "services/internal/auth/repo"
	auth_service "services/internal/auth/service"

	province_repo "services/internal/province/repo"
	province_service "services/internal/province/service"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kahlery/pkg/go/log/util"
)

// Clients
var (
	mongoClient *mongo.Client
)

// Services
var (
	authService     *auth_service.AuthService
	provinceService *province_service.ProvinceService
)

// Repos
var (
	userRepo     *auth_repo.UserRepo
	provinceRepo *province_repo.ProvinceRepo
)

// Main --------------------------------------------------------------------

func init() {
	envType := os.Getenv("ENV")
	if envType != "prod" {
		if err := godotenv.Load("../../.env"); err != nil {
			util.LogError("error loading .env"+err.Error(), "main.init()", "")
		}
	} else {
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

	dir, err := os.Getwd()
	if err != nil {
		util.LogError("failed to get working directory: "+err.Error(), "main.init()", "")
	} else {
		util.LogSuccess("working directory can be reached:", "main.init()", "")
		fmt.Println(dir)
	}
}

func main() {
	mux := http.NewServeMux()
	setupRoutes(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: withCORS(mux),
	}

	util.LogSuccess("💣 Server starting on port "+port, "main.main()", "")
	if err := server.ListenAndServe(); err != nil {
		util.LogError("Server failed to start: "+err.Error(), "main.main()", "")
		os.Exit(1)
	}
}

// Inits --------------------------------------------------------------------

func initClients() {
	setupDBConnection()
}

func initRepos() {
	db := mongoClient.Database("nuky_db")

	userRepo = auth_repo.NewUserRepo(db.Collection("users"))
	provinceRepo = province_repo.NewProvinceRepo(db.Collection("provinces"))

	util.LogSuccess("Repositories initialized", "main.initRepos()", "")
}

func initServices() {
	authService = auth_service.NewAuthService(userRepo)

	startDateStr := os.Getenv("GAME_START_DATE")
	if startDateStr == "" {
		log.Fatal("GAME_START_DATE environment variable is required")
	}

	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		log.Fatalf("Invalid GAME_START_DATE format: %v", err)
	}

	provinceService = province_service.NewProvinceService(provinceRepo, startDate)

	util.LogSuccess("Services initialized", "main.initServices()", "")
}

// Setups --------------------------------------------------------------------

func setupDBConnection() {
	var err error
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DB")))
	if err != nil {
		panic(err)
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		util.LogError("Failed to connect to MongoDB: "+err.Error(), "main.setupDBConnection()", "")
		panic(err)
	}

	util.LogSuccess("Connected to MongoDB", "main.setupDBConnection()", "")
}

func setupRoutes(mux *http.ServeMux) {
	// Public Auth routes
	mux.HandleFunc("/api/auth/register", authService.Register)
	mux.HandleFunc("/api/auth/login", authService.Login)

	// Province routes
	mux.HandleFunc("/api/province", provinceService.GetAllProvinces)
	mux.HandleFunc("/api/province/top", provinceService.GetTopProvinces)
	mux.HandleFunc("/api/province/attack", provinceService.AttackProvince)
	mux.HandleFunc("/api/province/support", provinceService.SupportProvince)
	mux.HandleFunc("/api/province/round", provinceService.GetCurrentRoundHandler)

	util.LogSuccess("Routes initialized", "main.setupRoutes()", "")
}

func withCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
