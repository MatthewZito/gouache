package com.github.exbotanical.session.controllers;

import javax.validation.Valid;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.*;

import com.github.exbotanical.session.entities.User;
import com.github.exbotanical.session.models.UserModel;
import com.github.exbotanical.session.models.UserSessionResponse;
import com.github.exbotanical.session.services.UserService;


@RestController
public class UserController {

  @Value("${app.locale}")
  private String locale;

  @Autowired
  private UserService userService;

  private final Logger LOGGER = LoggerFactory.getLogger(UserController.class);

  @PostMapping("/user/register")
  public UserSessionResponse register(@Valid @RequestBody UserModel userModel)
      throws Exception {
    UserModel updatedModel = userService.createUser(userModel);

    UserSessionResponse res = UserSessionResponse.builder()
        .username(updatedModel.getUsername())
        .exp(3600 /* @todo */)
        .build();

    return res;
  }

  @PostMapping("/user/login")
  public UserSessionResponse login(@Valid @RequestBody UserModel userModel) throws Exception {
    User user = userService.getUserByUsername(userModel.username);

    // @todo
    UserSessionResponse res = UserSessionResponse.builder()
        .username(user.getUsername())
        .exp(3600 /* @todo */)
        .build();

    return res;
  }

  @PostMapping("/user/logout")
  public void logout() {

  }

  @PostMapping("/user/renew")
  public UserSessionResponse renewSession() {
    return null;
  }
}
