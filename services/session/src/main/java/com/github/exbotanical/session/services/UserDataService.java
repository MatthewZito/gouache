package com.github.exbotanical.session.services;

import com.github.exbotanical.session.entities.User;
import com.github.exbotanical.session.errors.UserInvariantViolationException;
import com.github.exbotanical.session.errors.UserNotFoundException;
import com.github.exbotanical.session.models.UserCredentials;
import com.github.exbotanical.session.repositories.UserRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
public class UserDataService implements UserService {
  @Autowired
  private UserRepository userRepository;

  @Autowired
  private PasswordEncoder passwordEncoder;

  @Override
  public User createUser(UserCredentials credentials) {
    User user = User.builder()
        .passwordHash(passwordEncoder.encode(credentials.getPassword()))
        .username(credentials.getUsername())
        .build();

    return userRepository.createUser(user);
  }

  @Override
  public User getUserByUsername(String username)
      throws UserNotFoundException, UserInvariantViolationException {
    return userRepository.getUserByUsername(username);
  }

  @Override
  public boolean authenticate(UserCredentials credentials)
      throws UserNotFoundException, UserInvariantViolationException {
    User user = getUserByUsername(credentials.getUsername());

    return passwordEncoder.matches(credentials.getPassword(), user.getPasswordHash());
  }
}
