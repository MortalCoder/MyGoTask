package words

import "time"

type Word struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Translation string `json:"translation"`
}

type Report struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
