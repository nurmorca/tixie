package dto

type TicketReservationDTO struct {
	SeatID  int64  `json:"seat_id"`
	EventID int64  `json:"event_id"`
	UserID  *int64 `json:"user_id,omitempty"`
}
