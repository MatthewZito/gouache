package com.github.exbotanical.session.models;

import lombok.AllArgsConstructor;
import lombok.Builder;

@Builder
@AllArgsConstructor
public class UserSessionResponse {
  public final String username;

  public final int exp;
}
