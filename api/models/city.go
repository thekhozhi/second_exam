package models

type City struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type CreateCity struct {
	Name string `json:"name"`
}

