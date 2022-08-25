package com.github.exbotanical.resource.services;

import com.github.exbotanical.resource.entities.Session;
import com.github.exbotanical.resource.repositories.SessionRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

/**
 * An implementation of SessionService.
 */
@Service
public class SessionDataService implements SessionService {

  @Autowired
  private SessionRepository sessionRepository;

  @Override
  public Session getSessionBySessionId(String sid) {
    return sessionRepository.getById(sid);
  }
}
