package com.tixie.booking.service;

import com.tixie.booking.data.dto.BookingRequestDTO;
import com.tixie.booking.data.entity.Booking;

import java.util.List;

public interface BookingService {

    List<Booking> getAllBookings();

    Booking getBookingById(int BookingId);

    void cancelBooking(int BookingId); // signature might change

    String getStatusForBooking(int BookingId);

    Booking createBooking(BookingRequestDTO bookingRequestDTO);
}
