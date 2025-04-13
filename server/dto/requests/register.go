package dto_req

// DTO для запроса регистрации пользователя
type Register struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=client moderator"`
}
