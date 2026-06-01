package dto

import "github.com/shopspring/decimal"

type BookingTicketDTO struct {
	TicketId  int64           `db:"ti_id" json:"ticketId"`
	Price     decimal.Decimal `db:"ti_price" json:"price"`
	Status    string          `db:"ti_status" json:"ticketStatus"`
	TiEventID int64           `db:"ti_event_id" json:"eventId"`
}
