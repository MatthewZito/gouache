package com.github.exbotanical.session.config.filters;


import java.io.IOException;
import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.exbotanical.session.entities.User;

import org.springframework.http.HttpMethod;
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
        User u = MAPPER.readValue(httpServletRequest.getInputStream(), User.class);
        System.out.println(u);

      } catch (Exception e) {
        System.out.println("NOO");
        // TODO: handle exception
      }
      // SecurityContextHolder.getContext().setAuthentication(
      // new UsernamePasswordAuthenticationToken(credentialsDto.getLogin(),
      // credentialsDto.getPassword()));

      System.out.println("NOO");

    }

    filterChain.doFilter(httpServletRequest, httpServletResponse);
  }
}
