package domain

import "time"

type Ticket struct {
	TiID          int64      `db:"TI_ID" json:"id"`
	TiEventID     string     `db:"TI_EVENT_ID" json:"event_id"`
	TiSeatNumber  string     `db:"TI_SEAT_NUMBER" json:"seat_number"`
	TiSection     string     `db:"TI_SECTION" json:"section"`
	TiPrice       float64    `db:"TI_PRICE" json:"price"`
	TiStatus      string     `db:"TI_STATUS" json:"status"` // "available", "locked", "sold"
	TiLockedBy    *string    `db:"TI_LOCKED_BY" json:"locked_by,omitempty"`
	TiLockedUntil *time.Time `db:"TI_LOCKED_UNTIL" json:"locked_until,omitempty"`
	TiCreatedAt   time.Time  `db:"TI_CREATED_AT" json:"created_at"`
}
