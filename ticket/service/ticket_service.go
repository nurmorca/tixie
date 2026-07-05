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

const STATUS_AVAILABLE string = "AVAILABLE"
const STATUS_SOLD string = "SOLD"
const STATUS_RESERVED string = "RESERVED"

// TODO: separate validation logic
type ITicketService interface {
	GetAllEvents() ([]domain.Event, error)
	CreateEvent(event domain.Event) error
	GetEventById(eventId int64) (*domain.Event, error)
	DeleteEvent(eventId int64) error
	UpdateDescription(eventId int64, description string) error
	GetAvailableTicketsForEvent(eventId int64) ([]domain.Ticket, error)
	GetTicketsForEvent(eventId int64) ([]domain.Ticket, error)
	LockTickets(ctx context.Context, userBooking dto.UserTicketsDTO) error
	ReleaseTickets(ctx context.Context, userBooking dto.UserTicketsDTO) error
	GetAllTicketsForEvent(eventId int64) ([]dto.TicketDTO, error)
	CheckUserReservations(ctx context.Context, userBooking dto.UserTicketsDTO) bool
	UpdateTicketStatus(ctx context.Context, ticketIds []int64, status string) error
	GetBookingTickets(ctx context.Context, userTickets dto.UserTicketsDTO) ([]dto.BookingTicketDTO, error)
	InitiateTicketReservation(ctx context.Context, userTickets dto.UserTicketsDTO) ([]dto.BookingTicketDTO, error)
	CompleteReservation(ctx context.Context, userTickets dto.UserTicketsDTO) error
	FreeTickets(ctx context.Context, userTickets dto.UserTicketsDTO) error
}

type TicketService struct {
	ticketRepository repository.ITicketRepository
	redisClient      *redis.Client
}

func (ticketService *TicketService) FreeTickets(ctx context.Context, userTickets dto.UserTicketsDTO) error {
	err := ticketService.UpdateTicketStatus(ctx, userTickets.TicketIds, STATUS_AVAILABLE)
	if err != nil {
		return err
	}

	return nil
}

func NewTicketService(repository repository.ITicketRepository, redisClient *redis.Client) ITicketService {
	return &TicketService{
		ticketRepository: repository,
		redisClient:      redisClient,
	}
}

func (ticketService *TicketService) CompleteReservation(ctx context.Context, userTickets dto.UserTicketsDTO) error {
	bookingTickets, err := ticketService.GetBookingTickets(ctx, userTickets)
	if err != nil {
		return err
	}

	id := checkStatusOfTickets(bookingTickets, STATUS_AVAILABLE)
	if id != 0 {
		return errors.New(fmt.Sprintf("Ticket with ID %d not reserved for user!", id))
	}

	err = ticketService.UpdateTicketStatus(ctx, userTickets.TicketIds, STATUS_SOLD)
	if err != nil {
		return err
	}

	return nil
}

func (ticketService *TicketService) InitiateTicketReservation(ctx context.Context, userTickets dto.UserTicketsDTO) ([]dto.BookingTicketDTO, error) {
	if ticketService.CheckUserReservations(ctx, userTickets) == false {
		return nil, errors.New("Ticket(s) not assigned to user")
	}

	bookingTickets, err := ticketService.GetBookingTickets(ctx, userTickets)
	if err != nil {
		return nil, err
	}

	id := checkStatusOfTickets(bookingTickets, STATUS_RESERVED)
	if id != 0 {
		return nil, errors.New(fmt.Sprintf("Ticket with ID %d already reserved!", id))
	}

	err = ticketService.UpdateTicketStatus(ctx, userTickets.TicketIds, STATUS_RESERVED)
	if err != nil {
		return nil, err
	}

	return bookingTickets, nil
}

func (ticketService *TicketService) GetBookingTickets(ctx context.Context, userTickets dto.UserTicketsDTO) ([]dto.BookingTicketDTO, error) {
	bookingTickets, err := ticketService.ticketRepository.GetBookingTickets(userTickets)
	if err != nil {
		return nil, err
	}
	return bookingTickets, nil
}

func (ticketService *TicketService) GetAllTicketsForEvent(eventId int64) ([]dto.TicketDTO, error) {
	if eventId == 0 {
		return nil, errors.New("event id is not valid")
	}
	return ticketService.ticketRepository.GetAllTicketsForEvent(eventId)
}

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

func (ticketService *TicketService) LockTickets(ctx context.Context, userTicket dto.UserTicketsDTO) error {
	if *userTicket.UserID == 0 {
		return errors.New("user_id value is not entered")
	}
	for _, id := range userTicket.TicketIds {
		key := fmt.Sprintf("lock:%d", id)
		ok, _ := ticketService.redisClient.SetNX(ctx, key, userTicket.UserID, 10*time.Minute).Result()
		if !ok {
			return errors.New("seat already locked")
		}
	}
	return nil
}

func (ticketService *TicketService) UpdateTicketStatus(ctx context.Context, ticketIds []int64, status string) error {
	if ticketIds == nil {
		return errors.New("id values are not entered")
	}
	ticketService.ticketRepository.UpdateTicketStatus(ctx, ticketIds, status)

	return nil
}

func (ticketService *TicketService) ReleaseTickets(ctx context.Context, userTicket dto.UserTicketsDTO) error {
	for _, id := range userTicket.TicketIds {
		key := fmt.Sprintf("lock:%d", id)
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

func isTicketBookingRequestValid(userId int64, ticketId int64) bool {
	if userId == 0 || ticketId == 0 {
		return false
	}
	return true
}

func checkStatusOfTickets(tickets []dto.BookingTicketDTO, status string) int64 {
	for _, bookingTicket := range tickets {
		if bookingTicket.Status == status {
			return bookingTicket.TicketId
		}
	}

	return 0
}

func (ticketService *TicketService) CheckUserReservations(ctx context.Context, userTicket dto.UserTicketsDTO) bool {
	for _, id := range userTicket.TicketIds {
		key := fmt.Sprintf("lock:%d", id)
		value, _ := ticketService.redisClient.Get(ctx, key).Int64()
		if value != *userTicket.UserID {
			return false
		}
	}
	return true
}
