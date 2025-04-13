package handlers

import (
	dto "avito-tech-spring-2025/dto/requests"
	dto_req "avito-tech-spring-2025/dto/requests"
	"avito-tech-spring-2025/middleware"
	"avito-tech-spring-2025/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DummyLogin(c *gin.Context) {
	var userType dto.DummyLogin
	if err := c.ShouldBindJSON(&userType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный запрос"})
		return
	}

	token, err := middleware.GenerateToken("dummyuser", userType.Role)
	if err != nil {
		c.JSON(400, err.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(c *gin.Context) {
	var request dto_req.Login
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := services.Login(&request)
	if err != nil {
		c.JSON(400, err.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	var request dto_req.Register
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := services.Register(&request)
	if err != nil {
		c.JSON(400, err.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
