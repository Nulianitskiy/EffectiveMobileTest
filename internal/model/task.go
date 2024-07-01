package model

import (
	"fmt"
	"time"
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
		return fmt.Errorf("Task is already started ")
	}
	t.TimeStart = time.Now()

	return nil
}

func (t *Task) EndTask() error {
	if !t.TimeEnd.IsZero() {
		return fmt.Errorf("Task is already ended ")
	}
	t.TimeEnd = time.Now()

	return nil
}

func (t *Task) Assign(peopleId int) {
	t.PeopleId = peopleId
}
