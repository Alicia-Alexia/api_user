package models


type Users struct {
	ID				int 	 `json:"id"`
	Username        string   `json:"username"`
	Firstname       string   `json:"firstname"`
	Lastname        string   `json:"lastname"`
	Email           string   `json:"email"`
	Password        string   `json:"password"`
	Phone           string   `json:"phone"`
	Userstatus      int      `json:"userStatus"`
}