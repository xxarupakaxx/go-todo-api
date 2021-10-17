package model

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	LimitDate string `json:"limit_date"`
	Status    bool   `json:"status"`
}
