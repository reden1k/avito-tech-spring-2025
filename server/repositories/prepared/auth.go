package prepared

import (
	"avito-tech-spring-2025/db"
	"database/sql"
	"log"
)

var LoginStmt, RegisterStmt, RegisterCheck *sql.Stmt

func PrepareAuthQueries() {
	var err error

	LoginStmt, err = db.DB.Prepare(`
		SELECT uuid, email, role FROM users WHERE email = $1 AND password = $2
	`)
	if err != nil {
		log.Fatalf("Ошибка при подготовке запроса: %v", err)
	}

	RegisterStmt, err = db.DB.Prepare(`
		INSERT INTO users (uuid, email, password, role) VALUES ($1, $2, $3, $4)
	`)
	if err != nil {
		log.Fatalf("Ошибка при подготовке запроса: %v", err)
	}

	RegisterCheck, err = db.DB.Prepare(`
		SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)
	`)
	if err != nil {
		log.Fatalf("Ошибка при подготовке запроса: %v", err)
	}
}
