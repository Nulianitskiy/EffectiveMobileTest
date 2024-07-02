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

// GetAllPeople godoc
//
//	@Summary		Получить всех сотрудников
//	@Description	Возвращает список всех сотрудников с возможностью фильтрации
//	@Tags			employees
//	@Accept			json
//	@Produce		json
//	@Param			filterParams	body		FilterParams	true	"Параметры фильтрации"
//	@Success		200				{array}		model.People
//	@Failure		400				{object}	ErrorResponse
//	@Failure		500				{object}	ErrorResponse
//	@Router			/allPeople [get]
func GetAllPeople(ctx *gin.Context) {
	var filterParams FilterParams
	if err := ctx.ShouldBindJSON(&filterParams); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	people, err := db.GetAllPeople(filterParams.Page, filterParams.PageSize, filterParams.Filters)
	if err != nil {
		logger.Error("Ошибка при получении списка сотрудников", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, people)
	logger.Info("Успешно получен список сотрудников")
}

// AddPeople godoc
//
//	@Summary		Добавить сотрудника
//	@Description	Добавляет нового сотрудника по номеру паспорта
//	@Tags			employees
//	@Accept			json
//	@Produce		json
//	@Param			passportNumber	query	string	true	"Номер паспорта (серия и номер через пробел)"
//	@Success		200
//	@Failure		400	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/people [post]
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
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	err = db.AddPeople(people)
	if err != nil {
		logger.Error("Ошибка при добавлении сотрудника", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
	logger.Info("Сотрудник успешно добавлен")
}

// UpdatePeople godoc
//
//	@Summary		Обновить информацию о сотруднике
//	@Description	Обновляет информацию о сотруднике
//	@Tags			employees
//	@Accept			json
//	@Produce		json
//	@Param			people	body	model.People	true	"Информация о сотруднике"
//	@Success		200
//	@Failure		400	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/people [put]
func UpdatePeople(ctx *gin.Context) {
	var p model.People
	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	err = db.UpdatePeople(p)
	if err != nil {
		logger.Error("Ошибка при обновлении информации о сотруднике", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
	logger.Info("Информация о сотруднике успешно обновлена")
}

// DeletePeople godoc
//
//	@Summary		Удалить сотрудника
//	@Description	Удаляет сотрудника по идентификатору
//	@Tags			employees
//	@Accept			json
//	@Produce		json
//	@Param			id	query	int	true	"Идентификатор сотрудника"
//	@Success		200
//	@Failure		400	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/people [delete]
func DeletePeople(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	err = db.DeletePeople(id)
	if err != nil {
		logger.Error("Ошибка при удалении информации о сотруднике", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
	logger.Info("Информация о сотруднике успешно удалена")
}
