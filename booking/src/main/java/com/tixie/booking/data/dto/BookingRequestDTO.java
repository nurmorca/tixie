package com.tixie.booking.data.dto;

public class BookingRequestDTO {

    private int[] ticketIds;
    private int eventId;
    private int userId;

    public int getEventId() {
        return eventId;
    }

    public void setEventId(int eventId) {
        this.eventId = eventId;
    }

    public int getUserId() {
        return userId;
    }

    public void setUserId(int userId) {
        this.userId = userId;
    }

    public int[] getTicketIds() {
        return ticketIds;
    }

    public void setTicketIds(int[] ticketIds) {
        this.ticketIds = ticketIds;
    }
}
