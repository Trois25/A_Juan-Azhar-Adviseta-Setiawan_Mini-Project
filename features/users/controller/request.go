package controller

type UserRequest struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Date_of_birth string `json:"date_of_birth"`
	RoleId        uint64 `json:"role_id"`
}
