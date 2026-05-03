package com.tixie.booking.client;

import com.tixie.booking.data.dto.UserDTO;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClient;

@Service
public class UserApiClient {

    private final RestClient restClient;

    public UserApiClient(@Qualifier("userRestClient") RestClient restClient) {
        this.restClient = restClient;
    }

    public UserDTO getUserById(long userId) {
        return restClient.get()
                .uri("/api/user/{id}", userId)
                .retrieve()
                .body(UserDTO.class);
    }
}
