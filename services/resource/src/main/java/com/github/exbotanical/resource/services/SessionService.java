package com.github.exbotanical.resource.services;

import com.github.exbotanical.resource.entities.Session;

/**
 * Service for Session data.
 */
public interface SessionService {

  /**
   * Retrieve a Session from cache by its session id.
   *
   * @param sid A unique Session id identifying the Session to retrieve.
   * @return A Session, or null if not found.
   */
  Session getSessionBySessionId(String sid);
}
