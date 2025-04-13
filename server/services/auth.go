package services

import (
	"avito-tech-spring-2025/dto"
	dto_req "avito-tech-spring-2025/dto/requests"
	"avito-tech-spring-2025/middleware"
	"avito-tech-spring-2025/repositories"
)

func Login(req *dto_req.Login) (string, *dto.Error) {
	result, err := repositories.FetchLogin(req.Email, req.Password)
	if err != nil {
		return "", err
	}

	if result != nil {
		token, tokenErr := middleware.GenerateToken(result.Email, result.Role)
		if tokenErr != nil {
			return "", tokenErr
		}
		return token, nil
	}
	return "", nil
}

func Register(req *dto_req.Register) (string, *dto.Error) {
	result, err := repositories.FetchRegister(req.Email, req.Password, req.Role)
	if err != nil {
		return "", err
	}

	if result != nil {
		token, tokenErr := middleware.GenerateToken(result.Email, result.Role)
		if tokenErr != nil {
			return "", tokenErr
		}
		return token, nil
	}
	return "", nil
}
