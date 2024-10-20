package model

import "time"

type Todo struct {
	ID        int64     `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Completed bool      `db:"completed" json:"completed"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
