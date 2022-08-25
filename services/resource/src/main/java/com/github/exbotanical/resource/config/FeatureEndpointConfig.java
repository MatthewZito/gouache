package com.github.exbotanical.resource.config;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.boot.actuate.endpoint.annotation.Endpoint;
import org.springframework.boot.actuate.endpoint.annotation.ReadOperation;
import org.springframework.boot.actuate.endpoint.annotation.Selector;
import org.springframework.stereotype.Component;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

/**
 * Internal metadata and health / status endpoint config for Spring Actuator.
 */
@Component
@Endpoint(id = "features")
public class FeatureEndpointConfig {
  private final Map<String, Feature> featuresMap = new ConcurrentHashMap<>();

  public FeatureEndpointConfig() {
    featuresMap.put("resource", new Feature(true));
    featuresMap.put("cache", new Feature(false));
    featuresMap.put("roles", new Feature(false));
  }

  @ReadOperation
  public Map<String, Feature> features() {
    return featuresMap;
  }

  @ReadOperation
  public Feature feature(@Selector String featureKey) {
    return featuresMap.get(featureKey);
  }

  @Data
  @NoArgsConstructor
  @AllArgsConstructor
  private static class Feature {
    private boolean isEnabled;
  }
}
