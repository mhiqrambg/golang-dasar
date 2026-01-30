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
	// Initialize viper
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Error reading config file, %s", err)
	}

	cfg := Config{
		ServerPort:   viper.GetString("SERVER_PORT"),
		DBConnection: viper.GetString("DB_CONNECTION"),
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

