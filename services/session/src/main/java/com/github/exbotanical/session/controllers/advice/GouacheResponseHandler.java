package com.github.exbotanical.session.controllers.advice;

import java.util.ArrayList;
import java.util.Arrays;

import org.springframework.core.MethodParameter;
import org.springframework.http.MediaType;
import org.springframework.http.converter.HttpMessageConverter;
import org.springframework.http.server.ServerHttpRequest;
import org.springframework.http.server.ServerHttpResponse;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.servlet.mvc.method.annotation.ResponseBodyAdvice;

import com.github.exbotanical.session.annotations.IgnoreGouacheResponseBinding;
import com.github.exbotanical.session.controllers.UserController;
import com.github.exbotanical.session.models.GouacheResponse;


@ControllerAdvice
public class GouacheResponseHandler implements ResponseBodyAdvice<Object> {
  private final ArrayList<Class<?>> SUPPORTED_CONTROLLERS = new ArrayList<>(
      Arrays.asList(UserController.class));

  @Override
  public boolean supports(MethodParameter returnType,
      Class<? extends HttpMessageConverter<?>> converterType) {
    return SUPPORTED_CONTROLLERS.contains(returnType.getContainingClass());
  }

  @Override
  public Object beforeBodyWrite(Object body, MethodParameter returnType,
      MediaType selectedContentType, Class<? extends HttpMessageConverter<?>> selectedConverterType,
      ServerHttpRequest request, ServerHttpResponse response) {

    if (returnType.getContainingClass().isAnnotationPresent(IgnoreGouacheResponseBinding.class)) {
      return body;
    }

    if (returnType.getParameterType() == GouacheResponse.class) {
      return body;
    }

    return new GouacheResponse(null, null, body, 0);
  }
}
