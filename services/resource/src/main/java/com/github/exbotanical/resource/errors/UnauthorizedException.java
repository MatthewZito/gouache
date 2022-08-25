package com.github.exbotanical.resource.errors;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(value = HttpStatus.UNAUTHORIZED)
public class UnauthorizedException extends GouacheException {
  private static final String friendly = "You are not authorized to access this resource.";

  public UnauthorizedException() {}

  public UnauthorizedException(String internal) {
    super(friendly, internal);
  }

  public UnauthorizedException(String internal, Throwable cause) {
    super(friendly, internal, cause);
  }

  public UnauthorizedException(Throwable cause) {
    super(cause);
  }

  public UnauthorizedException(String internal, Throwable cause,
      boolean enableSuppression, boolean writableStackTrace) {
    super(friendly, internal, cause, enableSuppression, writableStackTrace);
  }
}
