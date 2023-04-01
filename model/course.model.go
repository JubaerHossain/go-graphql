package model

type Course struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	User        int    `json:"user"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
}
