package domain

import "time"

type Ticket struct {
	TiID        int64     `db:"ti_id" json:"id"`
	TiEventID   int64     `db:"ti_event_id" json:"event_id"`
	TiSeatID    int64     `db:"ti_seat_id" json:"seat_id"`
	TiUserID    int64     `db:"ti_user_id" json:"user_id"`
	TiCreatedAt time.Time `db:"ti_created_at" json:"created_at"`
}
