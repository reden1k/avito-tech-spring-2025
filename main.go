package main

import (
	"avito-tech-spring-2025/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/dummyLogin", handlers.DummyLogin)
	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	r.POST("/pvz", handlers.CreatePVZ)
	r.GET("/pvz", handlers.GetPVZ)
	r.POST("/pvz/:id/close_last_reception", handlers.CloseLastReception)
	r.POST("/pvz/:id/delete_last_product", handlers.DeleteLastProduct)

	r.POST("/receptions", handlers.Receptions)
	r.POST("/products", handlers.Products)

	r.Run(":8080") // Здесь мы указываем порт
}
