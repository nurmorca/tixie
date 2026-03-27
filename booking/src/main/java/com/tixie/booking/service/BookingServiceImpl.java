package com.tixie.booking.service;

import com.tixie.booking.data.dto.BookingRequestDTO;
import com.tixie.booking.data.entity.Booking;

import java.util.List;

public class BookingServiceImpl implements BookingService {
    @Override
    public List<Booking> getAllBookings() {
        return List.of();
    }

    @Override
    public Booking getBookingById(int BookingId) {
        return null;
    }

    @Override
    public void cancelBooking(int BookingId) {

    }

    @Override
    public String getStatusForBooking(int BookingId) {
        return "";
    }

    @Override
    public Booking createBooking(BookingRequestDTO bookingRequestDTO) {
        return null;
    }
}
