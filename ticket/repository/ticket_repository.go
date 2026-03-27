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
	GetSeatsForEvent(id int64, isAvailable bool) ([]dto.EventSeatDTO, error)
	UpdateSeatStatus(ctx context.Context, seatId int64, newStatus string) error
	CreateTicket(ticket domain.Ticket) error
	GetAllTicketsForEvent(eventId int64) ([]dto.UserTicketDTO, error)
	GetAllTicketsForUser(userId int64) ([]dto.UserTicketDTO, error)
	GetUserTicketsForEvent(eventId int64, userId int64) ([]dto.UserTicketDTO, error)
}

type TicketRepository struct {
	Pool *pgxpool.Pool
}

func NewTicketRepository(pool *pgxpool.Pool) ITicketRepository {
	return &TicketRepository{Pool: pool}
}

func (ticketRepository *TicketRepository) GetUserTicketsForEvent(eventId int64, userId int64) ([]dto.UserTicketDTO, error) {
	ctx := context.Background()
	var tickets []dto.UserTicketDTO
	query := `SELECT
	t.ti_id,
	es.es_id,
	t.ti_user_id,
    s.se_seat_row,
    s.se_seat_number,
    s.se_seat_type,
    es.es_status,
    es.es_price,
	e.ev_host,
	e.ev_name,
	e.ev_venue,
	e.ev_description,
	e.ev_datetime
    FROM ticket t
    JOIN event_seat es ON t.ti_seat_id = es.es_id
	JOIN seat s ON es.es_seat_id = s.se_id
	JOIN events e ON t.ti_event_id = e.ev_id
    WHERE t.ti_event_id = $1 AND t.ti_user_id = $2
	ORDER BY s.se_seat_row, s.se_seat_number`

	err := pgxscan.Select(ctx, ticketRepository.Pool, &tickets, query, eventId, userId)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (ticketRepository *TicketRepository) GetAllTicketsForUser(userId int64) ([]dto.UserTicketDTO, error) {
	ctx := context.Background()
	var tickets []dto.UserTicketDTO
	query := `SELECT
	t.ti_id,
	es.es_id,
	t.ti_user_id,
    s.se_seat_row,
    s.se_seat_number,
    s.se_seat_type,
    es.es_status,
    es.es_price,
	e.ev_host,
	e.ev_name,
	e.ev_venue,
	e.ev_description,
	e.ev_datetime
    FROM ticket t
    JOIN event_seat es ON t.ti_seat_id = es.es_id
	JOIN seat s ON es.es_seat_id = s.se_id
	JOIN events e ON t.ti_event_id = e.ev_id
    WHERE t.ti_user_id = $1
	ORDER BY s.se_seat_row, s.se_seat_number`

	err := pgxscan.Select(ctx, ticketRepository.Pool, &tickets, query, userId)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (ticketRepository *TicketRepository) GetAllTicketsForEvent(eventId int64) ([]dto.UserTicketDTO, error) {
	ctx := context.Background()
	var tickets []dto.UserTicketDTO
	query := `SELECT
	t.ti_id,
	es.es_id,
	t.ti_user_id,
    s.se_seat_row,
    s.se_seat_number,
    s.se_seat_type,
    es.es_status,
    es.es_price,
	e.ev_host,
	e.ev_name,
	e.ev_venue,
	e.ev_description,
	e.ev_datetime
    FROM ticket t
    JOIN event_seat es ON t.ti_seat_id = es.es_id
	JOIN seat s ON es.es_seat_id = s.se_id
	JOIN events e ON t.ti_event_id = e.ev_id
    WHERE t.ti_event_id = $1
	ORDER BY s.se_seat_row, s.se_seat_number`

	err := pgxscan.Select(ctx, ticketRepository.Pool, &tickets, query, eventId)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (ticketRepository *TicketRepository) UpdateSeatStatus(ctx context.Context, seatId int64, newStatus string) error {
	query := `UPDATE event_seat SET es_status=$1 WHERE ES_ID=$2`
	result, err := ticketRepository.Pool.Exec(
		ctx,
		query,
		newStatus,
		seatId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		log.Error("event seat not found. ", err)
	}

	return nil
}

func (ticketRepository *TicketRepository) GetSeatsForEvent(eventId int64, isAvailable bool) ([]dto.EventSeatDTO, error) {
	ctx := context.Background()
	var seats []dto.EventSeatDTO
	query := `SELECT
	es.es_id,
    s.se_seat_row,
    s.se_seat_number,
    s.se_seat_type,
    es.es_status,
    es.es_price
    FROM event_seat es
    JOIN seat s ON es.es_seat_id = s.se_id
    WHERE es.es_event_id = $1`

	if isAvailable {
		query += ` AND es.es_status = 'available'`
	}

	query += ` ORDER BY s.se_seat_row, s.se_seat_number`

	err := pgxscan.Select(ctx, ticketRepository.Pool, &seats, query, eventId)
	if err != nil {
		return nil, err
	}
	return seats, nil
}

func (ticketRepository *TicketRepository) CreateTicket(ticket domain.Ticket) error {
	ctx := context.Background()
	query := `INSERT INTO ticket (ti_event_id, ti_seat_id, ti_user_id) VALUES ($1, $2, $3) RETURNING ti_id, ti_created_at`

	err := ticketRepository.Pool.QueryRow(
		ctx,
		query,
		ticket.TiEventID,
		ticket.TiSeatID,
		ticket.TiUserID).Scan(&ticket.TiID, &ticket.TiCreatedAt)
	if err != nil {
		return err
	}

	log.Info("Ticket created: ", ticket)

	return nil
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
