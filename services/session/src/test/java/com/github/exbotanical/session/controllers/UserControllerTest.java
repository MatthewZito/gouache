package com.github.exbotanical.session.controllers;

import com.github.exbotanical.session.controllers.UserController;
import com.github.exbotanical.session.entities.User;
import com.github.exbotanical.session.services.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.web.servlet.MockMvc;

@WebMvcTest(UserController.class)
class UserControllerTest {

  @Autowired
  private MockMvc mockMvc;

  @MockBean
  private UserService userService;

  private User testUser;

  // @BeforeEach
  // void setUp() {
  // testUser = User.builder()
  // .id(1L)
  // .name("t")
  // .code("abc")
  // .data("test")
  // .build();
  // }
  //
  // @Test
  // void createUser() throws Exception {
  // User inputUser = User.builder()
  // .id(1L)
  // .name("t")
  // .code("abc")
  // .data("test")
  // .build();
  //
  // Mockito.when(userService.createUser(inputUser)).thenReturn(testUser);
  //
  // mockMvc.perform(
  // post("/User")
  // .contentType(MediaType.APPLICATION_JSON)
  // .content("todo"))
  // .andExpect(status().isOk());
  // }
  //
  // @Test
  // void getUser() throws Exception {
  // Mockito.when(userService.getUser(testUser.getId()))
  // .thenReturn(testUser);
  //
  // mockMvc.perform(
  // get("/User/1")
  // .contentType(MediaType.APPLICATION_JSON)
  // )
  // .andExpect(status().isOk())
  // .andExpect(jsonPath("$.name").value(testUser.getName()));
  // }
  //
  // @Test
  // void deleteUser() {
  // }
  //
  // @Test
  // void getUsers() {
  // }
  //
  // @Test
  // void updateUser() {
  // }
}
