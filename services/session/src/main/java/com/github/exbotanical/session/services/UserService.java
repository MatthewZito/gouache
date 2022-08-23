package com.github.exbotanical.session.services;

import com.github.exbotanical.session.entities.User;
import com.github.exbotanical.session.errors.UserInvariantViolationException;
import com.github.exbotanical.session.errors.UserNotFoundException;
import com.github.exbotanical.session.models.UserModel;

public interface UserService {
  public UserModel createUser(UserModel userModel);

  public User getUserByUsername(String username)
      throws UserNotFoundException, UserInvariantViolationException;
}
