package com.github.exbotanical.session.services;

import com.github.exbotanical.session.entities.User;
import com.github.exbotanical.session.errors.UserInvariantViolationException;
import com.github.exbotanical.session.errors.UserNotFoundException;
import com.github.exbotanical.session.models.UserModel;
import com.github.exbotanical.session.repositories.UserRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UserDataService implements UserService {
  @Autowired
  private UserRepository userRepository;

  @Override
  public UserModel createUser(UserModel userModel) {
    return userRepository.createUser(userModel);
  }

  @Override
  public User getUserByUsername(String username)
      throws UserNotFoundException, UserInvariantViolationException {
    return userRepository.getUserByUsername(username);
  }
}
