package com.github.exbotanical.resource.middleware;

import com.github.exbotanical.resource.entities.Session;
import com.github.exbotanical.resource.errors.UnauthorizedException;
import com.github.exbotanical.resource.meta.Constants;
import com.github.exbotanical.resource.models.ReportName;
import com.github.exbotanical.resource.services.QueueSenderService;
import com.github.exbotanical.resource.services.SessionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.util.WebUtils;

import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.util.Date;

/**
 * An auth interceptor - used to verify user access and authorization.
 */
@Component
public class AuthInterceptor implements HandlerInterceptor {

  @Autowired
  private SessionService sessionService;

  @Autowired
  private QueueSenderService queueSenderService;

  @Value("${app.cookie_name}")
  String cookieName;

  @Override
  public boolean preHandle(HttpServletRequest req, HttpServletResponse res, Object handler)
    throws Exception {

    if ("OPTIONS".equals(req.getMethod())) {
      return HandlerInterceptor.super.preHandle(req, res, handler);
    }

    // @todo add ignore auth annotation
    Cookie cookie = WebUtils.getCookie(req, cookieName);

    if (cookie == null) {

      throw deriveUnauthorizedExceptionAndSendReport(Constants.E_COOKIE_NOT_FOUND);
    }

    String sid = cookie.getValue();

    if (sid == null) {
      throw deriveUnauthorizedExceptionAndSendReport(Constants.E_SESSION_ID_NOT_FOUND);
    }

    Session session = null;
    try {
      session = sessionService.getSessionBySessionId(sid);
    } catch (Exception ex) {
      throw deriveUnauthorizedExceptionAndSendReport(String.format(Constants.E_SESSION_NOT_FOUND_FMT, sid));
    }

    if (session == null) {
      throw deriveUnauthorizedExceptionAndSendReport(String.format(Constants.E_SESSION_NOT_FOUND_FMT, sid));
    }

    if (session.expiry.before(new Date())) {
      throw deriveUnauthorizedExceptionAndSendReport(String.format(Constants.E_SESSION_EXPIRED_FMT,
        sid, session.username, session.expiry));
    }

    return HandlerInterceptor.super.preHandle(req, res, handler);
  }

  private UnauthorizedException deriveUnauthorizedExceptionAndSendReport(String message) {
    queueSenderService.sendMessage(message, ReportName.UNAUTHORIZED_ACCESS);

    return new UnauthorizedException(message);
  }
}
