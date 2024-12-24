package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"servers/logger"
)

func main() {
	var logger = logger.InitLogger()

	router := gin.Default()
	router.GET("/Hi", func(c *gin.Context) {
		logger.Info("Получен запрос на /")
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Logger!"})
	})

	if err := router.Run(":8080"); err != nil {
		logger.Error("Ошибка при запуске сервера:", err)
	}
}
