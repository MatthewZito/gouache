package com.github.exbotanical.session.services;

import com.github.exbotanical.session.repositories.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;

@SpringBootTest
class UserServiceTest {
  @Autowired
  private UserService userService;

  @MockBean
  private UserRepository userRepository;

  // User testUser = User.builder()
  // .id(1L)
  // .name("t")
  // .code("abc")
  // .data("test")
  // .build();

  // @BeforeEach
  // void setUp() {
  //
  // Mockito.when(userRepository.findById(1L)).thenReturn(Optional.ofNullable(testUser));
  // }
  //
  // @AfterEach
  // void tearDown() {
  // }
  //
  // @Test
  // @DisplayName("get User by id")
  // void shouldGetUserById() {
  // try {
  // User found = userService.getUser(testUser.getId());
  // assertEquals(testUser, found);
  // } catch (UserNotFoundException e) {
  // fail(e);
  // }
  // }
  //
  // @Test
  // void createUser() {
  // }
  //
  // @Test
  // void getUser() {
  // }
  //
  // @Test
  // void getUsers() {
  // }
  //
  // @Test
  // void deleteUser() {
  // }
  //
  // @Test
  // void updateUser() {
  // }
}
