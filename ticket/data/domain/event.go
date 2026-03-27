package domain

import "time"

type Event struct {
	EvID          int64     `db:"ev_id" json:"id,omitempty"`
	EvHost        string    `db:"ev_host" json:"hosted_by"`
	EvName        string    `db:"ev_name" json:"name"`
	EvVenue       string    `db:"ev_venue" json:"venue"`
	EvDescription string    `db:"ev_description" json:"description"`
	EvCreatedAt   time.Time `db:"ev_created_at" json:"created_At"`
	EvDateTime    time.Time `db:"ev_datetime" json:"date_time"`
}
