package com.tixie.booking.data.dto;

import java.util.Arrays;

public class BookingRequestDTO {

    private int[] ticketIds;
    private int userId;

    public BookingRequestDTO(int[] ticketIds) {
        this.ticketIds = ticketIds;
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

    @Override
    public String toString() {
        return "BookingRequestDTO{" +
                "ticketIds=" + Arrays.toString(ticketIds) +
                ", userId=" + userId +
                '}';
    }
}
