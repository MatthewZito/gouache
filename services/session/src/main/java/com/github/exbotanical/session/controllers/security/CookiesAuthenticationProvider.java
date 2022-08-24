package com.github.exbotanical.session.controllers.security;

import java.util.Collections;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationProvider;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.web.authentication.preauth.PreAuthenticatedAuthenticationToken;
import org.springframework.stereotype.Component;

import com.github.exbotanical.session.models.UserCredentials;
import com.github.exbotanical.session.services.UserService;

@Component
public class CookiesAuthenticationProvider implements AuthenticationProvider {
  @Autowired
  private UserService userService;

  @Override
  public boolean supports(Class<?> authentication) {

    return true;
  }

  @Override
  public Authentication authenticate(Authentication authentication) throws AuthenticationException {
    System.out.println("auth invoke");

    try {
      if (authentication instanceof UsernamePasswordAuthenticationToken) {
        UserCredentials credentials = UserCredentials.builder()
            .username((String) authentication.getPrincipal())
            .password((String) authentication.getCredentials())
            .build();

        if (userService.authenticate(credentials)) {
          return new UsernamePasswordAuthenticationToken(credentials, Collections.emptyList());
        } else if (authentication instanceof PreAuthenticatedAuthenticationToken) {
          // verify user's session
          return null;
        }

      }

    } catch (Exception e) {
      System.out.println(e);

      // throw new AuthenticationException(e) ;
    }

    return null;
  }

}
