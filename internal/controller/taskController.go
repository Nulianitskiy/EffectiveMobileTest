package controller

import (
	"GoTimeTracker/database"
	"GoTimeTracker/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// AddTask godoc
//
//	@Summary		Добавить задачу
//	@Description	Добавляет новую задачу
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//
//	@Param			name		query	string	true	"Название задачи"	example(Новая задача)
//	@Param			description	query	string	true	"Описание задачи"	example(Описание...)
//
//	@Success		200
//	@Failure		400	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/task [post]
func AddTask(ctx *gin.Context) {
	name := ctx.Query("name")
	description := ctx.Query("description")

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	err = db.AddTask(name, description)
	if err != nil {
		logger.Error("Ошибка при добавлении задачи", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
	logger.Info("Задача успешно добавлена")
}

// AssignPeopleOnTask godoc
//
//	@Summary		Назначить сотрудников на задачу
//	@Description	Назначает сотрудников на указанную задачу
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//
//	@Param			id			query	int	true	"Идентификатор задачи"		example(0)
//	@Param			people_id	query	int	true	"Идентификатор работника"	example(0)
//
//	@Success		200
//	@Failure		400	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/taskAssign [put]
func AssignPeopleOnTask(ctx *gin.Context) {
	id := ctx.Query("id")
	peopleId := ctx.Query("people_id")

	idValue, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("Ошибка при парсинге id", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	peopleIdValue, err := strconv.Atoi(peopleId)
	if err != nil {
		logger.Error("Ошибка при парсинге people_id", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	err = db.AssignPeopleOnTask(idValue, peopleIdValue)
	if err != nil {
		logger.Error("Ошибка при назначении сотрудников на задачу", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
	logger.Info("Сотрудники успешно назначены на задачу")
}

// StartTask godoc
//
//	@Summary		Начать задачу
//	@Description	Начинает отслеживание времени задачи
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//
//	@Param			id	query	int	true	"Идентификатор задачи"	example(0)
//
//	@Success		200
//	@Failure		400	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/taskStart [put]
func StartTask(ctx *gin.Context) {
	id := ctx.Query("id")

	idValue, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("Ошибка при парсинге id", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	err = db.StartTaskTime(idValue)
	if err != nil {
		logger.Error("Ошибка при начале отслеживания времени задачи", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
	logger.Info("Время начала задачи успешно обновлено")
}

// EndTask godoc
//
//	@Summary		Завершить задачу
//	@Description	Завершает отслеживание времени задачи
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//
//	@Param			id	query	int	true	"Идентификатор задачи"	example(0)
//
//	@Success		200
//	@Failure		400	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/taskEnd [put]
func EndTask(ctx *gin.Context) {
	id := ctx.Query("id")

	idValue, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("Ошибка при парсинге id", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	err = db.EndTaskTime(idValue)
	if err != nil {
		logger.Error("Ошибка при завершении отслеживания времени задачи", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
	logger.Info("Время завершения задачи успешно обновлено")
}

// GetTasks godoc
//
//	@Summary		Получить задачи сотрудника
//	@Description	Возвращает список задач для указанного сотрудника
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//
//	@Param			people_id	query		int	true	"Идентификатор работника"	example(0)
//
//	@Success		200			{array}		model.Task
//	@Failure		400			{object}	ErrorResponse
//	@Failure		500			{object}	ErrorResponse
//	@Router			/task [get]
func GetTasks(ctx *gin.Context) {
	peopleId := ctx.Query("people_id")

	peopleIdValue, err := strconv.Atoi(peopleId)
	if err != nil {
		logger.Error("Ошибка при парсинге people_id", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	tasks, err := db.GetPeopleTasks(peopleIdValue)
	if err != nil {
		logger.Error("Ошибка при получении задач для сотрудника", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
	logger.Info("Успешно получен список задач для сотрудника")
}
