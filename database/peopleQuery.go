package database

import (
	"GoTimeTracker/internal/model"
	"GoTimeTracker/pkg/logger"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

// GetAllPeople возвращает список сотрудников из базы данных с фильтрами и пагинацией
func (d *Database) GetAllPeople(page int, pageSize int, filters map[string]interface{}) ([]model.People, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	offset := (page - 1) * pageSize

	var query strings.Builder
	query.WriteString("SELECT id, name, surname, patronymic, address FROM people")

	var args []interface{}
	var index int

	// Генерация условий WHERE для каждого фильтра
	if len(filters) > 0 {
		query.WriteString(" WHERE ")
		for key, value := range filters {
			if index > 0 {
				query.WriteString(" AND ")
			}
			query.WriteString(fmt.Sprintf("%s = $%d", key, index+1))
			args = append(args, value)
			index++
		}
	}

	query.WriteString(fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", index+1, index+2))
	args = append(args, pageSize, offset)

	var peoples []model.People
	err := d.db.Select(&peoples, query.String(), args...)
	if err != nil {
		logger.Error("Ошибка при получении списка сотрудников", zap.Error(err))
		return nil, err
	}
	logger.Info("Получен список сотрудников", zap.Int("count", len(peoples)))
	return peoples, nil
}

// AddPeople добавление сотрудника в формате
func (d *Database) AddPeople(p model.People) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	query := `INSERT INTO people (name, surname, patronymic, address) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := d.db.QueryRow(query, p.Name, p.Surname, p.Patronymic, p.Address).Scan(&id)
	if err != nil {
		logger.Error("Ошибка при добавлении сотрудника", zap.Error(err))
		return err
	}
	p.Id = id
	logger.Info("Сотрудник успешно добавлен", zap.Int("id", p.Id))
	return nil
}

// UpdatePeople обновление информации о сотруднике
func (d *Database) UpdatePeople(p model.People) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	query := `UPDATE people SET name = $2, surname = $3, patronymic = $4, address = $5 WHERE id = $1`
	_, err := d.db.Exec(query, p.Id, p.Name, p.Surname, p.Patronymic, p.Address)
	if err != nil {
		logger.Error("Ошибка при обновлении информации о сотруднике", zap.Error(err), zap.Int("id", p.Id))
		return err
	}
	logger.Info("Информация о сотруднике успешно обновлена", zap.Int("id", p.Id))
	return nil
}

// DeletePeople удаление информации о сотруднике
func (d *Database) DeletePeople(id int) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	query := `DELETE FROM people WHERE id = $1`
	_, err := d.db.Exec(query, id)
	if err != nil {
		logger.Error("Ошибка при удалении информации о сотруднике", zap.Error(err), zap.Int("id", id))
		return err
	}
	logger.Info("Информация о сотруднике успешно удалена", zap.Int("id", id))
	return nil
}
