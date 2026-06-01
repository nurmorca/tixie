package repository

import (
	"context"
	"fmt"
	"ticket/data/domain"
	"ticket/data/dto"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/gommon/log"
)

type ITicketRepository interface {
	GetAllEvents() ([]domain.Event, error)
	CreateEvent(event domain.Event) error
	GetEventById(eventId int64) (*domain.Event, error)
	DeleteEvent(eventId int64) error
	UpdateDescription(eventId int64, description string) error
	GetTicketsForEvent(id int64, isAvailable bool) ([]domain.Ticket, error)
	UpdateTicketStatus(ctx context.Context, ticketIds []int64, newStatus string) error
	GetAllTicketsForEvent(eventId int64) ([]dto.TicketDTO, error)
	GetBookingTickets(userTickets dto.UserTicketsDTO) ([]dto.BookingTicketDTO, error)
}

type TicketRepository struct {
	Pool *pgxpool.Pool
}

func NewTicketRepository(pool *pgxpool.Pool) ITicketRepository {
	return &TicketRepository{Pool: pool}
}

func (ticketRepository *TicketRepository) GetTicketById(eventId int64) ([]dto.TicketDTO, error) {
	ctx := context.Background()
	var tickets []dto.TicketDTO
	query := `SELECT
	t.ti_id,
    s.se_seat_row,
    s.se_seat_number,
    s.se_seat_type,
    t.ti_status,
    t.ti_price,
	e.ev_host,
	e.ev_name,
	e.ev_venue,
	e.ev_description,
	e.ev_datetime
    FROM ticket t
	JOIN seat s ON t.ti_seat_id = s.se_id
	JOIN events e ON t.ti_event_id = e.ev_id
    WHERE t.ti_id = $1
	ORDER BY s.se_seat_row, s.se_seat_number`

	err := pgxscan.Select(ctx, ticketRepository.Pool, &tickets, query, eventId)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (ticketRepository *TicketRepository) GetBookingTickets(userTickets dto.UserTicketsDTO) ([]dto.BookingTicketDTO, error) {
	ctx := context.Background()
	var tickets []dto.BookingTicketDTO
	query := `SELECT
	t.ti_id,
    t.ti_status,
    t.ti_price,
	t.ti_event_id
    FROM ticket t
    WHERE t.ti_id = ANY($1)`

	err := pgxscan.Select(ctx, ticketRepository.Pool, &tickets, query, userTickets.TicketIds)
	if err != nil {
		return nil, err
	}
	fmt.Println(tickets)
	return tickets, nil
}

func (ticketRepository *TicketRepository) GetAllTicketsForEvent(eventId int64) ([]dto.TicketDTO, error) {
	ctx := context.Background()
	var tickets []dto.TicketDTO
	query := `SELECT
	t.ti_id,
    s.se_seat_row,
    s.se_seat_number,
    s.se_seat_type,
    t.ti_status,
    t.ti_price,
	e.ev_host,
	e.ev_name,
	e.ev_venue,
	e.ev_description,
	e.ev_datetime
    FROM ticket t
	JOIN seat s ON t.ti_seat_id = s.se_id
	JOIN events e ON t.ti_event_id = e.ev_id
    WHERE t.ti_event_id = $1
	ORDER BY s.se_seat_row, s.se_seat_number`

	err := pgxscan.Select(ctx, ticketRepository.Pool, &tickets, query, eventId)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (ticketRepository *TicketRepository) UpdateTicketStatus(ctx context.Context, ticketIds []int64, newStatus string) error {
	query := `UPDATE ticket SET ti_status=$1 WHERE ti_id = ANY($2)`
	result, err := ticketRepository.Pool.Exec(
		ctx,
		query,
		newStatus,
		ticketIds)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		log.Error("tickets not found. ", err)
	}
	return nil
}

func (ticketRepository *TicketRepository) GetTicketsForEvent(eventId int64, isAvailable bool) ([]domain.Ticket, error) {
	ctx := context.Background()
	var tickets []domain.Ticket
	query := `SELECT
	t.ti_id,
	t.ti_seat_id,
	t.ti_event_id,
    s.se_seat_row as "seat.se_seat_row",
    s.se_seat_number as "seat.se_seat_number",
    s.se_seat_type as "seat.se_seat_type",
    t.ti_status,
    t.ti_price
    FROM ticket t
    JOIN seat s ON t.ti_seat_id = s.se_id
    WHERE t.ti_event_id = $1`

	if isAvailable {
		query += ` AND t.ti_status = 'available'`
	}

	query += ` ORDER BY s.se_seat_row, s.se_seat_number`

	err := pgxscan.Select(ctx, ticketRepository.Pool, &tickets, query, eventId)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (ticketRepository *TicketRepository) GetAllEvents() ([]domain.Event, error) {
	ctx := context.Background()
	var events []domain.Event
	query := `SELECT * FROM EVENTS ORDER BY EV_DATETIME DESC`

	err := pgxscan.Select(ctx, ticketRepository.Pool, &events, query)
	if err != nil {
		fmt.Println("err", events, err)
		return nil, err
	}
	fmt.Println("hello", events, err)

	return events, nil
}

func (ticketRepository *TicketRepository) GetEventById(id int64) (*domain.Event, error) {
	ctx := context.Background()
	var event domain.Event
	query := `SELECT * FROM EVENTS WHERE EV_ID=$1`

	err := pgxscan.Get(ctx, ticketRepository.Pool, &event, query, id)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (ticketRepository *TicketRepository) DeleteEvent(id int64) error {
	ctx := context.Background()
	query := `DELETE * FROM EVENTS WHERE EV_ID=$1`

	result, err := ticketRepository.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		log.Error("event not found. ", err)
	}

	return nil
}

func (ticketRepository *TicketRepository) CreateEvent(event domain.Event) error {
	ctx := context.Background()
	query := `INSERT INTO EVENTS (EV_NAME, EV_HOST, EV_VENUE, EV_DATETIME,  EV_DESCRIPTION) VALUES ($1, $2, $3, $4, $5) RETURNING EV_ID, EV_CREATED_AT`

	err := ticketRepository.Pool.QueryRow(
		ctx,
		query,
		event.EvName,
		event.EvHost,
		event.EvVenue,
		event.EvDateTime,
		event.EvDescription).Scan(&event.EvID, &event.EvCreatedAt)
	if err != nil {
		return err
	}

	log.Info("User created: ", event)

	return nil
}

func (ticketRepository *TicketRepository) UpdateDescription(eventId int64, description string) error {
	ctx := context.Background()
	query := `UPDATE EVENTS SET EV_DESCRIPTION=$1 WHERE EV_ID=$2`

	result, err := ticketRepository.Pool.Exec(
		ctx,
		query,
		description,
		eventId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		log.Error("event not found. ", err)
	}

	return nil
}
