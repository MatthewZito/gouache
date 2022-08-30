package com.github.exbotanical.resource.models.logs;

import lombok.Getter;
import lombok.experimental.SuperBuilder;

import java.util.Map;

@SuperBuilder
@Getter
public class RequestLog {
  String path;

  String method;

  Map<String, Object> parameters;
}

