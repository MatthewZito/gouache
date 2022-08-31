package com.github.exbotanical.resource.models;

import lombok.Builder;
import lombok.Getter;

@Builder
@Getter
public class GouacheReport {
  String name;

  String caller;

  String data;
}
