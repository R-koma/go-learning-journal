package todo

import "time"

type Item struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}
