package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

var (
	OrderBy = []string{"id","email","created_at","updated_at"}
)
