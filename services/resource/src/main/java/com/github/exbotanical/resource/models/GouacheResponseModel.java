package com.github.exbotanical.resource.models;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * GouacheResponse represents a normalized, formatted response object.
 */
@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class GouacheResponseModel {
  String internal;

  String friendly;

  Object data;

  int flags;
}
