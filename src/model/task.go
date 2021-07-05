package model

import "time"

type Task struct {
	ID int `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	updatedAt time.Time `json:"updated_at"`
	Title string `json:"title"`
}
