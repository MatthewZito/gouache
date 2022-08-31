package com.github.exbotanical.resource.services;

import com.github.exbotanical.resource.entities.Resource;
import com.github.exbotanical.resource.models.ResourceModel;
import java.util.ArrayList;

/**
 * Service for Resource data.
 */
public interface ResourceService {

  /**
   * Retrieve a Resource by its id.
   *
   * @param id A unique Resource id identifying the Resource to retrieve.
   * @return A Resource, or null if not found.
   */
  Resource getResourceById(String id);

  /**
   * Retrieve all Resources.
   *
   * @return A list of Resources.
   */
  ArrayList<Resource> getAllResources();

  /**
   * Create a new Resource.
   *
   * @param resourceModel A ResourceModel containing the requisite data to create a Resource.
   * @return The newly-created Resource.
   */
  Resource createResource(ResourceModel resourceModel);

  /**
   * Delete a Resource by its id.
   *
   * @param id A unique Resource id identifying the Resource to delete.
   */
  void deleteResourceById(String id);

  /**
   * Update a Resource by its id.
   *
   * @param id A unique Resource id identifying the Resource to update.
   * @param resourceModel A ResourceModel containing the data to patch into the id-resolved
   *        Resource.
   */
  void updateResourceById(String id, ResourceModel resourceModel);
}
