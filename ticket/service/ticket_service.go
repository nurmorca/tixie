package service

import (
	"ticket/data/domain"
	"ticket/repository"
)

// add validation logic
type ITicketService interface {
	GetAllEvents() ([]domain.Event, error)
	CreateEvent(event domain.Event) error
	GetEventById(eventId int64) (*domain.Event, error)
	DeleteEvent(eventId int64) error
	UpdateDescription(eventId int64, description string) error
}

type TicketService struct {
	ticketRepository repository.ITicketRepository
}

func NewTicketService(repository repository.ITicketRepository) ITicketService {
	return &TicketService{
		ticketRepository: repository,
	}
}

func (ticketService *TicketService) CreateEvent(event domain.Event) error {
	return ticketService.ticketRepository.CreateEvent(event)
}

func (ticketService *TicketService) GetAllEvents() ([]domain.Event, error) {
	return ticketService.ticketRepository.GetAllEvents()
}

func (ticketService *TicketService) GetEventById(eventId int64) (*domain.Event, error) {
	return ticketService.ticketRepository.GetEventById(eventId)
}

func (ticketService *TicketService) DeleteEvent(eventId int64) error {
	return ticketService.ticketRepository.DeleteEvent(eventId)
}

func (ticketService *TicketService) UpdateDescription(eventId int64, description string) error {
	return ticketService.ticketRepository.UpdateDescription(eventId, description)
}
