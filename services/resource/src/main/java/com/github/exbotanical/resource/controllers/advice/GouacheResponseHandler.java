package com.github.exbotanical.resource.controllers.advice;

import com.github.exbotanical.resource.annotations.IgnoreGouacheResponseBinding;
import com.github.exbotanical.resource.controllers.ResourceController;
import com.github.exbotanical.resource.models.GouacheResponse;
import org.springframework.core.MethodParameter;
import org.springframework.http.MediaType;
import org.springframework.http.converter.HttpMessageConverter;
import org.springframework.http.server.ServerHttpRequest;
import org.springframework.http.server.ServerHttpResponse;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.servlet.mvc.method.annotation.ResponseBodyAdvice;

import java.util.ArrayList;
import java.util.Arrays;

@ControllerAdvice
public class GouacheResponseHandler implements ResponseBodyAdvice<Object> {
  private final ArrayList<Class<?>> SUPPORTED_CONTROLLERS = new ArrayList<>(
    Arrays.asList(ResourceController.class));

  @Override
  public boolean supports(MethodParameter returnType,
                          Class<? extends HttpMessageConverter<?>> converterType) {
    return SUPPORTED_CONTROLLERS.contains(returnType.getContainingClass());
  }

  @Override
  public Object beforeBodyWrite(Object body, MethodParameter returnType,
                                MediaType selectedContentType, Class<? extends HttpMessageConverter<?>> selectedConverterType,
                                ServerHttpRequest request, ServerHttpResponse response) {
    System.out.println("HERE _+" + body);

    // If the controller handler bears the @IgnoreGouacheResponseBinding annotation, return the body as-is.
    if (returnType.getContainingClass().isAnnotationPresent(IgnoreGouacheResponseBinding.class)) {
      return body;
    }

    // If the response body is already normalized as a GouacheResponse, return as-is.
    if (returnType.getParameterType() == GouacheResponse.class) {
      return body;
    }

    // Normalize into a GouacheResponse.
    return new GouacheResponse(null, null, body, 0);
  }
}
