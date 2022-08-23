package com.github.exbotanical.session.errors;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(value = HttpStatus.NOT_FOUND)
public class UserNotFoundException extends Exception {
  public UserNotFoundException() {
    super();
  }

  public UserNotFoundException(String username) {
    super(String.format("user with username %s not found", username));
  }

  public UserNotFoundException(String message, Throwable cause) {
    super(message, cause);
  }

  public UserNotFoundException(Throwable cause) {
    super(cause);
  }

  protected UserNotFoundException(String message, Throwable cause, boolean enableSuppression,
      boolean writableStackTrace) {
    super(message, cause, enableSuppression, writableStackTrace);
  }
}
