package com.tixie.booking.data.entity;

import jakarta.persistence.*;

import java.math.BigDecimal;
import java.sql.Timestamp;
import java.util.List;

@Entity
@Table(name="BOOKING")
public class Booking {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name="BO_ID")
    private int boId;
    @Column(name="BO_USER_ID", nullable = false)
    private int boUserId;
    @Column(name="BO_TOTAL_PRICE", nullable = false)
    private BigDecimal boTotalPrice;
    @Column(name="BO_STATUS", nullable = false, length = 80)
    private String boStatus;
    @Column(name="BO_CREATED_AT")
    private Timestamp boCreatedAt;

    @OneToMany(fetch = FetchType.LAZY) // FetchType.EAGER fetches the associated Product automatically
    @JoinColumn(name = "BI_BOOKING_ID") // Specifies the foreign key column in the 'transaction' table
    private List<BookingItems> boBookingItems;

    public List<BookingItems> getBookingItems() {
        return boBookingItems;
    }

    public void setBookingItems(List<BookingItems> bookingItems) {
        this.boBookingItems = bookingItems;
    }

    public int getBoId() {
        return boId;
    }

    public void setBoId(int boId) {
        this.boId = boId;
    }

    public int getBoUserId() {
        return boUserId;
    }

    public void setBoUserId(int boUserId) {
        this.boUserId = boUserId;
    }

    public BigDecimal getBoTotalPrice() {
        return boTotalPrice;
    }

    public void setBoTotalPrice(BigDecimal boTotalPrice) {
        this.boTotalPrice = boTotalPrice;
    }

    public String getBoStatus() {
        return boStatus;
    }

    public void setBoStatus(String boStatus) {
        this.boStatus = boStatus;
    }

    public Timestamp getBoCreatedAt() {
        return boCreatedAt;
    }

    public void setBoCreatedAt(Timestamp boCreatedAt) {
        this.boCreatedAt = boCreatedAt;
    }
}


