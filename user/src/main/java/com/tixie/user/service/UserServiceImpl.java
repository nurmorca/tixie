package com.tixie.user.service;

import com.tixie.user.data.dto.UserRegistrationRequest;
import com.tixie.user.data.dto.UserResponse;
import com.tixie.user.data.entity.User;
import com.tixie.user.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.web.server.ResponseStatusException;

import java.util.List;
import java.util.Optional;

@Service
public class UserServiceImpl implements UserService{

    private UserRepository userRepository;
    private final BCryptPasswordEncoder passwordEncoder;

    @Autowired
    public UserServiceImpl (UserRepository repository) {
        this.userRepository = repository;
        this.passwordEncoder = new BCryptPasswordEncoder();
    }

    @Override
    public List<UserResponse> findAll() {
        List<UserResponse> responses = userRepository.findAll().stream()
                .map(this::convertToUserResponse)
                .toList();
        if (responses.isEmpty()) {
            throw new ResponseStatusException(HttpStatus.NOT_FOUND, "Users not found");
        }
        return responses;
    }

    @Override
    public UserResponse findById(int Id) {
        Optional<User> result = userRepository.findById(Id);
        if (result.isEmpty()) {
          throw new ResponseStatusException(HttpStatus.NOT_FOUND, "User with not found with ID: " + Id);
        }
        return convertToUserResponse(result.get());
    }

    @Override
    public UserResponse registerUser(UserRegistrationRequest request) {
        if (userRepository.existsByusEmail(request.getEmail())) {
            throw new ResponseStatusException(HttpStatus.CONFLICT, "E-mail already registered");
        }
        // validate password logic here
        request.setPassword(passwordEncoder.encode(request.getPassword()));
        return convertToUserResponse(userRepository.save(convertToUserEntity(request)));
    }

    @Override
    public void deleteById(int Id) {
       userRepository.deleteById(Id);
    }

    @Override
    public UserResponse updateUser(int id, UserRegistrationRequest request) {
        User user = userRepository.findById(id)
                .orElseThrow(() -> new ResponseStatusException(
                        HttpStatus.NOT_FOUND,
                        "User not found with id: " + id
                ));

        if (request.getFirstName() != null) {
            user.setUsFirstName(request.getFirstName());
        }
        if (request.getLastName() != null) {
            user.setUsLastName(request.getLastName());
        }
        if (request.getPhoneNumber() != null) {
            user.setUsPhoneNumber(request.getPhoneNumber());
        }
        if (request.getEmail() != null) {
            user.setUsEmail(request.getEmail());
        }
        if (request.getPassword() != null) {
            // validate password logic here
            request.setPassword(passwordEncoder.encode(request.getPassword()));
            user.setUsPassword(request.getPassword());
        }

        User updatedUser = userRepository.save(user);
        return convertToUserResponse(updatedUser);
    }

    private User convertToUserEntity(UserRegistrationRequest request) {
        User user = new User();
        user.setUsFirstName(request.getFirstName());
        user.setUsLastName(request.getLastName());
        user.setUsEmail(request.getEmail());
        user.setUsPhoneNumber(request.getPhoneNumber());
        user.setUsPassword(request.getPassword());
        return user;
    }

    private UserResponse convertToUserResponse(User user) {
        return new UserResponse(user.getUsId(), user.getUsEmail(), user.getUsFirstName(), user.getUsLastName(), user.getUsPhoneNumber(), user.getUsCreatedAt());
    }
}
