package com.github.exbotanical.resource.entities;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.Date;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * A spec-compliant user session object qua the gouache/cache and gouache/auth services.
 */
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
