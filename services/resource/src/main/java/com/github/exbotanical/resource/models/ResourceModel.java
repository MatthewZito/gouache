package com.github.exbotanical.resource.models;

import java.util.List;
import javax.validation.constraints.NotNull;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.NoArgsConstructor;

/**
 * ResourceModel represents the necessary input data a client must provide to create a new
 * `Resource`.
 */
@Getter
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class ResourceModel {
  @NotNull
  private String title;

  @NotNull
  private List<String> tags;
}
