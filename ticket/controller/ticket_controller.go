package controller

import (
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
	e.GET("/api/event/:eventId/user/:userId", ticketController.GetUserTicketsForEvent)

	// ticket handling
	e.GET("/api/ticket/available/:eventId", ticketController.GetAvailableSeatsForEvent)
	e.POST("/api/ticket/lock", ticketController.LockSeat)
	e.POST("/api/ticket/release", ticketController.ReleaseSeat)
	e.POST("/api/ticket/confirm", ticketController.CreateTicket)
	e.GET("/api/ticket/user/:id", ticketController.GetAllTicketsForUser)
}

func (ticketController TicketController) GetUserTicketsForEvent(c echo.Context) error {
	param1 := c.Param("eventId")
	eventId, _ := strconv.ParseInt(param1, 10, 64)
	param2 := c.Param("userId")
	userId, _ := strconv.ParseInt(param2, 10, 64)

	tixes, err := ticketController.ticketService.GetUserTicketsForEvent(eventId, userId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, tixes)
}

func (ticketController TicketController) GetAllTicketsForUser(c echo.Context) error {
	param := c.Param("id")
	userId, _ := strconv.ParseInt(param, 10, 64)

	tixes, err := ticketController.ticketService.GetAllTicketsForEvent(userId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, tixes)
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

func (ticketController TicketController) CreateTicket(c echo.Context) error {
	var purchaseTicketRequest dto.TicketReservationDTO
	err := c.Bind(&purchaseTicketRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ticketController.ticketService.CreateTicket(purchaseTicketRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func (ticketController TicketController) GetSeatsForEvent(c echo.Context) error {
	param := c.Param("id")
	eventId, _ := strconv.ParseInt(param, 10, 64)

	seats, err := ticketController.ticketService.GetSeatsForEvent(eventId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, seats)
}

func (ticketController TicketController) GetAvailableSeatsForEvent(c echo.Context) error {
	param := c.Param("eventId")
	eventId, _ := strconv.ParseInt(param, 10, 64)

	seats, err := ticketController.ticketService.GetAvailableSeatsForEvent(eventId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, seats)
}

func (ticketController TicketController) LockSeat(c echo.Context) error {
	var lockSeatRequest dto.TicketReservationDTO
	err := c.Bind(&lockSeatRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ticketController.ticketService.LockSeat(c.Request().Context(), lockSeatRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func (ticketController TicketController) ReleaseSeat(c echo.Context) error {
	var releaseSeatRequest dto.TicketReservationDTO
	err := c.Bind(&releaseSeatRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ticketController.ticketService.ReleaseSeat(c.Request().Context(), releaseSeatRequest)
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
