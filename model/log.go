package model

import "time"

// Control struct has the common time control members of main structs
type Control struct {
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Deleted   bool      `db:"deleted" json:"deleted"`
}
