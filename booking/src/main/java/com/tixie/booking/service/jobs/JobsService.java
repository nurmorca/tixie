package com.tixie.booking.service.jobs;

import com.tixie.booking.client.TicketApiClient;
import com.tixie.booking.data.dto.BookingRequestDTO;
import com.tixie.booking.data.entity.Booking;
import com.tixie.booking.data.entity.BookingItems;
import com.tixie.booking.repository.BookingRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import java.sql.Timestamp;
import java.time.LocalDateTime;
import java.util.List;
import java.util.concurrent.TimeUnit;

@Component
public class JobsService {

    private static final long fixedRate = 5;
    private BookingRepository bookingRepository;
    private TicketApiClient ticketApiClient;

    @Autowired
    public JobsService(BookingRepository repo, TicketApiClient client) {
        this.bookingRepository = repo;
        this.ticketApiClient = client;
    }

    @Scheduled(fixedRate = fixedRate, timeUnit = TimeUnit.MINUTES)
    public void freeAbandonedTickets() {
       List<Booking> unpaidBookings = bookingRepository.getNotPaidBookingsOlderThanFiveMins(Booking.PAYMENT_EXPECTED,
               Timestamp.valueOf(LocalDateTime.now().minusMinutes(fixedRate)));
       for (Booking b : unpaidBookings) {
           int[] itemIds = b.getBookingItems().stream().mapToInt(BookingItems::getBiTicketId).toArray();
           BookingRequestDTO dto = new BookingRequestDTO(itemIds);
           ticketApiClient.freeTickets(dto);
           b.setBoStatus(Booking.ABANDONED);
           bookingRepository.save(b);
       }
    }
}
