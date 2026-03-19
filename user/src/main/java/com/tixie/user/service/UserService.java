package com.tixie.user.service;

import com.tixie.user.data.dto.UserRegistrationRequest;
import com.tixie.user.data.dto.UserResponse;


import java.util.List;

public interface UserService {

    List<UserResponse> findAll();

    UserResponse findById(int Id);

    UserResponse registerUser(UserRegistrationRequest u);

    void deleteById(int Id);

    public UserResponse updateUser(int id, UserRegistrationRequest request);

}
