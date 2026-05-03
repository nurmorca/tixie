package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Ticket struct {
	TiID        int64           `db:"ti_id" json:"id"`
	TiEventID   int64           `db:"ti_event_id" json:"event_id"`
	TiSeatID    int64           `db:"ti_seat_id" json:"seat_id"`
	Status      string          `db:"ti_status" json:"seat_status"`
	Price       decimal.Decimal `db:"ti_price" json:"price"`
	TiCreatedAt time.Time       `db:"ti_created_at" json:"created_at"`
	TiSeat      Seat            `db:"seat" json:"seat,omitempty"`
}
