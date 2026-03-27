package domain

import "time"

type Venue struct {
	VeID        int64     `db:"ve_id" json:"id"`
	VeName      string    `db:"ve_name" json:"name"`
	VeCapacity  int64     `db:"ve_capacity" json:"capacity"`
	VeCreatedAt time.Time `db:"ve_created_at" json:"created_at"`
}
