package model

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

type Task struct {
	Id          int
	PeopleId    int
	Name        string
	Description string
	TimeStart   time.Time
	TimeEnd     time.Time
}

func (t *Task) StartTask() error {
	if !t.TimeStart.IsZero() {
		return fmt.Errorf("Задача уже начата")
	}
	t.TimeStart = time.Now()

	// Логирование начала выполнения задачи
	logger := zap.L()
	logger.Info("Задача начата",
		zap.Int("taskId", t.Id),
		zap.String("taskName", t.Name),
		zap.Time("timeStart", t.TimeStart),
	)

	return nil
}

func (t *Task) EndTask() error {
	if !t.TimeEnd.IsZero() {
		return fmt.Errorf("Задача уже завершена")
	}
	t.TimeEnd = time.Now()

	// Логирование завершения выполнения задачи
	logger := zap.L()
	logger.Info("Задача завершена",
		zap.Int("taskId", t.Id),
		zap.String("taskName", t.Name),
		zap.Time("timeEnd", t.TimeEnd),
	)

	return nil
}

func (t *Task) Assign(peopleId int) {
	t.PeopleId = peopleId

	// Логирование назначения человека на задачу
	logger := zap.L()
	logger.Info("Человек назначен на задачу",
		zap.Int("taskId", t.Id),
		zap.String("taskName", t.Name),
		zap.Int("peopleId", t.PeopleId),
	)
}
