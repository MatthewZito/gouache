package com.github.exbotanical.session.repositories;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;
import org.springframework.boot.test.autoconfigure.orm.jpa.TestEntityManager;

@DataJpaTest
class UserRepositoryTest {
  @Autowired
  private UserRepository userRepository;

  @Autowired
  private TestEntityManager entityManager;
  //
  // Resource testResource = Resource.builder()
  // .id(1L)
  // .name("t")
  // .code("abc")
  // .data("test")
  // .build();
  //
  // @BeforeEach
  // void setUp() {
  // entityManager.persist(testResource);
  // }
  //
  // @Test
  // @DisplayName("finds a resource by its id")
  // public void shouldFindById() {
  // Resource found = resourceRepository.findById(testResource.getId()).get();
  //
  // assertEquals(testResource, found);
  // }
}
