package controller

type UserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role_id  uuid.UUID `json:"role_id"`
}