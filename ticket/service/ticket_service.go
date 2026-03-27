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
	GetAvailableSeatsForEvent(eventId int64) ([]dto.EventSeatDTO, error)
	GetSeatsForEvent(eventId int64) ([]dto.EventSeatDTO, error)
	LockSeat(ctx context.Context, lockSeatRequest dto.TicketReservationDTO) error
	ReleaseSeat(ctx context.Context, releaseSeatRequest dto.TicketReservationDTO) error
	GetAllTicketsForEvent(eventId int64) ([]dto.UserTicketDTO, error)
	CreateTicket(ticketReservation dto.TicketReservationDTO) error
	GetAllTicketsForUser(userId int64) ([]dto.UserTicketDTO, error)
	GetUserTicketsForEvent(eventId int64, userId int64) ([]dto.UserTicketDTO, error)
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

func (ticketService *TicketService) GetUserTicketsForEvent(eventId int64, userId int64) ([]dto.UserTicketDTO, error) {
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
}

func (ticketService *TicketService) GetAllTicketsForEvent(eventId int64) ([]dto.UserTicketDTO, error) {
	if eventId == 0 {
		return nil, errors.New("event id is not valid")
	}
	return ticketService.ticketRepository.GetAllTicketsForEvent(eventId)
}

func (ticketService *TicketService) CreateTicket(ticketReservation dto.TicketReservationDTO) error {
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
}

func (ticketService *TicketService) GetSeatsForEvent(eventId int64) ([]dto.EventSeatDTO, error) {
	if eventId == 0 {
		return nil, errors.New("event id is not valid")
	}
	return ticketService.ticketRepository.GetSeatsForEvent(eventId, false)
}

func (ticketService *TicketService) GetAvailableSeatsForEvent(eventId int64) ([]dto.EventSeatDTO, error) {
	if eventId == 0 {
		return nil, errors.New("event id is not valid")
	}
	return ticketService.ticketRepository.GetSeatsForEvent(eventId, true)
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

func (ticketService *TicketService) LockSeat(ctx context.Context, lockSeatRequest dto.TicketReservationDTO) error {
	if lockSeatRequest.UserID == nil {
		return errors.New("user_id value is not entered")
	}
	key := fmt.Sprintf("lock:%d:%d", lockSeatRequest.EventID, lockSeatRequest.SeatID)
	ok, _ := ticketService.redisClient.SetNX(ctx, key, *lockSeatRequest.UserID, 10*time.Minute).Result()
	if !ok {
		return errors.New("seat already locked")
	}
	err := ticketService.ticketRepository.UpdateSeatStatus(ctx, lockSeatRequest.SeatID, "locked")
	return err
}

func (ticketService *TicketService) ReleaseSeat(ctx context.Context, releaseSeatRequest dto.TicketReservationDTO) error {
	err := ticketService.releaseSeat(ctx, releaseSeatRequest)
	if err != nil {
		return err
	}
	err = ticketService.ticketRepository.UpdateSeatStatus(ctx, releaseSeatRequest.SeatID, "available")
	return err
}

func (ticketService *TicketService) releaseSeat(ctx context.Context, releaseSeatRequest dto.TicketReservationDTO) error {
	key := fmt.Sprintf("lock:%d:%d", releaseSeatRequest.EventID, releaseSeatRequest.SeatID)
	return ticketService.redisClient.Del(ctx, key).Err()
}

func isEventRequestValid(event domain.Event) bool {
	if event.EvName == "" || event.EvHost == "" || event.EvVenue == "" || event.EvDateTime.IsZero() {
		return false
	}
	return true
}

func isTicketReservationRequestValid(ticketRes dto.TicketReservationDTO) bool {
	if ticketRes.SeatID == 0 || ticketRes.EventID == 0 || *ticketRes.UserID == 0 {
		return false
	}
	return true
}

func (ticketService *TicketService) isTicketReservedForUser(ctx context.Context, ticketRes dto.TicketReservationDTO) bool {
	key := fmt.Sprintf("lock:%d:%d", ticketRes.EventID, ticketRes.SeatID)
	value, _ := ticketService.redisClient.Get(ctx, key).Int64()

	if value == *ticketRes.UserID {
		return true
	}

	return false
}
