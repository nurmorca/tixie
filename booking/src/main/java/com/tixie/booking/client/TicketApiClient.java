package com.tixie.booking.client;

import com.tixie.booking.data.dto.BookingRequestDTO;
import com.tixie.booking.data.dto.TicketsDTO;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClient;

import java.util.List;

@Service
public class TicketApiClient {

    private final RestClient restClient;
    private static final Logger log = LoggerFactory.getLogger(TicketApiClient.class);

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

    public String confirmBooking(BookingRequestDTO dto) {
        return restClient.post()
                .uri("/api/ticket/completeReservation")
                .body(dto)
                .retrieve()
                .body(String.class);
    }

    public List<TicketsDTO> getTickets(BookingRequestDTO dto) {
        log.info("Sending booking request: {}", dto.toString());
        TicketsDTO[] tixes = restClient.post()
                .uri("/api/ticket/initiateTicketReservation")
                .body(dto)
                .retrieve()
                .body(TicketsDTO[].class);

        return List.of(tixes);
    }
}
