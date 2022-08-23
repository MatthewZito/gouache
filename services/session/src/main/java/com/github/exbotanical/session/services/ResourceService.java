package com.github.exbotanical.session.services;

import com.github.exbotanical.session.entities.Resource;
import com.github.exbotanical.session.models.ResourceModel;
import java.util.List;

public interface ResourceService {
  List<Resource> getResources();

  Resource getResourceById(String id);

  Resource createResource(ResourceModel resourceModel);

  void deleteResourceById(String id);

  void updateResourceById(String id, ResourceModel resourceModel);
}
