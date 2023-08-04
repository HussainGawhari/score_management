package models

// Player represents the player attributes
type Player struct {
	ID      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	Country string `json:"country,omitempty"`
	Score   int    `json:"score,omitempty"`
}

type Response struct {
	Status  int64    `json:"status"`
	Message string   `json:"message"`
	Data    []Player `json:"data"`
}
