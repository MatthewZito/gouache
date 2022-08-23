package com.github.exbotanical.session.errors;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(value = HttpStatus.INTERNAL_SERVER_ERROR)
public class UserInvariantViolationException extends Exception {
  public UserInvariantViolationException() {
    super();
  }

  public UserInvariantViolationException(String username) {
    // @todo friendly
    super(String.format("user with username %s was duplicated", username));
  }

  public UserInvariantViolationException(String message, Throwable cause) {
    super(message, cause);
  }

  public UserInvariantViolationException(Throwable cause) {
    super(cause);
  }

  protected UserInvariantViolationException(String message, Throwable cause,
      boolean enableSuppression, boolean writableStackTrace) {
    super(message, cause, enableSuppression, writableStackTrace);
  }
}
