package modules

import "time"

// Task структура задач
type Task struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"`
}
