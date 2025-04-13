package main

import (
	"avito-tech-spring-2025/db"
	"avito-tech-spring-2025/handlers"
	"avito-tech-spring-2025/middleware"
	"avito-tech-spring-2025/repositories/prepared"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	prepared.PrepareAuthQueries()

	r := gin.Default()

	r.POST("/dummyLogin", handlers.DummyLogin)
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.POST("/pvz", middleware.JWTMiddleware("moderator"), handlers.CreatePVZ)

	r.GET("/pvz", middleware.JWTMiddleware("moderator", "pvz_employee"), handlers.GetPVZ)

	pvz := middleware.JWTMiddleware("pvz_employee")
	r.POST("/receptions", pvz, handlers.Receptions)
	r.POST("/products", pvz, handlers.Products)
	r.POST("/pvz/:id/close_last_reception", pvz, handlers.CloseLastReception)
	r.POST("/pvz/:id/delete_last_product", pvz, handlers.DeleteLastProduct)

	r.Run(":8080")
}
