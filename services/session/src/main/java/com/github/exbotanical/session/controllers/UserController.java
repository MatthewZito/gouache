package com.github.exbotanical.session.controllers;

import javax.validation.Valid;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.web.server.Cookie;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.web.bind.annotation.*;

import com.github.exbotanical.session.config.filters.CookieValidatorFilter;
import com.github.exbotanical.session.entities.User;
import com.github.exbotanical.session.models.UserCredentials;
import com.github.exbotanical.session.models.UserSessionResponse;
import com.github.exbotanical.session.services.UserService;


@RestController
@RequestMapping("/user")
public class UserController {

  @Value("${app.locale}")
  private String locale;

  @Autowired
  private UserService userService;

  private final Logger LOGGER = LoggerFactory.getLogger(UserController.class);

  @PostMapping("/register")
  public UserSessionResponse register(@Valid @RequestBody UserCredentials credentials)
      throws Exception {
    System.out.println("DEFAULT HANDLER");
    User updatedModel = userService.createUser(credentials);

    UserSessionResponse res = UserSessionResponse.builder()
        .username(updatedModel.getUsername())
        .exp(3600 /* @todo */)
        .build();

    return res;
  }

  @PostMapping("/login")
  public UserSessionResponse login(@AuthenticationPrincipal UserCredentials credentials)
      throws Exception {

    User user = userService.getUserByUsername(credentials.getUsername());


    Cookie authCookie =
        new Cookie(CookieValidatorFilter.COOKIE_NAME, authenticationService.createToken(user));


    authCookie.setHttpOnly(true);
    authCookie.setSecure(true);
    authCookie.setMaxAge((int) Duration.of(1, ChronoUnit.DAYS).toSeconds());
    authCookie.setPath("/");

    // @todo
    UserSessionResponse res = UserSessionResponse.builder()
        .username(user.getUsername())
        .exp(3600 /* @todo */)
        .build();

    return res;
  }

  @PostMapping("/logout")
  public void logout() {
    System.out.println("DEFAULT HANDLER");

  }

  @PostMapping("/renew")
  public UserSessionResponse renewSession() {
    return null;
  }
}
