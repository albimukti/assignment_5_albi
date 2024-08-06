package main

import (
	"finance-planner/api"
	model "finance-planner/models" // Ensure this matches your actual models folder name
	"finance-planner/repository"
	"finance-planner/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=albi123 dbname=DB_tranning port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate models
	if err := db.AutoMigrate(&model.Category{}, &model.Record{}, &model.Transfer{}, &model.Wallet{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	// Initialize repositories
	categoryRepo := repository.NewCategoryRepository(db)
	recordRepo := repository.NewRecordRepository(db)
	transferRepo := repository.NewTransferRepository(db)
	walletRepo := repository.NewWalletRepository(db)

	// Initialize services
	categoryService := service.NewCategoryService(categoryRepo)
	recordService := service.NewRecordService(recordRepo)
	transferService := service.NewTransferService(transferRepo)
	walletService := service.NewWalletService(walletRepo)

	// Initialize handlers
	categoryHandler := api.NewCategoryHandler(categoryService)
	recordHandler := api.NewRecordHandler(recordService)
	transferHandler := api.NewTransferHandler(transferService)
	walletHandler := api.NewWalletHandler(walletService)

	router := gin.Default()

	// Category routes
	router.POST("/categories", categoryHandler.CreateCategory)
	router.GET("/categories/:id", categoryHandler.GetCategoryByID)
	router.GET("/categories", categoryHandler.ListCategories)

	// Wallet routes
	router.POST("/wallets", walletHandler.CreateWallet)
	// router.GET("/wallets/:ID", walletHandler.GetWalletByID)
	router.GET("/users/:userID/wallets", walletHandler.ListWalletsByUserID)

	// Record routes
	router.POST("/records", recordHandler.CreateRecord)
	// router.GET("/wallets/:ID/records", recordHandler.ListRecordsByWalletID) // Ensure this matches your handler method
	router.GET("/wallets/:walletID/records/time", recordHandler.ListRecordsByTime)

	// Transfer routes
	router.POST("/transfers", transferHandler.CreateTransfer)

	// Run the server
	router.Run(":8080")
}
