package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufekoanggoro/voucher-seat-be/internal/infrastucture"
	"github.com/yusufekoanggoro/voucher-seat-be/internal/modules/voucher/delivery/resthandler"
	"github.com/yusufekoanggoro/voucher-seat-be/internal/modules/voucher/repository"
	"github.com/yusufekoanggoro/voucher-seat-be/internal/modules/voucher/usecase"
)

func main() {
	db := infrastucture.InitDB("./data/vouchers.db")
	defer db.Close()

	voucherRepo := repository.NewVoucherRepository(db)
	voucherUsecase := usecase.NewVoucherUsecase(voucherRepo)
	voucherHandler := resthandler.NewRestHandler(voucherUsecase)

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/check", voucherHandler.CheckFlight)
		api.POST("/generate", voucherHandler.GenerateVoucher)
	}

	router.Run(":8080")
}
