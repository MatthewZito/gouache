package com.github.exbotanical.resource.errors;

import com.github.exbotanical.resource.meta.Constants;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

/**
 * An exception for invariant violations for a User object.
 */
@ResponseStatus(value = HttpStatus.INTERNAL_SERVER_ERROR)
public class UserInvariantViolationException extends GouacheException {

  private static String message;

  private static String format(final String username) {
    if (message == null) {
      message = String.format(Constants.E_USER_DUPE_INVARIANT_FMT, username);
    }

    return message;
  }

  public UserInvariantViolationException() {
    super();
  }

  public UserInvariantViolationException(String username) {
    // @todo friendly
    super(format(username), format(username));
  }

  public UserInvariantViolationException(String username, Throwable cause) {
    super(format(username), format(username), cause);
  }

  public UserInvariantViolationException(Throwable cause) {
    super(cause);
  }

  protected UserInvariantViolationException(String username, Throwable cause,
                                            boolean enableSuppression, boolean writableStackTrace) {
    super(format(username), format(username), cause, enableSuppression, writableStackTrace);
  }
}
