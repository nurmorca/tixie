package repository

import (
	"context"
	"ticket/data/domain"

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
}

type TicketRepository struct {
	Pool *pgxpool.Pool
}

func NewTicketRepository(pool *pgxpool.Pool) ITicketRepository {
	return &TicketRepository{Pool: pool}
}

func (ticketRepository *TicketRepository) GetAllEvents() ([]domain.Event, error) {
	ctx := context.Background()
	var events []domain.Event
	query := `SELECT * FROM EVENTS ORDER BY EV_DATETIME DESC`

	err := pgxscan.Select(ctx, ticketRepository.Pool, &events, query)
	if err != nil {
		return nil, err
	}
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
	query := `INSERT INTO EVENTS (EV_NAME, EV_HOST, EV_VENUE, EV_DATETIME, EV_TOTAL_SEATS, EV_TICKETS_SOLD, EV_DESCRIPTION) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING EV_ID, EV_CREATED_AT`

	err := ticketRepository.Pool.QueryRow(
		ctx,
		query,
		event.EvName,
		event.EvHost,
		event.EvVenue,
		event.EvDateTime,
		event.EvTotalSeats,
		event.EvTicketSold,
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
