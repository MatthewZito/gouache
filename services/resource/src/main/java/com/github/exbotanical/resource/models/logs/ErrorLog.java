package com.github.exbotanical.resource.models.logs;

import lombok.Getter;
import lombok.experimental.SuperBuilder;

/**
 * An entity for storing error metadata in a Report.
 */
@SuperBuilder
@Getter
public class ErrorLog extends RequestLog {
  String message;

  String cause;
}
