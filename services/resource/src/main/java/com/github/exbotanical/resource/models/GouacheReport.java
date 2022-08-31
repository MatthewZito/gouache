package com.github.exbotanical.resource.models;

import lombok.Builder;
import lombok.Getter;

/**
 * A spec-compliant report entity qua the gouache/reporting system.
 */
@Builder
@Getter
public class GouacheReport {
  String name;

  String caller;

  String data;
}
