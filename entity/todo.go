package entity

import (
	"database/sql"
	"time"
)

type Todo struct {
	Id         int
	Todo       string
	status     string
	created_at time.Time
	deleted_at sql.NullTime
}
