package model

import "github.com/lib/pq"

// Control struct has the common time control members of main structs
type Control struct {
	CreatedAt pq.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt pq.NullTime `db:"updated_at" json:"updated_at"`
	DeletedAt pq.NullTime `db:"deleted_at" json:"deleted_at"`
}
