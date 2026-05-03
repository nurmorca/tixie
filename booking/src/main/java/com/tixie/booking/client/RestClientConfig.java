package com.tixie.booking.client;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.client.HttpComponentsClientHttpRequestFactory;
import org.springframework.web.client.RestClient;

@Configuration
public class RestClientConfig
{
    @Value("${services.user.url}")
    private String userUrl;
    @Value("${services.ticket.url}")
    private String ticketUrl;

    @Bean
    public RestClient ticketRestClient()
    {
        return RestClient.builder()
                .baseUrl(ticketUrl)
                .requestFactory(getHttpComponentsClientHttpRequestFactory())
                .build();
    }

    @Bean
    public RestClient userRestClient() {
        return RestClient.builder()
                .baseUrl(userUrl)
                .requestFactory(getHttpComponentsClientHttpRequestFactory())
                .build();
    }

    private HttpComponentsClientHttpRequestFactory getHttpComponentsClientHttpRequestFactory() {
        HttpComponentsClientHttpRequestFactory factory = new HttpComponentsClientHttpRequestFactory();
        factory.setConnectionRequestTimeout(10_000);
        factory.setReadTimeout(10_000);
        return factory;
    }
}