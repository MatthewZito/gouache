package com.github.exbotanical.resource.controllers.advice;

import com.github.exbotanical.resource.annotations.IgnoreGouacheResponseBinding;
import com.github.exbotanical.resource.controllers.ResourceController;
import com.github.exbotanical.resource.models.GouacheResponseModel;
import org.springframework.core.MethodParameter;
import org.springframework.http.MediaType;
import org.springframework.http.converter.HttpMessageConverter;
import org.springframework.http.server.ServerHttpRequest;
import org.springframework.http.server.ServerHttpResponse;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.servlet.mvc.method.annotation.ResponseBodyAdvice;

import java.util.ArrayList;
import java.util.Arrays;

/**
 * Global response handler. Normalizes all responses to a GouacheResponse.
 */
@ControllerAdvice
public class GouacheResponseHandler implements ResponseBodyAdvice<Object> {
  private final ArrayList<Class<?>> supportedControllers = new ArrayList<>(
    Arrays.asList(ResourceController.class));

  @Override
  public boolean supports(MethodParameter returnType,
                          Class<? extends HttpMessageConverter<?>> converterType) {

    return supportedControllers.contains(returnType.getContainingClass());
  }

  @Override
  public Object beforeBodyWrite(Object body, MethodParameter returnType,
                                MediaType selectedContentType, Class<? extends HttpMessageConverter<?>> selectedConverterType,
                                ServerHttpRequest request, ServerHttpResponse response) {

    // If the controller handler bears the @IgnoreGouacheResponseBinding annotation, return the body
    // as-is.
    if (returnType.getContainingClass().isAnnotationPresent(IgnoreGouacheResponseBinding.class)) {
      return body;
    }

    // If the response body is already normalized as a GouacheResponse, return as-is.
    if (returnType.getParameterType() == GouacheResponseModel.class) {
      return body;
    }

    // Normalize into a GouacheResponse.
    return new GouacheResponseModel(null, null, body, 0);
  }
}
