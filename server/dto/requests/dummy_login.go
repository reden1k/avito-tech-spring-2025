package dto_req

type DummyLogin struct {
	Role string `json:"role" binding:"required,oneof=client moderator"`
}
