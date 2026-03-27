package com.tixie.booking.data.entity;


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
    @Column(name="BI_SEAT_NUMBER", nullable = false, length = 80)
    private String biSeatNumber;
    @Column(name="BO_CREATED_AT", nullable = false)
    private Timestamp biCreatedAt;

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

    public String getBiSeatNumber() {
        return biSeatNumber;
    }

    public void setBiSeatNumber(String biSeatNumber) {
        this.biSeatNumber = biSeatNumber;
    }

    public Timestamp getBiCreatedAt() {
        return biCreatedAt;
    }

    public void setBiCreatedAt(Timestamp biCreatedAt) {
        this.biCreatedAt = biCreatedAt;
    }
}
