package service

import (
	"context"
	"errors"
	"fmt"
	"ticket/data/domain"
	"ticket/data/dto"
	"ticket/repository"
	"time"

	"github.com/redis/go-redis/v9"
)

// TODO: separate validation logic
type ITicketService interface {
	GetAllEvents() ([]domain.Event, error)
	CreateEvent(event domain.Event) error
	GetEventById(eventId int64) (*domain.Event, error)
	DeleteEvent(eventId int64) error
	UpdateDescription(eventId int64, description string) error
	GetAvailableTicketsForEvent(eventId int64) ([]domain.Ticket, error)
	GetTicketsForEvent(eventId int64) ([]domain.Ticket, error)
	LockTickets(ctx context.Context, userBooking dto.UserBookingDTO) error
	ReleaseTickets(ctx context.Context, userBooking dto.UserBookingDTO) error
	GetAllTicketsForEvent(eventId int64) ([]dto.TicketDTO, error)
	CheckUserReservations(ctx context.Context, userBooking dto.UserBookingDTO) bool
	CompleteReservation(ctx context.Context, ticketIds []int64) error
	// CreateTicket(ticketReservation dto.TicketReservationDTO) error
	// GetAllTicketsForUser(userId int64) ([]dto.UserTicketDTO, error)
	// GetUserTicketsForEvent(eventId int64, userId int64) ([]dto.UserTicketDTO, error)
}

type TicketService struct {
	ticketRepository repository.ITicketRepository
	redisClient      *redis.Client
}

func NewTicketService(repository repository.ITicketRepository, redisClient *redis.Client) ITicketService {
	return &TicketService{
		ticketRepository: repository,
		redisClient:      redisClient,
	}
}

/* func (ticketService *TicketService) GetUserTicketsForEvent(eventId int64, userId int64) ([]dto.UserTicketDTO, error) {
	if userId == 0 {
		return nil, errors.New("user id is not valid")
	}
	if eventId == 0 {
		return nil, errors.New("event id is not valid")
	}
	return ticketService.ticketRepository.GetUserTicketsForEvent(eventId, userId)
}

func (ticketService *TicketService) GetAllTicketsForUser(userId int64) ([]dto.UserTicketDTO, error) {
	if userId == 0 {
		return nil, errors.New("user id is not valid")
	}
	return ticketService.ticketRepository.GetAllTicketsForUser(userId)
} */

func (ticketService *TicketService) GetAllTicketsForEvent(eventId int64) ([]dto.TicketDTO, error) {
	if eventId == 0 {
		return nil, errors.New("event id is not valid")
	}
	return ticketService.ticketRepository.GetAllTicketsForEvent(eventId)
}

/* func (ticketService *TicketService) CreateTicket(ticketReservation dto.TicketReservationDTO) error {
	ctx := context.Background()
	if !isTicketReservationRequestValid(ticketReservation) {
		return errors.New("please make sure all event id, seat id and user id correctly entered")
	}

	if !ticketService.isTicketReservedForUser(ctx, ticketReservation) {
		return errors.New("ticket is not reserved for user")
	}

	ticket := domain.Ticket{
		TiEventID: ticketReservation.EventID,
		TiSeatID:  ticketReservation.SeatID,
		TiUserID:  *ticketReservation.UserID,
	}

	err := ticketService.ticketRepository.CreateTicket(ticket)
	if err != nil {
		return err
	}

	err = ticketService.ticketRepository.UpdateSeatStatus(ctx, ticket.TiSeatID, "sold")
	if err != nil {
		return err
	}

	return ticketService.releaseSeat(ctx, ticketReservation)
} */

func (ticketService *TicketService) GetTicketsForEvent(eventId int64) ([]domain.Ticket, error) {
	if eventId == 0 {
		return nil, errors.New("event id is not valid")
	}
	return ticketService.ticketRepository.GetTicketsForEvent(eventId, false)
}

func (ticketService *TicketService) GetAvailableTicketsForEvent(eventId int64) ([]domain.Ticket, error) {
	if eventId == 0 {
		return nil, errors.New("event id is not valid")
	}
	return ticketService.ticketRepository.GetTicketsForEvent(eventId, true)
}

func (ticketService *TicketService) CreateEvent(event domain.Event) error {
	if !isEventRequestValid(event) {
		return errors.New("event is not valid. please enter all fields for event")
	}
	return ticketService.ticketRepository.CreateEvent(event)
}

func (ticketService *TicketService) GetAllEvents() ([]domain.Event, error) {
	return ticketService.ticketRepository.GetAllEvents()
}

func (ticketService *TicketService) GetEventById(eventId int64) (*domain.Event, error) {
	if eventId == 0 {
		return nil, errors.New("event id is not valid")
	}
	return ticketService.ticketRepository.GetEventById(eventId)
}

func (ticketService *TicketService) DeleteEvent(eventId int64) error {
	if eventId == 0 {
		return errors.New("event id is not valid")
	}
	return ticketService.ticketRepository.DeleteEvent(eventId)
}

func (ticketService *TicketService) UpdateDescription(eventId int64, description string) error {
	if eventId == 0 {
		return errors.New("event id is not valid")
	}
	if description == "" {
		return errors.New("description is not valid")
	}
	return ticketService.ticketRepository.UpdateDescription(eventId, description)
}

func (ticketService *TicketService) LockTickets(ctx context.Context, userBooking dto.UserBookingDTO) error {
	if *userBooking.UserID == 0 {
		return errors.New("user_id value is not entered")
	}
	for _, item := range userBooking.Items {
		key := fmt.Sprintf("lock:%d:%d", item.EventID, item.TicketID)
		ok, _ := ticketService.redisClient.SetNX(ctx, key, userBooking.UserID, 10*time.Minute).Result()
		if !ok {
			return errors.New("seat already locked")
		}
	}
	return nil
}

func (ticketService *TicketService) CompleteReservation(ctx context.Context, ticketIds []int64) error {
	if ticketIds == nil {
		return errors.New("id values are not entered")
	}
	ticketService.ticketRepository.UpdateTicketStatus(ctx, ticketIds, "sold")
	return nil
}

func (ticketService *TicketService) ReleaseTickets(ctx context.Context, userBooking dto.UserBookingDTO) error {
	for _, item := range userBooking.Items {
		key := fmt.Sprintf("lock:%d:%d", item.EventID, item.TicketID)
		err := ticketService.redisClient.Del(ctx, key).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func isEventRequestValid(event domain.Event) bool {
	if event.EvName == "" || event.EvHost == "" || event.EvVenue == "" || event.EvDateTime.IsZero() {
		return false
	}
	return true
}

func isUserBookingRequestValid(userBooking dto.UserBookingDTO) bool {
	if *userBooking.UserID == 0 || userBooking.Items == nil {
		return false
	}
	return true
}

func (ticketService *TicketService) CheckUserReservations(ctx context.Context, userBooking dto.UserBookingDTO) bool {
	for _, item := range userBooking.Items {
		key := fmt.Sprintf("lock:%d:%d", item.EventID, item.TicketID)
		value, _ := ticketService.redisClient.Get(ctx, key).Int64()
		if value != *userBooking.UserID {
			return false
		}
	}
	return true
}
