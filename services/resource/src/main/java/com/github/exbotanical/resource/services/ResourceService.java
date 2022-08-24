package com.github.exbotanical.resource.services;

import com.github.exbotanical.resource.entities.Resource;
import com.github.exbotanical.resource.models.ResourceModel;

public interface ResourceService {
  Resource getResourceById(String id);

  Resource createResource(ResourceModel resourceModel);

  void deleteResourceById(String id);

  void updateResourceById(String id, ResourceModel resourceModel);
}
