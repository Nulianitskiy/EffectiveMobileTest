package main

import (
	dbase "GoTimeTracker/database"
	_ "GoTimeTracker/docs"
	"GoTimeTracker/internal/routes"
	"GoTimeTracker/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"os"
)

//	@title			Task Tracker
//	@version		1.0
//	@description	Тестовое задание для Effective Mobile.

//	@contact.name	Никита Ульяницкий
//	@contact.url	https://t.me/Nulianitskiy

// @host		localhost:8080
// @BasePath	/
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

	//router.LoadHTMLGlob("web/pages/*")
	//router.Static("/pages", "./web/pages")
	//router.Static("/js", "./web/js")
	//router.Static("/styles", "./web/styles")

	routes.SetupRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logger.Info("Запуск сервера на порту", zap.String("port", port))
	err = router.Run(":" + port)
	if err != nil {
		logger.Fatal("Ошибка запуска сервиса", zap.Error(err))
	}
}
