package com.github.exbotanical.session.services;

import com.github.exbotanical.session.entities.Resource;
import com.github.exbotanical.session.errors.ResourceNotFoundException;
import com.github.exbotanical.session.repositories.ResourceRepository;
import java.util.List;
import java.util.Optional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class ResourceDataService implements ResourceService {
  @Autowired
  private ResourceRepository resourceRepository;

  @Override
  public Resource createResource(Resource resource) {
    return resourceRepository.save(resource);
  }

  @Override
  public Resource getResourceById(Long id) throws ResourceNotFoundException {
    Optional<Resource> maybeResource = resourceRepository.findById(id);

    if (maybeResource.isEmpty()) {
      throw new ResourceNotFoundException(id);
    }

    return maybeResource.get();
  }

  @Override
  public List<Resource> getResources() {
    return resourceRepository.findAll();
  }

  @Override
  public void deleteResourceById(Long id) throws ResourceNotFoundException {
    if (!resourceRepository.existsById(id)) {
      throw new ResourceNotFoundException(id);
    }

    resourceRepository.deleteById(id);
  }

  @Override
  public Long updateResourceById(Resource resource) {
    resourceRepository.save(resource);

    return resource.getId();
  }
}