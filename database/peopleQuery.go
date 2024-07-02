package database

import (
	"GoTimeTracker/internal/model"
	"fmt"
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
		return nil, err
	}
	return peoples, nil
}

// AddPeople добавление сотрудника в формате
func (d *Database) AddPeople(p model.People) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	query := `INSERT INTO people (passport_serie, passport_number) VALUES ($1, $2) RETURNING id`
	var id int
	err := d.db.QueryRow(query, p.PassportSerie, p.PassportNumber).Scan(&id)
	if err != nil {
		return err
	}
	p.Id = id
	return nil
}

// UpdatePeople обновление информации о сотруднике
func (d *Database) UpdatePeople(p model.People) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	query := `UPDATE people SET name = $2, surname = $3, patronymic = $4, address = $5 WHERE id = $1`
	_, err := d.db.Exec(query, p.Id, p.Name, p.Surname, p.Patronymic, p.Address)
	if err != nil {
		return err
	}
	return nil
}

// DeletePeople удаление информации о сотруднике
func (d *Database) DeletePeople(id int) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	query := `DELETE FROM people WHERE id = $1`
	_, err := d.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
