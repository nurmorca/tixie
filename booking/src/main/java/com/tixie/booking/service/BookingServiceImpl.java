package com.tixie.booking.service;

import com.tixie.booking.client.TicketApiClient;
import com.tixie.booking.client.UserApiClient;
import com.tixie.booking.data.dto.BookingRequestDTO;
import com.tixie.booking.data.dto.ConfirmBookingDTO;
import com.tixie.booking.data.dto.UserDTO;
import com.tixie.booking.data.entity.Booking;
import com.tixie.booking.data.entity.BookingItems;
import com.tixie.booking.repository.BookingItemRepository;
import com.tixie.booking.repository.BookingRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

@Service
public class BookingServiceImpl implements BookingService {

    private BookingRepository bookingRepository;
    private BookingItemRepository bookingItemRepository;
    private UserApiClient userApiClient;
    private Logger logger = LoggerFactory.getLogger(BookingServiceImpl.class);
    private TicketApiClient ticketApiClient;

    @Autowired
    public BookingServiceImpl (BookingRepository bookingRepository, BookingItemRepository bookingItemRepository, UserApiClient userApiClient, TicketApiClient ticketApiClient) {
        this.bookingRepository = bookingRepository;
        this.bookingItemRepository = bookingItemRepository;
        this.userApiClient = userApiClient;
        this.ticketApiClient = ticketApiClient;
    }

    @Override
    public List<Booking> getAllBookings() {
        return bookingRepository.findBookingAndItems();
    }

    @Override
    public Booking getBookingById(int bookingId) {
        return bookingRepository.findBookingAndItemsById(bookingId);
    }

    @Override
    public void cancelBooking(int bookingId) {
        Optional<Booking> canceledBooking  = bookingRepository.findById(bookingId);
        if (canceledBooking.isPresent()) {
            Booking booking = canceledBooking.get();
            booking.setBoStatus("CANCELLED");
            bookingRepository.save(booking);
        }
        // will be adding a custom error message in case of an unsuccessful retrieval.
    }

    @Override
    public String getStatusForBooking(int bookingId) {
        return bookingRepository.getBookingStatusById(bookingId);
    }

    @Override
    public Booking createBooking(BookingRequestDTO bookingRequestDTO) {
        if (verifyUserExist(bookingRequestDTO.getUserId()) == null) {
            return null;
        }
        Booking newBooking = new Booking();
        newBooking.setBoUserId(bookingRequestDTO.getUserId());
        newBooking.setBoStatus("WAITING_CONFIRMATION");
        newBooking.setBoTotalPrice(BigDecimal.valueOf(888L));
       Booking booking =  bookingRepository.save(newBooking);
        for (int id: bookingRequestDTO.getTicketIds()) {
            BookingItems item = new BookingItems();
            item.setBiEventId(bookingRequestDTO.getEventId());
            item.setBiTicketId(id);
            item.setBiBookingId(booking.getBoId());
            item.setBiPrice(newBooking.getBoTotalPrice());
            bookingItemRepository.save(item);
        }
        return booking;
    }

    @Override
    public List<Booking> confirmBooking(ConfirmBookingDTO confirmBookingDTO) {
        Booking booking = getBookingById(confirmBookingDTO.)
        ticketApiClient.confirmBooking(confirmBookingDTO);
        List<Booking> bookings = new ArrayList<>();

       for (int bookingId : confirmBookingDTO.getTicketIds()) {
           booking = getBookingById(bookingId);
           booking.setBoStatus("BOOKED");
           bookings.add(bookingRepository.save(booking));
       }
        return bookings;
    }

    private UserDTO verifyUserExist(int userId) {
        if (userId == 0) {
            return null;
        }
        return userApiClient.getUserById(userId);
    }
}
