package dto

import "github.com/shopspring/decimal"

type EventSeatDTO struct {
	SeatID     int64           `db:"es_id" json:"id"`
	SeatRow    string          `db:"se_seat_row" json:"row"`
	SeatNumber int             `db:"se_seat_number" json:"number"`
	SeatType   string          `db:"se_seat_type" json:"seat_type"`
	Status     string          `db:"es_status" json:"seat_status"`
	Price      decimal.Decimal `db:"es_price" json:"price"`
}
