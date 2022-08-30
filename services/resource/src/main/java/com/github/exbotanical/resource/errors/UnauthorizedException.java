package com.github.exbotanical.resource.errors;

import com.github.exbotanical.resource.meta.Constants;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(value = HttpStatus.UNAUTHORIZED)
public class UnauthorizedException extends GouacheException {

  public UnauthorizedException() {
  }

  public UnauthorizedException(String internal) {
    super(Constants.E_UNAUTHORIZED, internal);
  }

  public UnauthorizedException(String internal, Throwable cause) {
    super(Constants.E_UNAUTHORIZED, internal, cause);
  }

  public UnauthorizedException(Throwable cause) {
    super(cause);
  }

  public UnauthorizedException(String internal, Throwable cause,
                               boolean enableSuppression, boolean writableStackTrace) {
    super(Constants.E_UNAUTHORIZED, internal, cause, enableSuppression, writableStackTrace);
  }
}
