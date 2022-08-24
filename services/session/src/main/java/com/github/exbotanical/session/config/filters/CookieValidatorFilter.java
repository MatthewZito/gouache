package com.github.exbotanical.session.config.filters;

import java.io.IOException;
import java.util.Optional;
import java.util.stream.Stream;
import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.web.authentication.preauth.PreAuthenticatedAuthenticationToken;
import org.springframework.web.filter.OncePerRequestFilter;

public class CookieValidatorFilter extends OncePerRequestFilter {

  public static final String COOKIE_NAME = "gouache_session";

  @Override
  protected void doFilterInternal(HttpServletRequest httpServletRequest,
      HttpServletResponse httpServletResponse,
      FilterChain filterChain) throws ServletException, IOException {
    System.out.println("COOKIE FILTER");

    Optional<Cookie> cookieAuth =
        Stream.of(Optional.ofNullable(httpServletRequest.getCookies()).orElse(new Cookie[0]))
            .filter(cookie -> COOKIE_NAME.equals(cookie.getName()))
            .findFirst();

    if (cookieAuth.isPresent()) {
      SecurityContextHolder.getContext().setAuthentication(
          new PreAuthenticatedAuthenticationToken(cookieAuth.get().getValue(), null));
    }

    filterChain.doFilter(httpServletRequest, httpServletResponse);
  }
}
