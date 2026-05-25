package com.tixie.booking.repository;

import com.tixie.booking.data.entity.BookingItems;
import org.springframework.data.jpa.repository.JpaRepository;

public interface BookingItemRepository extends JpaRepository<BookingItems, Integer> {
}
