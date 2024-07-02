package controller

import (
	"GoTimeTracker/database"
	"GoTimeTracker/internal/model"
	"GoTimeTracker/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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
//	@Param			task	body	model.Task	true	"Информация о задаче"
//	@Success		200
//	@Failure		400	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/task [post]
func AddTask(ctx *gin.Context) {
	var t model.Task
	if err := ctx.ShouldBindJSON(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	err = db.AddTask(t)
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
//	@Param			task	body	model.Task	true	"Информация о задаче"
//	@Success		200
//	@Failure		400	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/taskAssign [patch]
func AssignPeopleOnTask(ctx *gin.Context) {
	var t model.Task
	if err := ctx.ShouldBindJSON(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	err = db.AssignPeopleOnTask(t)
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
//	@Param			task	body	model.Task	true	"Информация о задаче"
//	@Success		200
//	@Failure		400	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/taskStart [patch]
func StartTask(ctx *gin.Context) {
	var t model.Task
	if err := ctx.ShouldBindJSON(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	err = db.StartTaskTime(t)
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
//	@Param			task	body	model.Task	true	"Информация о задаче"
//	@Success		200
//	@Failure		400	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/taskEnd [patch]
func EndTask(ctx *gin.Context) {
	var t model.Task
	if err := ctx.ShouldBindJSON(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	db, err := database.GetInstance()
	if err != nil {
		logger.Error("Ошибка получения экземпляра базы данных", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	err = db.EndTaskTime(t)
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
//	@Param			people	body		model.People	true	"Информация о сотруднике"
//	@Success		200		{array}		model.Task
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/task [get]
func GetTasks(ctx *gin.Context) {
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

	tasks, err := db.GetPeopleTasks(p)
	if err != nil {
		logger.Error("Ошибка при получении задач для сотрудника", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
	logger.Info("Успешно получен список задач для сотрудника")
}
