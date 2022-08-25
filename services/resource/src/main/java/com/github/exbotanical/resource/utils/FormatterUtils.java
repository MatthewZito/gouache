package com.github.exbotanical.resource.utils;

import java.util.stream.Collectors;

import org.springframework.validation.BindingResult;

/**
 * Shared utilities for formatting and normalization.
 */
public class FormatterUtils {

  /**
   * Extract from and format all field validation errors from a given BindingResult object `result`
   * and return as a single comma-delimited string, where each delimited substring contains the
   * format <objectName>.<fieldName> <errorMessage>.
   *
   * @param result A BindingResult as returned by Spring via the @Valid annotation.
   * @return A comma-delimited string of validation error messages.
   */
  public static String formatValidationErrors(BindingResult result) {
    return result.getFieldErrors().stream()
        .map(f -> String.format("%s.%s %s", f.getObjectName(), f.getField(),
            f.getDefaultMessage()))
        .collect(Collectors.joining(", "));
  }
}
