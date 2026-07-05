package controller

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
	"ticket/data/domain"
	"ticket/data/dto"
	"ticket/service"

	"github.com/labstack/echo/v4"
)

type TicketController struct {
	ticketService service.ITicketService
}

func NewTicketController(service service.ITicketService) *TicketController {
	return &TicketController{
		ticketService: service,
	}
}

func (ticketController *TicketController) RegisterRoutes(e *echo.Echo) {
	// event handling
	e.GET("/api/event/:id", ticketController.GetEventById)
	e.GET("/api/event/", ticketController.GetAllEvents)
	e.POST("/api/event/", ticketController.CreateEvent)
	e.PUT("/api/event/:id", ticketController.UpdateDescription)
	e.DELETE("/api/event/:id", ticketController.DeleteEvent)
	e.GET("/api/event/:id/seats", ticketController.GetSeatsForEvent)
	e.GET("/api/event/:id/tickets", ticketController.GetAllTicketsForEvent)

	// ticket handling
	e.GET("/api/ticket/available/:eventId", ticketController.GetAvailableSeatsForEvent)
	e.POST("/api/ticket/lock", ticketController.LockTickets)
	e.POST("/api/ticket/release", ticketController.ReleaseTickets)
	e.POST("/api/ticket/checkReservation", ticketController.CheckUserReservation)
	e.POST("/api/ticket/completeReservation", ticketController.CompleteReservation)
	e.POST("/api/ticket/initiateTicketReservation", ticketController.InitiateTicketReservation)
	e.POST("/api/ticket/freeTickets", ticketController.FreeTickets)
}

func (ticketController TicketController) FreeTickets(c echo.Context) error {
	body, _ := io.ReadAll(c.Request().Body)
	c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

	var userTickets dto.UserTicketsDTO
	err := c.Bind(&userTickets)
	if err != nil {
		log.Printf("Bind error: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = ticketController.ticketService.FreeTickets(c.Request().Context(), userTickets)
	if err != nil {
		log.Printf("Service error: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (ticketController TicketController) InitiateTicketReservation(c echo.Context) error {
	body, _ := io.ReadAll(c.Request().Body)
	c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

	var userTickets dto.UserTicketsDTO
	err := c.Bind(&userTickets)
	if err != nil {
		log.Printf("Bind error: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	bookingTickets, err := ticketController.ticketService.InitiateTicketReservation(c.Request().Context(), userTickets)
	if err != nil {
		log.Printf("Service error: %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, bookingTickets)
}

func (ticketController TicketController) CompleteReservation(c echo.Context) error {
	var userTickets dto.UserTicketsDTO
	err := c.Bind(&userTickets)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ticketController.ticketService.CompleteReservation(c.Request().Context(), userTickets)

	if err != nil {
		return c.JSON(http.StatusNotFound, "ticket(s) couldnt be reserved for user")
	}

	return c.JSON(http.StatusOK, "completed successfully")
}

func (ticketController TicketController) CheckUserReservation(c echo.Context) error {
	var userBooking dto.UserTicketsDTO
	err := c.Bind(&userBooking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := ticketController.ticketService.CheckUserReservations(c.Request().Context(), userBooking)

	if result == false {
		return c.JSON(http.StatusNotFound, "ticket(s) are not reserved for user")
	}

	return c.JSON(http.StatusOK, result)
}

func (ticketController TicketController) GetAllTicketsForEvent(c echo.Context) error {
	param := c.Param("id")
	eventId, _ := strconv.ParseInt(param, 10, 64)

	tixes, err := ticketController.ticketService.GetAllTicketsForEvent(eventId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, tixes)
}

func (ticketController TicketController) GetSeatsForEvent(c echo.Context) error {
	param := c.Param("id")
	eventId, _ := strconv.ParseInt(param, 10, 64)

	seats, err := ticketController.ticketService.GetTicketsForEvent(eventId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, seats)
}

func (ticketController TicketController) GetAvailableSeatsForEvent(c echo.Context) error {
	param := c.Param("eventId")
	eventId, _ := strconv.ParseInt(param, 10, 64)

	seats, err := ticketController.ticketService.GetAvailableTicketsForEvent(eventId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, seats)
}

func (ticketController TicketController) LockTickets(c echo.Context) error {
	var userBooking dto.UserTicketsDTO
	err := c.Bind(&userBooking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ticketController.ticketService.LockTickets(c.Request().Context(), userBooking)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func (ticketController TicketController) ReleaseTickets(c echo.Context) error {
	var userBooking dto.UserTicketsDTO
	err := c.Bind(&userBooking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ticketController.ticketService.ReleaseTickets(c.Request().Context(), userBooking)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func (ticketController TicketController) GetEventById(c echo.Context) error {
	param := c.Param("id")
	eventId, _ := strconv.ParseInt(param, 10, 64)

	event, err := ticketController.ticketService.GetEventById(eventId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, event)
}

func (ticketController TicketController) GetAllEvents(c echo.Context) error {
	events, err := ticketController.ticketService.GetAllEvents()

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, events)
}

func (ticketController TicketController) CreateEvent(c echo.Context) error {
	var eventRequest domain.Event
	err := c.Bind(&eventRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ticketController.ticketService.CreateEvent(eventRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func (ticketController TicketController) UpdateDescription(c echo.Context) error {
	param := c.Param("id")
	eventId, _ := strconv.ParseInt(param, 10, 64)
	newDescParam := c.QueryParam("description")
	if len(newDescParam) == 0 {
		return c.JSON(http.StatusBadRequest, "A valid description should be entered")
	}
	err := ticketController.ticketService.UpdateDescription(eventId, newDescParam)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (ticketController TicketController) DeleteEvent(c echo.Context) error {
	param := c.Param("id")
	eventId, _ := strconv.ParseInt(param, 10, 64)
	err := ticketController.ticketService.DeleteEvent(eventId)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
