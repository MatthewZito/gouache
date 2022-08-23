package com.github.exbotanical.session.controllers.security;

import java.util.Collections;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationProvider;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.web.authentication.preauth.PreAuthenticatedAuthenticationToken;

import com.github.exbotanical.session.models.UserCredentials;
import com.github.exbotanical.session.services.UserDataService;
import com.github.exbotanical.session.services.UserService;

public class CookiesAuthenticationProvider implements AuthenticationProvider {
  @Autowired
  private UserService userService;

  @Override
  public boolean supports(Class<?> authentication) {
    System.out.println("HIwdXXX");

    return true;
  }


  @Override
  public Authentication authenticate(Authentication authentication) throws AuthenticationException {

    try {
      if (authentication instanceof UsernamePasswordAuthenticationToken) {
        UserCredentials credentials = UserCredentials.builder()
            .username((String) authentication.getPrincipal())
            .password((String) authentication.getCredentials())
            .build();

        if (userService.authenticate(credentials)) {
          return new UsernamePasswordAuthenticationToken(credentials, Collections.emptyList());
        } else if (authentication instanceof PreAuthenticatedAuthenticationToken) {
          return null;
        }

      }

    } catch (Exception e) {
      // throw new AuthenticationException(e) ;
      return null;
    }
  }

}
