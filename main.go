package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/factory"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/infrastructure"
)

func main() {
	db := infrastructure.InitDB("./data/vouchers.db")
	defer db.Close()

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

	apiGroup := router.Group("/api")

	modules := factory.InitAllModule(db)

	for _, m := range modules {
		m.RestHandler().RegisterRoutes(apiGroup)
	}

	router.Run(":8080")
}
