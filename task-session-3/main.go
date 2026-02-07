package main

import (
	"log"
	"net/http"

	"github.com/mhiqrambg/golang-dasar/task-session-3/internal/config"
	"github.com/mhiqrambg/golang-dasar/task-session-3/internal/handler"
	"github.com/mhiqrambg/golang-dasar/task-session-3/internal/repository"
	"github.com/mhiqrambg/golang-dasar/task-session-3/internal/service"
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

	// Product
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Transaction
	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// Report
	reportRepo := repository.NewReportRepository(db)
	reportService := service.NewReportService(reportRepo)
	reportHandler := handler.NewReportHandler(reportService)

	// Register Routes
	http.HandleFunc("/api/produk", productHandler.HandleProducts)       // GET (with ?name=), POST
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)   // GET, PUT, DELETE by ID
	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout) // POST
	http.HandleFunc("/api/report/hari-ini", reportHandler.HandleReportHariIni) // GET
	http.HandleFunc("/api/report", reportHandler.HandleReport)                 // GET ?start_date=&end_date=

	log.Printf("Server starting on port %s...\n", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
