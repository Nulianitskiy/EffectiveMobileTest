package controller

import (
	"GoTimeTracker/database"
	"GoTimeTracker/internal/model"
	"GoTimeTracker/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// GetAllPeople godoc
//
//	@Summary		Получить всех сотрудников
//	@Description	Возвращает список всех сотрудников с возможностью фильтрации
//	@Tags			people
//	@Accept			json
//	@Produce		json
//
//	@Param			page		query		int		true	"Страница"													example(0)
//	@Param			page_size	query		int		true	"Количество объектов на странице"							example(5)
//	@Param			filter		query		string	false	"Фильтр (название параметра и параметр через двоеточие)"	example(name:Иванов)
//
//	@Success		200			{array}		model.People
//	@Failure		400			{object}	ErrorResponse
//	@Failure		500			{object}	ErrorResponse
//	@Router			/allPeople [get]
func GetAllPeople(ctx *gin.Context) {
	page := ctx.Query("page")
	pageSize := ctx.Query("page_size")

	pageValue, err := strconv.Atoi(page)
	if err != nil {
		logger.Error("Ошибка при парсинге значении страницы", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	pageSizeValue, err := strconv.Atoi(pageSize)
	if err != nil {
		logger.Error("Ошибка при парсинге количества страниц", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	filter := ctx.Query("filter")
	params := []string{"", ""}
	if len(filter) != 0 {
		params = strings.Split(filter, ":")
	}

	people, err := db.GetAllPeople(pageValue, pageSizeValue, params[0], params[1])
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
//	@Tags			people
//	@Accept			json
//	@Produce		json
//	@Param			passportNumber	query	string	true	"Номер паспорта (серия и номер через пробел)"	example(1234 567890)
//	@Success		200
//	@Failure		400	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/people [post]
func AddPeople(ctx *gin.Context) {
	passportParam := ctx.Query("passportNumber")
	passport, err := url.QueryUnescape(passportParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат номера паспорта"})
		return
	}

	passportParts := strings.Split(passport, " ")
	if len(passportParts) != 2 {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат номера паспорта"})
		return
	}

	serie, err := strconv.Atoi(passportParts[0])
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат серии паспорта"})
		return
	}

	number, err := strconv.Atoi(passportParts[1])
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат номера паспорта"})
		return
	}

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
//	@Tags			people
//	@Accept			json
//	@Produce		json
//	@Param			people	body	model.People	true	"Информация о сотруднике (серия и номер не изменяются)"
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
//	@Tags			people
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
