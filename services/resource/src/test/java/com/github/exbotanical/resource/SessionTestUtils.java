package com.github.exbotanical.resource;

import java.time.Instant;
import java.util.Date;

import javax.servlet.http.Cookie;

import com.github.exbotanical.resource.entities.Session;

public class SessionTestUtils {
    public static final Cookie cookie = new Cookie("gouache_session", "test");

    public static final Session session = Session.builder().username("username").expiry(Date.from(Instant.parse("9999-12-31T00:00:00Z"))).build();
}
