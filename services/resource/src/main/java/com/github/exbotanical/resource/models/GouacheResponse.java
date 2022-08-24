package com.github.exbotanical.resource.models;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class GouacheResponse {
  String internal;
  String friendly;
  Object data;
  int flags;
}
