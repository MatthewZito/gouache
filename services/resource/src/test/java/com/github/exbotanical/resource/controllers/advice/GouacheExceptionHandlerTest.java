package com.github.exbotanical.resource.controllers.advice;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.delete;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import com.github.exbotanical.resource.SessionTestUtils;
import com.github.exbotanical.resource.controllers.ResourceController;
import com.github.exbotanical.resource.errors.GouacheException;
import com.github.exbotanical.resource.errors.OperationFailedException;
import com.github.exbotanical.resource.meta.Constants;
import com.github.exbotanical.resource.services.QueueSenderService;
import com.github.exbotanical.resource.services.ResourceService;
import com.github.exbotanical.resource.services.SessionService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentMatchers;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.cloud.aws.messaging.listener.QueueMessageHandler;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.web.client.HttpServerErrorException.InternalServerError;

@WebMvcTest(ResourceController.class)
@DisplayName("Test the ResourceController and evaluate its formatted responses")
class GouacheExceptionHandlerTest {

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

  @BeforeEach
  void setUp() {
    Mockito
        .when(sessionService.getSessionBySessionId(ArgumentMatchers.anyString()))
        .thenReturn(SessionTestUtils.session);
  }

  @Test
  @DisplayName("Trigger the exception handler advice with a GouacheException")
  void shouldHandleGouacheException() throws Exception {
    Mockito
        .when(resourceController.deleteResourceById(ArgumentMatchers.any()))
        .thenThrow(new GouacheException("x", "y"));

    mockMvc.perform(
        delete("/resource/a66de382-a9df-4fab-9d34-616e01e3e054")
            .contentType(MediaType.APPLICATION_JSON)
            .cookie(SessionTestUtils.cookie))

        .andExpect(status().isBadRequest())
        .andExpect(jsonPath("$.data").isEmpty())
        .andExpect(jsonPath("$.friendly").value("x"))
        .andExpect(jsonPath("$.internal").value("y"));
  }

  @Test
  @DisplayName("Trigger the exception handler advice with a GouacheException subclass")
  void shouldHandleGouacheExceptionSubclass() throws Exception {
    Mockito
        .when(resourceController.deleteResourceById(ArgumentMatchers.any()))
        .thenThrow(new OperationFailedException("x", "y"));

    mockMvc.perform(
        delete("/resource/a66de382-a9df-4fab-9d34-616e01e3e054")
            .contentType(MediaType.APPLICATION_JSON)
            .cookie(SessionTestUtils.cookie))

        .andExpect(status().isBadRequest())
        .andExpect(jsonPath("$.data").isEmpty())
        .andExpect(jsonPath("$.friendly").value("x"))
        .andExpect(jsonPath("$.internal").value("y"));
  }

  @Test
  @DisplayName("Trigger the exception handler advice with a generic Exception")
  void shouldHandleException() throws Exception {
    Mockito
        .when(resourceController.getResourceById(ArgumentMatchers.any()))
        .thenThrow(new RuntimeException("x"));

    mockMvc.perform(
        get("/resource/a66de382-a9df-4fab-9d34-616e01e3e054")
            .contentType(MediaType.APPLICATION_JSON)
            .cookie(SessionTestUtils.cookie))

        .andExpect(status().isBadRequest())
        .andExpect(jsonPath("$.data").isEmpty())
        .andExpect(jsonPath("$.friendly").value(Constants.E_GENERIC))
        .andExpect(jsonPath("$.internal").value("x"));
  }

  @Test
  @DisplayName("Trigger the exception handler advice with a non-existent route")
  void shouldHandleRouteNotFound() throws Exception {
    String reqPath = "/resourcx";

    mockMvc.perform(
        get(reqPath)
            .contentType(MediaType.APPLICATION_JSON)
            .cookie(SessionTestUtils.cookie))

        .andExpect(status().isNotFound())
        .andExpect(jsonPath("$.data").isEmpty())
        .andExpect(
            jsonPath("$.friendly").value(String.format(Constants.E_ROUTE_NOT_FOUND_FMT, reqPath)))
        .andExpect(jsonPath("$.internal").isString());
  }

  @Test
  @DisplayName("Trigger the exception handler advice with an existing route but invalid method")
  void shouldHandleMethodNotAllowed() throws Exception {
    String reqPath = "/resource/a66de382-a9df-4fab-9d34-616e01e3e054";

    mockMvc.perform(
        post(reqPath)
            .contentType(MediaType.APPLICATION_JSON)
            .cookie(SessionTestUtils.cookie))

        .andExpect(status().isMethodNotAllowed())
        .andExpect(jsonPath("$.data").isEmpty())
        .andExpect(jsonPath("$.friendly")
            .value(String.format(Constants.E_METHOD_NOT_ALLOWED_FMT, "POST", reqPath)))
        .andExpect(jsonPath("$.internal").isString());
  }

  @Test
  @DisplayName("Trigger the exception handler advice with an existing route but invalid method")
  void shouldHandleInternalServerError() throws Exception {
    Mockito
        .when(resourceController.deleteResourceById(ArgumentMatchers.any()))
        .thenThrow(InternalServerError.class);

    mockMvc.perform(
        delete("/resource/a66de382-a9df-4fab-9d34-616e01e3e054")
            .contentType(MediaType.APPLICATION_JSON)
            .cookie(SessionTestUtils.cookie))

        .andExpect(status().isInternalServerError())
        .andExpect(jsonPath("$.data").isEmpty())
        .andExpect(jsonPath("$.friendly").value(Constants.E_INTERNAL_SERVER_ERROR))
        // The InternalServerError constructor is private and we are therefore unable to explicitly
        // set a message.
        .andExpect(jsonPath("$.internal").isEmpty());
  }
}
