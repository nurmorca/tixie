package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type UserTicketDTO struct {
	TicketID      int64           `db:"ti_id" json:"id"`
	SeatID        int64           `db:"es_id" json:"seat_id"`
	UserID        int64           `db:"ti_user_id" json:"user_id"`
	SeatRow       string          `db:"se_seat_row" json:"row"`
	SeatNumber    int             `db:"se_seat_number" json:"number"`
	SeatType      string          `db:"se_seat_type" json:"seat_type"`
	Status        string          `db:"es_status" json:"seat_status"`
	Price         decimal.Decimal `db:"es_price" json:"price"`
	EvHost        string          `db:"ev_host" json:"hosted_by"`
	EvName        string          `db:"ev_name" json:"name"`
	EvVenue       string          `db:"ev_venue" json:"venue"`
	EvDescription string          `db:"ev_description" json:"description"`
	EvDateTime    time.Time       `db:"ev_datetime" json:"date_time"`
}
