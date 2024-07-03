package database

import (
	"GoTimeTracker/internal/model"
	"GoTimeTracker/pkg/logger"
	"go.uber.org/zap"
	"time"
)

// AddTask Добавить задачу
func (d *Database) AddTask(name, description string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	query := `INSERT INTO task (name, description) VALUES ($1, $2) RETURNING id`
	var id int
	err := d.db.QueryRow(query, name, description).Scan(&id)
	if err != nil {
		logger.Error("Ошибка при добавлении задачи", zap.Error(err))
		return err
	}
	logger.Info("Задача успешно добавлена", zap.Int("id", id))
	return nil
}

// AssignPeopleOnTask Назначить сотрудников на задачу
func (d *Database) AssignPeopleOnTask(id, peopleId int) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	query := `UPDATE task SET people_id = $2 WHERE id = $1`
	_, err := d.db.Exec(query, id, peopleId)
	if err != nil {
		logger.Error("Ошибка при назначении сотрудников на задачу", zap.Error(err), zap.Int("taskId", id))
		return err
	}
	logger.Info("Сотрудники успешно назначены на задачу", zap.Int("taskId", id))
	return nil
}

// StartTaskTime Начать отслеживание времени задачи
func (d *Database) StartTaskTime(id int) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	query := `UPDATE task SET time_start = $2 WHERE id = $1`
	_, err := d.db.Exec(query, id, time.Now())
	if err != nil {
		logger.Error("Ошибка при обновлении времени начала задачи", zap.Error(err), zap.Int("taskId", id))
		return err
	}
	logger.Info("Время начала задачи успешно обновлено", zap.Int("taskId", id))
	return nil
}

// EndTaskTime Завершить отслеживание времени задачи
func (d *Database) EndTaskTime(id int) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	query := `UPDATE task SET time_end = $2 WHERE id = $1`
	_, err := d.db.Exec(query, id, time.Now())
	if err != nil {
		logger.Error("Ошибка при обновлении времени завершения задачи", zap.Error(err), zap.Int("taskId", id))
		return err
	}
	logger.Info("Время завершения задачи успешно обновлено", zap.Int("taskId", id))
	return nil
}

// GetPeopleTasks Получить задачи для конкретного сотрудника
func (d *Database) GetPeopleTasks(peopleId int) ([]model.Task, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	var tasks []model.Task

	query := "SELECT *, TO_CHAR(time_end - time_start, 'HH24:MI:SS') AS duration  FROM task WHERE people_id = $1 ORDER BY duration DESC "

	err := d.db.Select(&tasks, query, peopleId)
	if err != nil {
		logger.Error("Ошибка при получении задач для сотрудника", zap.Error(err), zap.Int("peopleId", peopleId))
		return nil, err
	}
	logger.Info("Получен список задач для сотрудника", zap.Int("peopleId", peopleId), zap.Int("tasksCount", len(tasks)))
	return tasks, nil
}
