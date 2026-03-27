package domain

import "time"

type Seat struct {
	SeID         int64     `db:"se_id" json:"id"`
	SeVenueID    int64     `db:"se_venue_id" json:"venue_id"`
	SeSeatRow    string    `db:"se_seat_row" json:"seat_row"`
	SeSeatNumber int64     `db:"se_seat_number" json:"seat_number"`
	SeSeatType   string    `db:"se_seat_type" json:"seat_type"`
	SeCreatedAt  time.Time `db:"se_created_at" json:"created_at"`
}
