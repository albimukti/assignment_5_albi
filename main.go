package main

import (
	"finance-planner/api"
	model "finance-planner/models"
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

	db.AutoMigrate(&model.Category{}, &model.Record{}, &model.Transfer{}, &model.Wallet{})

	categoryRepo := repository.NewCategoryRepository(db)
	recordRepo := repository.NewRecordRepository(db)
	transferRepo := repository.NewTransferRepository(db)
	walletRepo := repository.NewWalletRepository(db)

	categoryService := service.NewCategoryService(categoryRepo)
	recordService := service.NewRecordService(recordRepo)
	transferService := service.NewTransferService(transferRepo)
	walletService := service.NewWalletService(walletRepo)

	categoryHandler := api.NewCategoryHandler(categoryService)
	recordHandler := api.NewRecordHandler(recordService)
	transferHandler := api.NewTransferHandler(transferService)
	walletHandler := api.NewWalletHandler(walletService)

	router := gin.Default()

	router.POST("/categories", categoryHandler.CreateCategory)
	router.GET("/categories/:id", categoryHandler.GetCategoryByID)
	router.GET("/categories", categoryHandler.ListCategories)

	router.POST("/wallets", walletHandler.CreateWallet)
	router.GET("/wallets/:id", walletHandler.GetWalletByID)
	router.GET("/users/:userID/wallets", walletHandler.ListWalletsByUserID)

	router.POST("/records", recordHandler.CreateRecord)
	router.GET("/wallets/:walletID/records", recordHandler.ListRecordsByWalletID)
	router.GET("/wallets/:walletID/records/time", recordHandler.ListRecordsByTime)

	router.POST("/transfers", transferHandler.CreateTransfer)

	router.Run(":8080")
}
