package com.tixie.booking.data.entity;


import com.fasterxml.jackson.annotation.JsonBackReference;
import jakarta.persistence.*;

import java.math.BigDecimal;
import java.sql.Timestamp;

@Entity
@Table(name="BOOKING_ITEMS")
public class BookingItems {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name="BI_ID")
    private int biId;
    @Column(name="BI_BOOKING_ID", nullable = false)
    private int biBookingId;
    @Column(name="BI_TICKET_ID", nullable = false)
    private int biTicketId;
    @Column(name="BI_EVENT_ID", nullable = false)
    private int biEventId;
    @Column(name="BI_PRICE", nullable = false)
    private BigDecimal biPrice;
    @Column(name="BO_CREATED_AT")
    private Timestamp biCreatedAt;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "BI_BOOKING_ID", referencedColumnName = "BO_ID", nullable = false, insertable = false, updatable = false)
    @JsonBackReference("booking-bookingItems")
    private Booking booking;

    public int getBiId() {
        return biId;
    }

    public void setBiId(int biId) {
        this.biId = biId;
    }

    public int getBiBookingId() {
        return biBookingId;
    }

    public void setBiBookingId(int biBookingId) {
        this.biBookingId = biBookingId;
    }

    public int getBiTicketId() {
        return biTicketId;
    }

    public void setBiTicketId(int biTicketId) {
        this.biTicketId = biTicketId;
    }

    public int getBiEventId() {
        return biEventId;
    }

    public void setBiEventId(int biEventId) {
        this.biEventId = biEventId;
    }

    public BigDecimal getBiPrice() {
        return biPrice;
    }

    public void setBiPrice(BigDecimal biPrice) {
        this.biPrice = biPrice;
    }

    public Timestamp getBiCreatedAt() {
        return biCreatedAt;
    }

    public void setBiCreatedAt(Timestamp biCreatedAt) {
        this.biCreatedAt = biCreatedAt;
    }

    public Booking getBooking() {
        return booking;
    }

    public void setBooking(Booking booking) {
        this.booking = booking;
    }
}
