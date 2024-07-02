package database

import "GoTimeTracker/internal/model"

// AddTask Добавить задачу
func (d *Database) AddTask(t model.Task) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	query := `INSERT INTO task (name, description) VALUES ($1, $2) RETURNING id`
	var id int
	err := d.db.QueryRow(query, t.Name, t.Description).Scan(&id)
	if err != nil {
		return err
	}
	t.Id = id
	return nil
}

func (d *Database) AssignPeopleOnTask(t model.Task) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	query := `UPDATE task SET people_id = $2 WHERE id = $1`
	_, err := d.db.Exec(query, t.Id, t.PeopleId)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) StartTaskTime(t model.Task) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if err := t.StartTask(); err != nil {
		return err
	}

	query := `UPDATE task SET time_start = $2 WHERE id = $1`
	_, err := d.db.Exec(query, t.Id, t.TimeStart)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) EndTaskTime(t model.Task) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if err := t.EndTask(); err != nil {
		return err
	}

	query := `UPDATE task SET time_end = $2 WHERE id = $1`
	_, err := d.db.Exec(query, t.Id, t.TimeEnd)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) GetPeopleTasks(p model.People) ([]model.Task, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	var tasks []model.Task

	query := "SELECT *, EXTRACT(epoch FROM (time_end - time_start)) AS duration FROM task WHERE id = $1 ORDER BY duration DESC "

	err := d.db.Select(tasks, query, p.Id)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
