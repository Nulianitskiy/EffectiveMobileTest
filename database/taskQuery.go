package database

import (
	"GoTimeTracker/internal/model"
	"GoTimeTracker/pkg/logger"
	"go.uber.org/zap"
)

// AddTask Добавить задачу
func (d *Database) AddTask(t model.Task) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	query := `INSERT INTO task (name, description) VALUES ($1, $2) RETURNING id`
	var id int
	err := d.db.QueryRow(query, t.Name, t.Description).Scan(&id)
	if err != nil {
		logger.Error("Ошибка при добавлении задачи", zap.Error(err))
		return err
	}
	t.Id = id
	logger.Info("Задача успешно добавлена", zap.Int("id", t.Id))
	return nil
}

// AssignPeopleOnTask Назначить сотрудников на задачу
func (d *Database) AssignPeopleOnTask(t model.Task) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	query := `UPDATE task SET people_id = $2 WHERE id = $1`
	_, err := d.db.Exec(query, t.Id, t.PeopleId)
	if err != nil {
		logger.Error("Ошибка при назначении сотрудников на задачу", zap.Error(err), zap.Int("taskId", t.Id))
		return err
	}
	logger.Info("Сотрудники успешно назначены на задачу", zap.Int("taskId", t.Id))
	return nil
}

// StartTaskTime Начать отслеживание времени задачи
func (d *Database) StartTaskTime(t model.Task) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if err := t.StartTask(); err != nil {
		logger.Error("Ошибка при начале отслеживания времени задачи", zap.Error(err), zap.Int("taskId", t.Id))
		return err
	}

	query := `UPDATE task SET time_start = $2 WHERE id = $1`
	_, err := d.db.Exec(query, t.Id, t.TimeStart)
	if err != nil {
		logger.Error("Ошибка при обновлении времени начала задачи", zap.Error(err), zap.Int("taskId", t.Id))
		return err
	}
	logger.Info("Время начала задачи успешно обновлено", zap.Int("taskId", t.Id))
	return nil
}

// EndTaskTime Завершить отслеживание времени задачи
func (d *Database) EndTaskTime(t model.Task) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if err := t.EndTask(); err != nil {
		logger.Error("Ошибка при завершении отслеживания времени задачи", zap.Error(err), zap.Int("taskId", t.Id))
		return err
	}

	query := `UPDATE task SET time_end = $2 WHERE id = $1`
	_, err := d.db.Exec(query, t.Id, t.TimeEnd)
	if err != nil {
		logger.Error("Ошибка при обновлении времени завершения задачи", zap.Error(err), zap.Int("taskId", t.Id))
		return err
	}
	logger.Info("Время завершения задачи успешно обновлено", zap.Int("taskId", t.Id))
	return nil
}

// GetPeopleTasks Получить задачи для конкретного сотрудника
func (d *Database) GetPeopleTasks(p model.People) ([]model.Task, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	var tasks []model.Task

	query := "SELECT *, EXTRACT(epoch FROM (time_end - time_start)) AS duration FROM task WHERE people_id = $1 ORDER BY duration DESC "

	err := d.db.Select(&tasks, query, p.Id)
	if err != nil {
		logger.Error("Ошибка при получении задач для сотрудника", zap.Error(err), zap.Int("peopleId", p.Id))
		return nil, err
	}
	logger.Info("Получен список задач для сотрудника", zap.Int("peopleId", p.Id), zap.Int("tasksCount", len(tasks)))
	return tasks, nil
}
