package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type TicketDTO struct {
	TicketID      int64           `db:"ti_id" json:"id"`
	SeatRow       string          `db:"se_seat_row" json:"row"`
	SeatNumber    int             `db:"se_seat_number" json:"number"`
	SeatType      string          `db:"se_seat_type" json:"seat_type"`
	Status        string          `db:"ti_status" json:"seat_status"`
	Price         decimal.Decimal `db:"ti_price" json:"price"`
	EvHost        string          `db:"ev_host" json:"hosted_by"`
	EvName        string          `db:"ev_name" json:"name"`
	EvVenue       string          `db:"ev_venue" json:"venue"`
	EvDescription string          `db:"ev_description" json:"description"`
	EvDateTime    time.Time       `db:"ev_datetime" json:"date_time"`
}
