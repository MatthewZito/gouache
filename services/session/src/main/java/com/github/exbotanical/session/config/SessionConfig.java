package com.github.exbotanical.session.config;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.data.redis.connection.RedisStandaloneConfiguration;
import org.springframework.data.redis.connection.lettuce.LettuceConnectionFactory;
import org.springframework.session.data.redis.config.annotation.web.http.EnableRedisHttpSession;


@EnableRedisHttpSession
public class SessionConfig {

  @Value("${spring.redis.host}")
  String redisHost;

  @Value("${spring.redis.port}")
  String redisPort;

  @Value("${spring.redis.password}")
  String redisPassword;

  @Bean
  public LettuceConnectionFactory connectionFactory() {
    RedisStandaloneConfiguration config =
        new RedisStandaloneConfiguration(redisHost, Integer.valueOf(redisPort));
    // config.setPassword(redisPassword);

    return new LettuceConnectionFactory(config);
  }
}
