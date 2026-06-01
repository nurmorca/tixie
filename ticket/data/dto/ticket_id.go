package dto

type UserTicketsDTO struct {
	UserID    *int64  `json:"userId,omitempty"`
	TicketIds []int64 `json:"ticketIds"`
}
