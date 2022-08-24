package com.github.exbotanical.session.config.filters;


import java.io.IOException;
import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.exbotanical.session.entities.User;
import com.github.exbotanical.session.models.UserCredentials;

import org.springframework.http.HttpMethod;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.filter.OncePerRequestFilter;

public class AuthenticationFilter extends OncePerRequestFilter {

  private static final ObjectMapper MAPPER = new ObjectMapper();

  @Override
  protected void doFilterInternal(
      HttpServletRequest httpServletRequest,
      HttpServletResponse httpServletResponse,
      FilterChain filterChain) throws ServletException, IOException {

    if ("/user/login".equals(httpServletRequest.getServletPath())
        && HttpMethod.POST.matches(httpServletRequest.getMethod())) {

      try {
        UserCredentials credentials =
            MAPPER.readValue(httpServletRequest.getInputStream(), UserCredentials.class);


        SecurityContextHolder.getContext().setAuthentication(
            new UsernamePasswordAuthenticationToken(credentials.getUsername(),
                credentials.getPassword()));
      } catch (Exception e) {
        System.out.println(e);
      }
      // SecurityContextHolder.getContext().setAuthentication(
      // new UsernamePasswordAuthenticationToken(credentialsDto.getLogin(),
      // credentialsDto.getPassword()));
    }

    filterChain.doFilter(httpServletRequest, httpServletResponse);
  }
}
