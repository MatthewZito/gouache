package com.github.exbotanical.session.services;

import com.github.exbotanical.session.entities.User;
import com.github.exbotanical.session.errors.UserInvariantViolationException;
import com.github.exbotanical.session.errors.UserNotFoundException;
import com.github.exbotanical.session.models.UserCredentials;

public interface UserService {
  public User createUser(UserCredentials user);

  public boolean authenticate(UserCredentials credentials)
      throws UserNotFoundException, UserInvariantViolationException;

  public User getUserByUsername(String username)
      throws UserNotFoundException, UserInvariantViolationException;
}
