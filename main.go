package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/infrastucture"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/modules/voucher/delivery/resthandler"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/modules/voucher/repository"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/modules/voucher/usecase"
)

func main() {
	db := infrastucture.InitDB("./data/vouchers.db")
	defer db.Close()

	voucherRepo := repository.NewVoucherRepository(db)
	voucherUsecase := usecase.NewVoucherUsecase(voucherRepo)
	voucherHandler := resthandler.NewRestHandler(voucherUsecase)

	router := gin.Default()

	// add config CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		api.POST("/check", voucherHandler.CheckFlight)
		api.POST("/generate", voucherHandler.GenerateVoucher)
	}

	router.Run(":8080")
}
