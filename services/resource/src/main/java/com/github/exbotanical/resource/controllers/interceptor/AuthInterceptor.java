package com.github.exbotanical.resource.controllers.interceptor;

import java.util.Date;

import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.util.WebUtils;

import com.github.exbotanical.resource.entities.Session;
import com.github.exbotanical.resource.errors.UnauthorizedException;
import com.github.exbotanical.resource.services.SessionService;

/**
 * An auth interceptor - used to verify user access and authorization.
 */
@Component
public class AuthInterceptor implements HandlerInterceptor {

  @Autowired
  private SessionService sessionService;

  @Value("${app.cookie_name}")
  String cookieName;

  @Override
  public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler)
      throws Exception {

    // @todo add ignore auth annotation
    Cookie cookie = WebUtils.getCookie(request, cookieName);

    String sid = cookie.getValue();

    if (sid == null) {
      throw new UnauthorizedException("No session Id found");
    }

    Session session = sessionService.getSessionBySessionId(sid);

    if (session == null) {
      throw new UnauthorizedException(String.format("No session found for id %s", sid));
    }

    if (session.expiry.before(new Date())) {
      throw new UnauthorizedException(String.format("Session with id %s for user %s expired on %s",
          sid, session.username, session.expiry));
    }

    return HandlerInterceptor.super.preHandle(request, response, handler);
  }

  private static String newRecord(String key, Object value) {
    return "\t" + key + ": " + value + ",\n";
  }
}
