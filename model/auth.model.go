package model

type Auth struct {
	Token         string    `json:"token"`
	RefreshToken  string `json:"refresh_token"`
}

