package com.github.exbotanical.session.services;

import com.github.exbotanical.session.repositories.ResourceRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;

@SpringBootTest
class ResourceServiceTest {
  @Autowired
  private ResourceService resourceService;

  @MockBean
  private ResourceRepository resourceRepository;

//  Resource testResource = Resource.builder()
//          .id(1L)
//          .name("t")
//          .code("abc")
//          .data("test")
//          .build();

//  @BeforeEach
//  void setUp() {
//
//    Mockito.when(resourceRepository.findById(1L)).thenReturn(Optional.ofNullable(testResource));
//  }
//
//  @AfterEach
//  void tearDown() {
//  }
//
//  @Test
//  @DisplayName("get resource by id")
//  void shouldGetResourceById() {
//    try {
//      Resource found = resourceService.getResource(testResource.getId());
//      assertEquals(testResource, found);
//    } catch (ResourceNotFoundException e) {
//      fail(e);
//    }
//  }
//
//  @Test
//  void createResource() {
//  }
//
//  @Test
//  void getResource() {
//  }
//
//  @Test
//  void getResources() {
//  }
//
//  @Test
//  void deleteResource() {
//  }
//
//  @Test
//  void updateResource() {
//  }
}