package com.github.exbotanical.resource.models.logs;

import lombok.Getter;
import lombok.experimental.SuperBuilder;

@SuperBuilder
@Getter
public class ErrorLog extends RequestLog {
  String message;

  String cause;
}

