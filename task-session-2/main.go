package main

import (
	"log"
	"net/http"

	"github.com/mhiqrambg/golang-dasar/task-session-2/internal/config"
	"github.com/mhiqrambg/golang-dasar/task-session-2/internal/handler"
	"github.com/mhiqrambg/golang-dasar/task-session-2/internal/repository"
	"github.com/mhiqrambg/golang-dasar/task-session-2/internal/service"
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort   string `mapstructure:"SERVER_PORT"`
	DBConnection string `mapstructure:"DB_CONNECTION"`
}

func main() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found, using environment variables")
	}

	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("PORT", "8080")

	port := viper.GetString("PORT")
	if port == "" {
		port = viper.GetString("SERVER_PORT")
	}

	cfg := Config{
		ServerPort:   port,
		DBConnection: viper.GetString("DB_CONNECTION"),
	}

	// Validasi DB_CONNECTION
	if cfg.DBConnection == "" {
		log.Fatal("Error: DB_CONNECTION environment variable is required")
	}

	log.Printf("Server is configured to run on port: %s\n", cfg.ServerPort)

	// Setup database
	db, err := config.InitDB(cfg.DBConnection)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Initialize Repository, Service, and Handler
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Register Routes
	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)

	log.Printf("Server starting on port %s...\n", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

