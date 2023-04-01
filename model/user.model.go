package model

type User struct {
	Id        int    `json:"id"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}
