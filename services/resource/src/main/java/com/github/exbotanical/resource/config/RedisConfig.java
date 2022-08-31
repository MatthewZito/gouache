package com.github.exbotanical.resource.config;

import com.github.exbotanical.resource.entities.Session;
import java.util.stream.Stream;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.env.Environment;
import org.springframework.data.redis.cache.RedisCacheConfiguration;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.connection.RedisStandaloneConfiguration;
import org.springframework.data.redis.connection.lettuce.LettuceConnectionFactory;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.serializer.StringRedisSerializer;

/**
 * Redis configuration.
 */
@Configuration
public class RedisConfig {

  @Autowired
  private Environment env;

  @Value("${app.redis.host}")
  private String redisHost;

  @Value("${app.redis.port}")
  private String redisPort;

  @Value("${app.redis.password}")
  private String redisPassword;

  @Bean
  public RedisCacheConfiguration redisCacheConfiguration() {
    return RedisCacheConfiguration.defaultCacheConfig();
  }

  /**
   * Create a Redis connection object.
   *
   * @return An initialized Redis connection object.
   */
  @Bean
  public RedisConnectionFactory redisConnectionFactory() {
    RedisStandaloneConfiguration connectionConfig =
        new RedisStandaloneConfiguration(redisHost, Integer.parseInt(redisPort));

    // No password necessary in local mode.
    if (Stream.of(env.getActiveProfiles()).noneMatch(v -> v.equals("local"))) {
      connectionConfig.setPassword(redisPassword);
    }

    return new LettuceConnectionFactory(connectionConfig);
  }

  /**
   * Create and return a Redis configuration.
   */
  @Bean
  public RedisTemplate<String, Session> redisTemplate() {
    RedisTemplate<String, Session> template = new RedisTemplate<>();
    template.setConnectionFactory(redisConnectionFactory());

    // Prevent keys from being serialized into byte strings.
    template.setDefaultSerializer(new StringRedisSerializer());
    template.setKeySerializer(new StringRedisSerializer());
    template.setValueSerializer(new StringRedisSerializer());

    return template;
  }
}
