package controller

import (
	"net/http"
	"strconv"
	"ticket/data/domain"
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
	e.GET("/api/event/:id", ticketController.GetEventById)
	e.GET("/api/event/", ticketController.GetAllEvents)
	e.POST("/api/event/", ticketController.CreateEvent)
	e.PUT("/api/event/:id", ticketController.UpdateDescription)
	e.DELETE("/api/event/:id", ticketController.DeleteEvent)
}

func (ticketController TicketController) GetEventById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.ParseInt(param, 10, 64)

	event, err := ticketController.ticketService.GetEventById(productId)

	if err != nil {
		return c.JSON(http.StatusNotFound, "Event not found")
	}

	return c.JSON(http.StatusOK, event)
}

func (ticketController TicketController) GetAllEvents(c echo.Context) error {
	events, err := ticketController.ticketService.GetAllEvents()

	if err != nil {
		return c.JSON(http.StatusNotFound, "Events not found")
	}

	return c.JSON(http.StatusOK, events)
}

func (ticketController TicketController) CreateEvent(c echo.Context) error {
	var eventRequest domain.Event
	err := c.Bind(&eventRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Event cannot be mapped")
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
		return c.JSON(http.StatusNotFound, "Cannot change description")
	}

	return c.NoContent(http.StatusOK)
}

func (ticketController TicketController) DeleteEvent(c echo.Context) error {
	param := c.Param("id")
	eventId, _ := strconv.ParseInt(param, 10, 64)
	err := ticketController.ticketService.DeleteEvent(eventId)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Cannot delete event")
	}
	return c.NoContent(http.StatusOK)
}
