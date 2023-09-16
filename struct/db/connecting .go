package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	j "yandex_go/struct/json"
	e  "yandex_go/struct/utils"

)

func InitializeDB() (*gorm.DB, error)  {
	// Подключение к базе данных SQLite
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	e.ConnectToDBError(err)
	// Автомиграция - создание таблицы (если она не существует)
	db.AutoMigrate(&j.Person{})

	return db, nil
}
func CreatePersonWithJob(db *gorm.DB, person j.Person, job j.Job) error {
	// Начало транзакции
	tx := db.Begin()

	// Создание записи в таблице jobs
	result := tx.Save(&job)
	if result.Error != nil {
		tx.Rollback() // Откат транзакции в случае ошибки
		return fmt.Errorf("error creating job: %w", result.Error)
	}

	// Создание записи в таблице people с ссылкой на запись в таблице jobs
	person.JobID = job.ID
	result = tx.Create(&person)
	if result.Error != nil {
		tx.Rollback() // Откат транзакции в случае ошибки
		return fmt.Errorf("error creating person: %w", result.Error)
	}

	tx.Commit()
// Фиксация транзакции

	return nil
}   