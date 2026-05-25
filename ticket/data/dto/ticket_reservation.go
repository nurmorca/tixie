package dto

type TicketReservationDTO struct {
	TicketID int64 `json:"ticket_id"`
	EventID  int64 `json:"event_id"`
}
