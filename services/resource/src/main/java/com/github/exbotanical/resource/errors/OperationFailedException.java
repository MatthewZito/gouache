package com.github.exbotanical.resource.errors;

public class OperationFailedException extends GouacheException {
  public OperationFailedException() {
  }

  public OperationFailedException(String friendly, String internal) {
    super(friendly, internal);
  }

  public OperationFailedException(String friendly, String internal, Throwable cause) {
    super(friendly, internal, cause);
  }

  public OperationFailedException(Throwable cause) {
    super(cause);
  }

  public OperationFailedException(String friendly, String internal, Throwable cause, boolean enableSuppression, boolean writableStackTrace) {
    super(friendly, internal, cause, enableSuppression, writableStackTrace);
  }
}
