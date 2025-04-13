package dto_entity

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"` // client / moderator / pvz_employee
}
