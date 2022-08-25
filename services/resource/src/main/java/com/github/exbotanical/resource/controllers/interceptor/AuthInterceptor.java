package com.github.exbotanical.resource.controllers.interceptor;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;

/**
 * An auth interceptor - used to verify user access and authorization.
 */
@Component
public class AuthInterceptor implements HandlerInterceptor {

  @Override
  public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler)
      throws Exception {

    System.out.println("{\n" + "\request.getSession():" + request.getSession() + ",\n" + "}");
    return HandlerInterceptor.super.preHandle(request, response, handler);
  }
}
