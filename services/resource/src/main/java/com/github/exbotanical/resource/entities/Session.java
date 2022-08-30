package com.github.exbotanical.resource.entities;

import java.util.Date;

import com.fasterxml.jackson.annotation.JsonProperty;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class Session {
  @JsonProperty("username")
  public String username;

  @JsonProperty("expiry")
  public Date expiry;
}
