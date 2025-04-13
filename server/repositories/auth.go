package repositories

import (
	"avito-tech-spring-2025/dto"
	dto_entity "avito-tech-spring-2025/dto/entity"
	"avito-tech-spring-2025/repositories/prepared"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

func FetchLogin(email, hashedPassword string) (*dto_entity.User, *dto.Error) {
	row := prepared.LoginStmt.QueryRow(email, hashedPassword)

	var user dto_entity.User
	err := row.Scan(&user.ID, &user.Email, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &dto.Error{Message: "Неверные учетные данные"}
		}
		return nil, &dto.Error{Message: "Ошибка при авторизации"}
	}

	return &user, nil
}

func FetchRegister(email, hashedPassword, role string) (*dto_entity.User, *dto.Error) {
	row := prepared.RegisterCheck.QueryRow(email)
	var exists bool
	err := row.Scan(&exists)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, &dto.Error{Message: "Ошибка при проверке существующего пользователя"}
	}
	if exists {
		return nil, &dto.Error{Message: "Пользователь с таким email уже существует"}
	}

	id := uuid.New().String()
	_, err = prepared.RegisterStmt.Exec(id, email, hashedPassword, role)
	if err != nil {
		return nil, &dto.Error{Message: err.Error()}
	}

	return &dto_entity.User{
		ID:    id,
		Email: email,
		Role:  role,
	}, nil
}
