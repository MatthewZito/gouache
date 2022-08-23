package com.github.exbotanical.session.controllers;

import com.github.exbotanical.session.entities.Resource;
import com.github.exbotanical.session.services.ResourceService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.web.servlet.MockMvc;

@WebMvcTest(ResourceController.class)
class ResourceControllerTest {

  @Autowired
  private MockMvc mockMvc;

  @MockBean
  private ResourceService resourceService;

  private Resource testResource;

//  @BeforeEach
//  void setUp() {
//    testResource = Resource.builder()
//            .id(1L)
//            .name("t")
//            .code("abc")
//            .data("test")
//            .build();
//  }
//
//  @Test
//  void createResource() throws Exception {
//    Resource inputResource = Resource.builder()
//            .id(1L)
//            .name("t")
//            .code("abc")
//            .data("test")
//            .build();
//
//    Mockito.when(resourceService.createResource(inputResource)).thenReturn(testResource);
//
//    mockMvc.perform(
//                    post("/resource")
//                            .contentType(MediaType.APPLICATION_JSON)
//                            .content("todo"))
//            .andExpect(status().isOk());
//  }
//
//  @Test
//  void getResource() throws Exception {
//    Mockito.when(resourceService.getResource(testResource.getId()))
//            .thenReturn(testResource);
//
//    mockMvc.perform(
//                    get("/resource/1")
//                            .contentType(MediaType.APPLICATION_JSON)
//            )
//            .andExpect(status().isOk())
//            .andExpect(jsonPath("$.name").value(testResource.getName()));
//  }
//
//  @Test
//  void deleteResource() {
//  }
//
//  @Test
//  void getResources() {
//  }
//
//  @Test
//  void updateResource() {
//  }
}