package com.github.exbotanical.resource.controllers;

import com.amazonaws.util.json.Jackson;
import com.github.exbotanical.resource.DynamoTestUtils;
import com.github.exbotanical.resource.entities.Resource;
import com.github.exbotanical.resource.models.ResourceModel;
import com.github.exbotanical.resource.services.ResourceService;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentMatchers;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;

import java.util.Arrays;
import java.util.Date;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(ResourceController.class)
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
      .andExpect(status().isOk())
      .andExpect(jsonPath("$.data").value(testResource));
  }

  @Test
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
      .andExpect(jsonPath("$.internal").value("resourceModel.tags must not be null, resourceModel.title must not be null"))
      .andExpect(jsonPath("$.friendly").value("The provided input was not valid."))
      .andExpect(jsonPath("$.data").isEmpty());
  }
  
  // @Test
  // void createResourceUnauthorized() throws Exception {
  // ResourceModel inputResource = ResourceModel.builder()
  // .title("title")
  // .tags(Arrays.asList("art", "music"))
  // .build();

  // Mockito.when(resourceService.createResource(inputResource)).thenReturn(testResource);

  // mockMvc.perform(
  // post("/resource")
  // .contentType(MediaType.APPLICATION_JSON)
  // .content("{\"title\": \"title\", \"tags\": [\"art\",\"music\"] }"))
  // .andExpect(status().isOk());
  // }
}
