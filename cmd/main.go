package main

import (
	"dictionary/internal/service"
	"dictionary/pkg/logs"

	"github.com/labstack/echo/v4"
)

func main() {
	// создаем логгер
	logger := logs.NewLogger(false)

	// подключаемся к базе
	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}

	svc := service.NewService(db, logger)

	router := echo.New()
	// создаем группу api
	api := router.Group("api")

	// прописываем пути
	api.GET("/word/:id", svc.GetWordById)
	api.POST("/words", svc.CreateWords)
	api.PUT("/word/:id", svc.UpdateWord)
	api.DELETE("/word/:id", svc.DeleteWord)

	api.GET("/search/ru", svc.SearchWords)

	// пути для репортов
	api.POST("/report", svc.CreateReport)
	api.GET("/report/:id", svc.GetReportById)
	api.PUT("/report/:id", svc.UpdateReport)
	api.DELETE("/report/:id", svc.DeleteReport)

	// запускаем сервер, чтобы слушал 8000 порт
	router.Logger.Fatal(router.Start(":8000"))
}
