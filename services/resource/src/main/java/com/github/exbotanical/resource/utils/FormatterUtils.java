package com.github.exbotanical.resource.utils;

import java.util.stream.Collectors;

import org.springframework.validation.BindingResult;

public class FormatterUtils {
  public static String formatValidationErrors(BindingResult result) {
    return result.getFieldErrors().stream()
        .map(f -> String.format("%s.%s %s", f.getObjectName(), f.getField(),
            f.getDefaultMessage()))
        .collect(Collectors.joining(", "));
  }
}
