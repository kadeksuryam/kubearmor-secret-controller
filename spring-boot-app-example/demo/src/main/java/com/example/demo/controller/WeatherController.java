package com.example.demo.controller;

import com.example.demo.dto.GetWeatherOutboundResDTO;
import com.example.demo.dto.GetWeatherResDTO;
import com.example.demo.outbound.RestService;
import com.example.demo.outbound.WeatherService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

@RestController
public class WeatherController {
    @Autowired
    private WeatherService weatherService;

    @GetMapping("/{query}/current")
    public GetWeatherResDTO getCurrentWeather(@PathVariable String query) {
        GetWeatherOutboundResDTO outboundResDTO = weatherService.getCurrentWeather(query);
        GetWeatherResDTO resDTO = new GetWeatherResDTO();

        resDTO.setTemperatureC(outboundResDTO.getCurrent().getTempC());
        resDTO.setHumidity(outboundResDTO.getCurrent().getHumidity());

        return resDTO;
    }
}
