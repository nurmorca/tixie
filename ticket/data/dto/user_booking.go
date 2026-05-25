package dto

type UserBookingDTO struct {
	UserID *int64                 `json:"user_id,omitempty"`
	Items  []TicketReservationDTO `json:"items"`
}
