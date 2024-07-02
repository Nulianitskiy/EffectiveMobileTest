package controller

import (
	"GoTimeTracker/database"
	"GoTimeTracker/internal/model"
	"GoTimeTracker/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

type FilterParams struct {
	Page     int                    `json:"page" binding:"required"`
	PageSize int                    `json:"page_size" binding:"required"`
	Filters  map[string]interface{} `json:"filters"`
}

func GetAllPeople(ctx *gin.Context) {
	var filterParams FilterParams
	if err := ctx.ShouldBindJSON(&filterParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	people, err := db.GetAllPeople(filterParams.Page, filterParams.PageSize, filterParams.Filters)
	if err != nil {
		logger.Error("Ошибка при получении списка сотрудников", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, people)
	logger.Info("Успешно получен список сотрудников")
}

func AddPeople(ctx *gin.Context) {
	passport := strings.Split(ctx.Query("passportNumber"), " ")
	serie, _ := strconv.Atoi(passport[0])
	number, _ := strconv.Atoi(passport[1])

	people := model.People{
		PassportSerie:  serie,
		PassportNumber: number,
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = db.AddPeople(people)
	if err != nil {
		logger.Error("Ошибка при добавлении сотрудника", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
	logger.Info("Сотрудник успешно добавлен")
}

func UpdatePeople(ctx *gin.Context) {
	var p model.People
	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = db.UpdatePeople(p)
	if err != nil {
		logger.Error("Ошибка при обновлении информации о сотруднике", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
	logger.Info("Информация о сотруднике успешно обновлена")
}

func DeletePeople(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = db.DeletePeople(id)
	if err != nil {
		logger.Error("Ошибка при удалении информации о сотруднике", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
	logger.Info("Информация о сотруднике успешно удалена")
}
