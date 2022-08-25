package com.github.exbotanical.resource.controllers;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.delete;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.patch;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import com.amazonaws.util.json.Jackson;
import com.github.exbotanical.resource.DynamoTestUtils;
import com.github.exbotanical.resource.entities.Resource;
import com.github.exbotanical.resource.models.ResourceModel;
import com.github.exbotanical.resource.services.ResourceService;
import java.util.Arrays;
import java.util.Date;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentMatchers;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;

@WebMvcTest(ResourceController.class)
@DisplayName("Test the ResourceController and evaluate its formatted responses")
class ResourceControllerTest {

  @Autowired
  private MockMvc mockMvc;

  @MockBean
  private ResourceService resourceService;

  private Resource testResource;

  @BeforeAll
  public static void setupDynamo() {
    DynamoTestUtils.setupDynamo();
  }

  @AfterAll
  public static void teardownDynamo() {
    DynamoTestUtils.teardownDynamo();
  }

  @BeforeEach
  void setUp() {

    testResource = Resource.builder()
        .id("a66de382-a9df-4fab-9d34-616e01e3e054")
        .title("title")
        .tags(Arrays.asList("art", "music"))
        .createdAt(new Date().toString())
        .updatedAt(new Date().toString())
        .build();
  }

  @Test
  @DisplayName("Create a resource successfully")
  void createResourceSuccess() throws Exception {
    ResourceModel inputResource = ResourceModel.builder()
        .title("title")
        .tags(Arrays.asList("art", "music"))
        .build();

    Mockito.when(resourceService.createResource(
        ArgumentMatchers.any())).thenReturn(testResource);

    mockMvc.perform(
        post("/resource")
            .contentType(MediaType.APPLICATION_JSON)
            .content(Jackson.toJsonString(inputResource)))
        .andExpect(status().isCreated())
        .andExpect(jsonPath("$.data").value(testResource));
  }

  @Test
  @DisplayName("Attempt to create a resource with an invalid input model")
  void createResourceInvalidInput() throws Exception {
    ResourceModel inputResource = ResourceModel.builder()
        .title("title")
        .tags(Arrays.asList("art", "music"))
        .build();

    Mockito.when(resourceService.createResource(inputResource)).thenReturn(testResource);

    mockMvc.perform(
        post("/resource")
            .contentType(MediaType.APPLICATION_JSON)
            .content("{\"titl ez\": \"title\", \"tagds\": [\"art\",\"music\"] }"))
        .andExpect(status().isBadRequest())
        .andExpect(jsonPath("$.internal").isString())
        .andExpect(jsonPath("$.friendly").value("The provided input was not valid."))
        .andExpect(jsonPath("$.data").isEmpty());
  }

  @Test
  @DisplayName("Retrieve a resource by ID")
  void getResourceById() throws Exception {
    String testId = "a66de382-a9df-4fab-9d34-616e01e3e054";

    Mockito.when(resourceService.getResourceById(testId)).thenReturn(testResource);

    mockMvc.perform(
        get(String.format("/resource/%s", testId)))
        .andExpect(status().isOk())
        .andExpect(jsonPath("$.data").value(testResource));
  }

  @Test
  @DisplayName("Attempt to retrieve a resource that does not exist")
  void getResourceByIdNotFound() throws Exception {
    String testId = "a66de382-a9df-4fab-9d34-616e01e3e054";

    Mockito.when(resourceService.getResourceById(testId)).thenReturn(testResource);

    mockMvc.perform(
        get(String.format("/resource/%s", testId + "1")))
        .andExpect(status().isNoContent())
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
            .content(Jackson.toJsonString(inputModel)))
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
            .content(Jackson.toJsonString(inputModel)))
        .andExpect(status().isBadRequest())
        .andExpect(jsonPath("$.internal").isString())
        .andExpect(jsonPath("$.friendly").value("The provided input was not valid."))
        .andExpect(jsonPath("$.data").isEmpty());
  }

  @Test
  @DisplayName("Delete a resource by ID")
  void deleteResourceById() throws Exception {
    String testId = "a66de382-a9df-4fab-9d34-616e01e3e054";

    mockMvc.perform(
        delete(String.format("/resource/%s", testId)))
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
        delete(String.format("/resource/%s", testId)))
        .andExpect(status().isBadRequest())
        .andExpect(jsonPath("$.data").isEmpty())
        .andExpect(jsonPath("$.friendly").value(
            String.format("An exception occurred while deleting the resource with id %s", testId)))
        .andExpect(jsonPath("$.internal").value("test"));
  }
}
