package com.github.exbotanical.resource.controllers;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.delete;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.patch;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import com.amazonaws.util.json.Jackson;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.exbotanical.resource.SessionTestUtils;
import com.github.exbotanical.resource.entities.Resource;
import com.github.exbotanical.resource.meta.Constants;
import com.github.exbotanical.resource.models.ResourceModel;
import com.github.exbotanical.resource.services.QueueSenderService;
import com.github.exbotanical.resource.services.ResourceService;
import com.github.exbotanical.resource.services.SessionService;
import java.time.Instant;
import java.util.Arrays;
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

@WebMvcTest(ResourceController.class)
@DisplayName("Test the ResourceController and evaluate its formatted responses")
class ResourceControllerTest {
  @Autowired
  ObjectMapper objectMapper;

  @Autowired
  private MockMvc mockMvc;

  @MockBean
  private ResourceService resourceService;

  @MockBean
  private SessionService sessionService;

  // No-op mock to prevent connection attempts.
  @MockBean
  private QueueMessageHandler queueMessageHandler;

  // No-op mock to prevent connection attempts.
  @MockBean
  private QueueSenderService queueSenderService;

  private Resource testResource;

  @BeforeEach
  void setUp() {
    Mockito
        .when(sessionService.getSessionBySessionId(ArgumentMatchers.anyString()))
        .thenReturn(SessionTestUtils.session);

    testResource = Resource.builder()
        .id("a66de382-a9df-4fab-9d34-616e01e3e054")
        .title("title")
        .tags(Arrays.asList("art", "music"))
        .createdAt(Instant.now())
        .updatedAt(Instant.now())
        .build();
  }

  @Test
  @DisplayName("Create a resource successfully")
  void createResourceSuccess() throws Exception {
    ResourceModel inputModel = ResourceModel.builder()
        .title("title")
        .tags(Arrays.asList("art", "music"))
        .build();

    Mockito.when(resourceService.createResource(
        ArgumentMatchers.any())).thenReturn(testResource);

    mockMvc.perform(
        post("/resource")
            .contentType(MediaType.APPLICATION_JSON)
            .content(Jackson.toJsonString(inputModel))
            .cookie(SessionTestUtils.cookie))
        .andExpect(status().isCreated())
        .andExpect(jsonPath("$.data.id").value(testResource.getId()))
        .andExpect(jsonPath("$.data.title").value(testResource.getTitle())

        );
  }

  @Test
  @DisplayName("Attempt to create a resource with an invalid input model")
  void createResourceInvalidInput() throws Exception {
    ResourceModel inputModel = ResourceModel.builder()
        .title("title")
        .tags(Arrays.asList("art", "music"))
        .build();

    Mockito.when(resourceService.createResource(inputModel)).thenReturn(testResource);

    mockMvc.perform(
        post("/resource")
            .contentType(MediaType.APPLICATION_JSON)
            .content("{\"titl ez\": \"title\", \"tagds\": [\"art\",\"music\"] }")
            .cookie(SessionTestUtils.cookie))
        .andExpect(status().isBadRequest())
        .andExpect(jsonPath("$.internal").isString())
        .andExpect(jsonPath("$.friendly").value(Constants.E_INVALID_INPUT))
        .andExpect(jsonPath("$.data").isEmpty());
  }

  @Test
  @DisplayName("Retrieve a resource by ID")
  void getResourceById() throws Exception {
    String testId = "a66de382-a9df-4fab-9d34-616e01e3e054";

    Mockito.when(resourceService.getResourceById(testId)).thenReturn(testResource);

    mockMvc.perform(
        get(String.format("/resource/%s", testId))

            .cookie(SessionTestUtils.cookie))
        .andExpect(status().isOk())
        .andExpect(jsonPath("$.data.id").value(testResource.getId()))
        .andExpect(jsonPath("$.data.title").value(testResource.getTitle())

        );
  }

  @Test
  @DisplayName("Attempt to retrieve a resource that does not exist")
  void getResourceByIdNotFound() throws Exception {
    String testId = "a66de382-a9df-4fab-9d34-616e01e3e054";

    Mockito.when(resourceService.getResourceById(testId)).thenReturn(testResource);

    mockMvc.perform(
        get(String.format("/resource/%s", testId + "1"))

            .cookie(SessionTestUtils.cookie))
        .andExpect(status().isNotFound())
        .andExpect(jsonPath("$.data").isEmpty());
  }

  @Test
  @DisplayName("Update a resource by ID")
  void updateResourceById() throws Exception {
    ResourceModel inputModel = ResourceModel.builder()
        .tags(Arrays.asList("test"))
        .title("test title")
        .build();

    mockMvc.perform(
        patch(String.format("/resource/%s", testResource.getId()))
            .contentType(MediaType.APPLICATION_JSON)
            .content(Jackson.toJsonString(inputModel))

            .cookie(SessionTestUtils.cookie))
        .andExpect(status().isOk())
        .andExpect(jsonPath("$.data").isEmpty());
  }

  @Test
  @DisplayName("Attempt to update a resource by ID with an invalid input model")
  void updateResourceByIdInvalidInput() throws Exception {
    ResourceModel inputModel = ResourceModel.builder().build();

    mockMvc.perform(
        patch(String.format("/resource/%s", testResource.getId()))
            .contentType(MediaType.APPLICATION_JSON)
            .content(Jackson.toJsonString(inputModel))
            .cookie(SessionTestUtils.cookie))
        .andExpect(status().isBadRequest())
        .andExpect(jsonPath("$.internal").isString())
        .andExpect(jsonPath("$.friendly").value(Constants.E_INVALID_INPUT))
        .andExpect(jsonPath("$.data").isEmpty());
  }

  @Test
  @DisplayName("Delete a resource by ID")
  void deleteResourceById() throws Exception {
    String testId = "a66de382-a9df-4fab-9d34-616e01e3e054";

    mockMvc.perform(
        delete(String.format("/resource/%s", testId))
            .cookie(SessionTestUtils.cookie))
        .andExpect(status().isOk())
        .andExpect(jsonPath("$.data").isEmpty())
        .andExpect(jsonPath("$.friendly").isEmpty())
        .andExpect(jsonPath("$.internal").isEmpty());
  }

  @Test
  @DisplayName("Erroneous delete resource by ID")
  void deleteResourceByIdError() throws Exception {
    String testId = "a66de382-a9df-4fab-9d34-616e01e3e055";

    Mockito.doThrow(new RuntimeException("test")).when(resourceService).deleteResourceById(testId);

    mockMvc.perform(
        delete(String.format("/resource/%s", testId))
            .cookie(SessionTestUtils.cookie))
        .andExpect(status().isBadRequest())
        .andExpect(jsonPath("$.data").isEmpty())
        .andExpect(jsonPath("$.friendly").value(
            String.format(Constants.E_RESOURCE_DELETE_FMT, testId)))
        .andExpect(jsonPath("$.internal").value("test"));
  }
}
