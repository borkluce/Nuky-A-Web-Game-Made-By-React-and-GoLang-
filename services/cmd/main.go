package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	auth_repo "services/internal/auth/repo"
	auth_service "services/internal/auth/service"
	province_repo "services/internal/province/repo"
	province_service "services/internal/province/service"

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
	authService     *auth_service.AuthService
	provinceService *province_service.ProvinceService
)

// Repos
var (
	userRepo     *auth_repo.UserRepo
	provinceRepo *province_repo.ProvinceRepo
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
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Setup routes
	setupRoutes(mux)

	// Apply middlewares
	finalHandler := setupMiddlewares()(mux)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: finalHandler,
	}

	util.LogSuccess("Server starting on port "+port, "main.main()", "")
	if err := server.ListenAndServe(); err != nil {
		util.LogError("Server failed to start: "+err.Error(), "main.main()", "")
		os.Exit(1)
	}
}

// Inits: --------------------------------------------------------------------

func initClients() {
	setupDBConnection()
}

func initRepos() {
	// Initialize repositories with database client
	db := mongoClient.Database("services_db")

	// User repository
	userRepo = auth_repo.NewUserRepo(db.Collection("users"))

	// Province repository
	provinceRepo = province_repo.NewProvinceRepo(db.Collection("provinces"))

	util.LogSuccess("Repositories initialized", "main.initRepos()", "")
}

func initServices() {
	// Initialize services with their dependencies
	authService = auth_service.NewAuthService(userRepo)
	provinceService = province_service.NewProvinceService(provinceRepo)

	util.LogSuccess("Services initialized", "main.initServices()", "")
}

// Setups: --------------------------------------------------------------------

func setupDBConnection() {
	var err error
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("CONNECTION_STRING")))
	if err != nil {
		panic(err)
	}

	// Ping the database to verify connection
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		util.LogError("Failed to connect to MongoDB: "+err.Error(), "main.setupDBConnection()", "")
		panic(err)
	}

	util.LogSuccess("Connected to MongoDB", "main.setupDBConnection()", "")
}

func setupRoutes(mux *http.ServeMux) {
	// Auth routes
	mux.HandleFunc("/api/auth/register", authService.RegisterHandler)
	mux.HandleFunc("/api/auth/login", authService.LoginHandler)
	mux.HandleFunc("/api/auth/verify", authService.VerifyHandler)

	// Province routes
	mux.HandleFunc("/api/provinces", provinceService.GetAllProvinces)
	//	mux.HandleFunc("/api/provinces/", provinceService.GetProvinceByID)

	util.LogSuccess("Routes initialized", "main.setupRoutes()", "")
}

func setupMiddlewares() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set common headers
			w.Header().Set("Content-Type", "application/json")

			// CORS middleware
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			// OPTIONS requests handling
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Logging middleware
			util.LogSuccess(fmt.Sprintf("%s %s", r.Method, r.URL.Path), "request", "")

			// Pass to the next handler
			next.ServeHTTP(w, r)
		})
	}
}
