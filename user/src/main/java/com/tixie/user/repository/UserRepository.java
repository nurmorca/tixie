package com.tixie.user.repository;

import com.tixie.user.data.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository extends JpaRepository<User, Integer> {

    boolean existsByusEmail(String email);


}
