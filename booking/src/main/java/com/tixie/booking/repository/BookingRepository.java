package com.tixie.booking.repository;

import com.tixie.booking.data.entity.Booking;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.sql.Timestamp;
import java.util.List;

public interface BookingRepository extends JpaRepository<Booking, Integer> {

    @Query("SELECT b FROM Booking b LEFT JOIN FETCH b.boBookingItems bi WHERE b.boId = :bookingId")
    Booking findBookingAndItemsById(@Param("bookingId") int bookingId);

    @Query("SELECT b FROM Booking b LEFT JOIN FETCH b.boBookingItems bi WHERE 1=1")
    List<Booking> findBookingAndItems();

    @Query("SELECT b.boStatus FROM Booking b LEFT JOIN b.boBookingItems bi WHERE b.boId = :bookingId")
    String getBookingStatusById(@Param("bookingId") int bookingId);

    @Query("SELECT b FROM Booking b LEFT JOIN FETCH b.boBookingItems bi WHERE b.boStatus = :bookingStatus AND b.boLastPaymentAttemptAt <= :time")
    List<Booking> getNotPaidBookingsOlderThanFiveMins(@Param("bookingStatus") String status, @Param("time") Timestamp time);

}
