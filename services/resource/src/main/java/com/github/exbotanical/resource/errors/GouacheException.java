package com.github.exbotanical.resource.errors;

import lombok.Builder;
import lombok.Getter;

@Getter
@Builder
public class GouacheException extends Exception {
  private String friendly;
  private String internal;

  public GouacheException() {
    super();
  }

  public GouacheException(String friendly, String internal) {
    super(formatMessageForLogging(friendly, internal));

    this.friendly = friendly;
    this.internal = internal;
  }

  public GouacheException(String friendly, String internal, Throwable cause) {
    super(formatMessageForLogging(friendly, internal), cause);

    this.friendly = friendly;
    this.internal = internal;
  }

  public GouacheException(Throwable cause) {
    super(cause);
  }

  protected GouacheException(String friendly, String internal, Throwable cause,
      boolean enableSuppression, boolean writableStackTrace) {
    super(formatMessageForLogging(friendly, internal), cause, enableSuppression,
        writableStackTrace);

    this.friendly = friendly;
    this.internal = internal;
  }


  private static String formatMessageForLogging(String friendly, String internal) {
    return String.format("{ \"friendly\": %s, \"internal\": %s }", friendly, internal);
  }
}
