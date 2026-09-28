package com.tixie.user.repository;

import com.tixie.user.data.entity.User;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository extends JpaRepository<User, Integer> {

    boolean existsByusEmail(String email);
     Optional<User> findByUsEmail(String email);
}
