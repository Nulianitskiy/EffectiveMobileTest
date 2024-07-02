package main

import (
	dbase "GoTimeTracker/database"
	"GoTimeTracker/internal/routes"
	"GoTimeTracker/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal("Ошибка загрузки переменных окружения")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := dbase.GetInstance()
	if err != nil {
		logger.Fatal("Ошибка подключения к базе данных", zap.Error(err))
	}
	defer db.Close()

	router := gin.Default()

	router.LoadHTMLGlob("web/pages/*")
	router.Static("/pages", "./web/pages")
	router.Static("/js", "./web/js")
	router.Static("/styles", "./web/styles")

	routes.SetupRoutes(router)

	logger.Info("Запуск сервера на порту", zap.String("port", port))
	err = router.Run(":" + port)
	if err != nil {
		logger.Fatal("Ошибка запуска сервиса", zap.Error(err))
	}
}
