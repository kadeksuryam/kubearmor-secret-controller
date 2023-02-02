package com.example.demo.outbound;

import com.example.demo.dto.GetWeatherOutboundResDTO;
import com.example.demo.dto.GetWeatherResDTO;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.stereotype.Service;

@Service
public class WeatherService extends RestService {
    @Value("${WEATHER_API_KEY}")
    private String API_KEY;
    public WeatherService(RestTemplateBuilder restTemplateBuilder) {
        super(restTemplateBuilder);
    }

    public GetWeatherOutboundResDTO getCurrentWeather(String query) {
        String serverURI = String.format("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no",API_KEY, query);

        // System.out.println(serverURI);
        return this.restTemplate.getForObject(serverURI, GetWeatherOutboundResDTO.class);
    }
}
