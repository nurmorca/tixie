package com.tixie.booking.client;

import com.tixie.booking.data.dto.BookingRequestDTO;
import com.tixie.booking.data.dto.UserDTO;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClient;

@Service
public class TicketApiClient {

    private final RestClient restClient;

    public TicketApiClient(@Qualifier("ticketRestClient") RestClient restClient) {
        this.restClient = restClient;
    }

    public String lockSeat(BookingRequestDTO bookingRequestDTO) {
        return restClient.post()
                .uri("/api/ticket/lock")
                .body(bookingRequestDTO)
                .retrieve()
                .body(String.class);

    }
}
