package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type Database struct {
	db    *sqlx.DB
	mutex sync.Mutex
}

var (
	instance *Database
	once     sync.Once
)

func GetInstance() (*Database, error) {
	var err error
	once.Do(func() {
		instance, err = newDatabase()
		if err != nil {
			log.Fatal("Ошибка создания экземпляра Database:", err)
		}
	})
	return instance, err
}

func newDatabase() (*Database, error) {
	connectionString := "postgres://dbuser:stoic@db:5436/tasktrackerdb?sslmode=disable"
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	// Ping базы данных для проверки подключения
	err = db.Ping()
	if err != nil {
		log.Fatal("Ошибка проверки подключения к базе данных:", err)
	}
	log.Println("Подключение к базе данных PostgreSQL успешно")

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
