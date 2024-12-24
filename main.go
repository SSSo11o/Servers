package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"servers/logger"
)

func main() {
	var logger = logger.InitLumberjackLogger()
	router := gin.Default()
	router.GET("/Hi", func(c *gin.Context) {
		logger.Info("Получен запрос на /")
		c.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
	})
	if err := router.Run(":8080"); err != nil {
		logger.Error("Ошибка при запуске сервера:", err)
	}
}
