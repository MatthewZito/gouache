package com.github.exbotanical.resource.models.reporting;

import lombok.Builder;
import lombok.Getter;

import java.util.Map;

/**
 * An entity for storing request metadata in a Report.
 */
@Builder
@Getter
public class RequestReport {
  String path;

  String method;

  Map<String, Object> parameters;

  String error;
}
