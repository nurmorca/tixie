package com.tixie.booking.service;

import com.tixie.booking.data.dto.BookingRequestDTO;
import com.tixie.booking.data.entity.Booking;
import com.tixie.booking.data.entity.BookingItems;
import com.tixie.booking.repository.BookingRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.util.List;
import java.util.Optional;

@Service
public class BookingServiceImpl implements BookingService {

    private BookingRepository bookingRepository;

    @Autowired
    public BookingServiceImpl (BookingRepository bookingRepository) {
        this.bookingRepository = bookingRepository;
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
        // implement checking ticket and user logic here.
        BookingItems item = new BookingItems();
        item.setBiEventId(bookingRequestDTO.getEventId());
        item.setBiTicketId(bookingRequestDTO.getTicketId());
        Booking newBooking = new Booking();
        newBooking.setBoUserId(bookingRequestDTO.getUserId());
        newBooking.setBoStatus("BOOKED");
        newBooking.setBoTotalPrice(BigDecimal.valueOf(888L)); //TODO: fix after service connection
        newBooking.setBookingItems(List.of(item));
        return bookingRepository.save(newBooking);
    }
}
