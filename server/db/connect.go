package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "admin"
	password = "1234"
	dbname   = "db"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error

	for i := 0; i < 5; i++ {
		DB, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Printf("Ошибка подключения к базе данных: %v", err)
			time.Sleep(2 * time.Second) // Задержка перед повторной попыткой
			continue
		}

		err = DB.Ping()
		if err != nil {
			log.Printf("Не удалось установить соединение с базой данных: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		log.Println("Успешное подключение к базе данных")
		return DB, nil
	}

	return nil, fmt.Errorf("не удалось подключиться к базе данных после нескольких попыток: %v", err)
}
