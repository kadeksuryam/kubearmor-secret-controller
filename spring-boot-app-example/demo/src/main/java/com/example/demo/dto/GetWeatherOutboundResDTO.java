package com.example.demo.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class GetWeatherOutboundResDTO {
    private Current current;
    @Getter
    @Setter
    public static class Current {
        @JsonProperty("temp_c")
        private Double tempC;
        @JsonProperty("humidity")
        private Double humidity;
    }
}

