package com.github.exbotanical.resource.services;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.junit.jupiter.api.Assertions.assertThrows;

import com.github.exbotanical.resource.SessionTestUtils;
import com.github.exbotanical.resource.entities.Resource;
import com.github.exbotanical.resource.models.ResourceModel;
import com.github.exbotanical.resource.repositories.ResourceRepository;
import java.util.Arrays;
import java.util.Date;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentMatchers;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;

/**
 * Test ResourceService.
 */
@SpringBootTest
@DisplayName("Test ResourceService")
public class ResourceServiceTest {

  @Autowired
  private ResourceService resourceService;

  @MockBean
  private SessionService sessionService;

  @MockBean
  private ResourceRepository resourceRepository;

  private Resource testResource;

  @BeforeEach
  void setUp() {
    Mockito.when(sessionService.getSessionBySessionId(ArgumentMatchers.anyString()))
        .thenReturn(SessionTestUtils.session);

    testResource = Resource.builder().id("a66de382-a9df-4fab-9d34-616e01e3e054").title("title")
        .tags(Arrays.asList("art", "music")).createdAt(new Date().toString())
        .updatedAt(new Date().toString()).build();
  }

  @Test
  @DisplayName("Create a resource")
  void createResourceSuccess() {
    ResourceModel inputResource =
        ResourceModel.builder().title("title").tags(Arrays.asList("art", "music")).build();

    Resource newResource =
        Resource.builder().title(inputResource.getTitle()).tags(inputResource.getTags()).build();

    Mockito.when(resourceRepository.save(newResource)).thenReturn(testResource);

    assertEquals(testResource, resourceService.createResource(inputResource));
  }

  @Test
  @DisplayName("Retrieve a resource by ID")
  void getResourceById() {
    String testId = "a66de382-a9df-4fab-9d34-616e01e3e054";

    Mockito.when(resourceRepository.getById(testId)).thenReturn(testResource);

    assertEquals(testResource, resourceService.getResourceById(testId));
  }

  @Test
  @DisplayName("Attempt to retrieve a resource that does not exist")
  void getResourceByIdNotFound() {
    String testId = "a66de382-a9df-4fab-9d34-616e01e3e054";

    Mockito.when(resourceRepository.getById(testId)).thenReturn(testResource);

    assertNull(resourceService.getResourceById(testId + 1));
  }

  @Test
  @DisplayName("Update a resource by ID")
  void updateResourceById() {
    ResourceModel inputModel =
        ResourceModel.builder().tags(Arrays.asList("test")).title("test title").build();

    assertDoesNotThrow(() -> resourceService.updateResourceById(testResource.getId(), inputModel));
  }

  @Test
  @DisplayName("Delete a resource by ID")
  void deleteResourceById() {
    assertDoesNotThrow(() -> resourceService.deleteResourceById(testResource.getId()));
  }

  @Test
  @DisplayName("Erroneous delete resource by ID")
  void deleteResourceByIdError() {
    String testId = testResource.getId();

    Mockito.doThrow(new RuntimeException("test")).when(resourceRepository).deleteById(testId);

    assertThrows(RuntimeException.class, () -> resourceService.deleteResourceById(testId));
  }
}
