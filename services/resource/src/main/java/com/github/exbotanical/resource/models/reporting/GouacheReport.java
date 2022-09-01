package com.github.exbotanical.resource.models.reporting;

import lombok.Builder;
import lombok.Getter;

/**
 * A spec-compliant report entity qua the gouache/reporting system.
 */
@Builder
@Getter
public class GouacheReport {
  GouacheReportName name;

  final String caller = "gouache/resource";

  Object data;
}
