package com.tixie.booking.service;

import com.tixie.booking.data.dto.BookingRequestDTO;
import com.tixie.booking.data.entity.Booking;
import com.tixie.booking.repository.BookingRepository;
import org.springframework.beans.factory.annotation.Autowired;

import java.util.List;
import java.util.Optional;

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
    public void cancelBooking(int BookingId) {

    }

    @Override
    public String getStatusForBooking(int bookingId) {
        return bookingRepository.getBookingStatusById(bookingId);
    }

    @Override
    public Booking createBooking(BookingRequestDTO bookingRequestDTO) {
        return null;
    }
}
