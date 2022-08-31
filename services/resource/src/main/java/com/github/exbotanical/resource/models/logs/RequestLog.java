package com.github.exbotanical.resource.models.logs;

import java.util.Map;
import lombok.Getter;
import lombok.experimental.SuperBuilder;

/**
 * An entity for storing request metadata in a Report.
 */
@SuperBuilder
@Getter
public class RequestLog {
  String path;

  String method;

  Map<String, Object> parameters;
}
