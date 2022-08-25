package com.github.exbotanical.resource.errors;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;


/**
 * An exception representing metadata about an invalid client-provided input.
 */
@ResponseStatus(value = HttpStatus.BAD_REQUEST)
public class InvalidInputException extends GouacheException {
  private static final String fallbackMessage = "The provided input was not valid.";

  public InvalidInputException() {
    super();
  }

  public InvalidInputException(String message) {
    super(fallbackMessage, message);
  }

  public InvalidInputException(String message, Throwable cause) {
    super(fallbackMessage, message, cause);
  }

  public InvalidInputException(Throwable cause) {
    super(cause);
  }

  protected InvalidInputException(String message, Throwable cause,
      boolean enableSuppression, boolean writableStackTrace) {
    super(fallbackMessage, message, cause, enableSuppression, writableStackTrace);
  }
}
