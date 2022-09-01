package com.github.exbotanical.resource.middleware;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import com.github.exbotanical.resource.SessionTestUtils;
import com.github.exbotanical.resource.controllers.ResourceController;
import com.github.exbotanical.resource.entities.Session;
import com.github.exbotanical.resource.meta.Constants;
import com.github.exbotanical.resource.services.QueueSenderService;
import com.github.exbotanical.resource.services.ResourceService;
import com.github.exbotanical.resource.services.SessionService;
import java.time.Instant;
import java.util.Date;
import javax.servlet.http.Cookie;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.cloud.aws.messaging.listener.QueueMessageHandler;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;

@WebMvcTest(ResourceController.class)
@DisplayName("Test the AuthInterceptor by replicating all scenarios under which a request is considered unauthorized")
class AuthInterceptorTest {

  @Autowired
  private MockMvc mockMvc;

  @MockBean
  private ResourceService resourceService;

  @MockBean
  private SessionService sessionService;

  @MockBean
  private ResourceController resourceController;

  // No-op mock to prevent connection attempts.
  @MockBean
  private QueueMessageHandler queueMessageHandler;

  // No-op mock to prevent connection attempts.
  @MockBean
  private QueueSenderService queueSenderService;

  @Test
  @DisplayName("Establish a baseline and expect authorized")
  void shouldAuthorizeRequestWithValidSession() throws Exception {
    Cookie sessionCookie = SessionTestUtils.cookie;

    Mockito
        .when(sessionService.getSessionBySessionId(sessionCookie.getValue()))
        .thenReturn(SessionTestUtils.session);

    mockMvc.perform(
        get("/api/resource/a66de382-a9df-4fab-9d34-616e01e3e054")
            .contentType(MediaType.APPLICATION_JSON)
            .cookie(sessionCookie))

        .andExpect(status().isOk());
  }

  @Test
  @DisplayName("Test authorization when the request has no cookie")
  void shouldRejectRequestWithNoCookie() throws Exception {
    mockMvc.perform(
        get("/api/resource/a66de382-a9df-4fab-9d34-616e01e3e054")
            .contentType(MediaType.APPLICATION_JSON))

        .andExpect(status().isUnauthorized())
        .andExpect(jsonPath("$.data").isEmpty())
        .andExpect(jsonPath("$.friendly").value(Constants.E_UNAUTHORIZED))
        .andExpect(jsonPath("$.internal").value(Constants.E_COOKIE_NOT_FOUND));
  }

  @Test
  @DisplayName("Test authorization when the request has a cookie but no session id therein")
  void shouldRejectRequestWithNoSessionId() throws Exception {
    mockMvc.perform(
        get("/api/resource/a66de382-a9df-4fab-9d34-616e01e3e054")
            .contentType(MediaType.APPLICATION_JSON)
            .cookie(new Cookie("gouache_session", null)))

        .andExpect(status().isUnauthorized())
        .andExpect(jsonPath("$.data").isEmpty())
        .andExpect(jsonPath("$.friendly").value(Constants.E_UNAUTHORIZED))
        .andExpect(jsonPath("$.internal").value(Constants.E_SESSION_ID_NOT_FOUND));
  }

  @Test
  @DisplayName("Test authorization when the request session id does not match a known session")
  void shouldRejectRequestWithNoSession() throws Exception {
    Cookie sessionCookie = SessionTestUtils.cookie;

    Mockito
        .when(sessionService.getSessionBySessionId(sessionCookie.getValue()))
        .thenReturn(null);

    mockMvc.perform(
        get("/api/resource/a66de382-a9df-4fab-9d34-616e01e3e054")
            .contentType(MediaType.APPLICATION_JSON)
            .cookie(sessionCookie))

        .andExpect(status().isUnauthorized())
        .andExpect(jsonPath("$.data").isEmpty())
        .andExpect(jsonPath("$.friendly").value(Constants.E_UNAUTHORIZED))
        .andExpect(jsonPath("$.internal").value(
            String.format(Constants.E_SESSION_NOT_FOUND_FMT, sessionCookie.getValue())));
  }

  @Test
  @DisplayName("Test authorization when the request session is expired")
  void shouldRejectRequestWithExpiredSession() throws Exception {
    Cookie sessionCookie = SessionTestUtils.cookie;
    String sid = sessionCookie.getValue();
    String testUsername = "username";
    Date testExpiry = Date.from(Instant.parse("1000-12-31T00:00:00Z"));

    Mockito
        .when(sessionService.getSessionBySessionId(sid))
        .thenReturn(Session
            .builder()
            .username(testUsername)
            .expiry(testExpiry)
            .build());

    mockMvc.perform(
        get("/api/resource/a66de382-a9df-4fab-9d34-616e01e3e054")
            .contentType(MediaType.APPLICATION_JSON)
            .cookie(sessionCookie))

        .andExpect(status().isUnauthorized())
        .andExpect(jsonPath("$.data").isEmpty())
        .andExpect(jsonPath("$.friendly").value(Constants.E_UNAUTHORIZED))
        .andExpect(jsonPath("$.internal").value(
            String.format(Constants.E_SESSION_EXPIRED_FMT, sid, testUsername, testExpiry)));
  }
}
