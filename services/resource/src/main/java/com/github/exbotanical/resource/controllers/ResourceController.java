package com.github.exbotanical.resource.controllers;

import com.github.exbotanical.resource.entities.Resource;
import com.github.exbotanical.resource.errors.GouacheException;
import com.github.exbotanical.resource.errors.InvalidInputException;
import com.github.exbotanical.resource.errors.OperationFailedException;
import com.github.exbotanical.resource.models.ResourceModel;
import com.github.exbotanical.resource.services.ResourceService;
import com.github.exbotanical.resource.utils.FormatterUtils;

import java.util.ArrayList;

import javax.validation.Valid;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PatchMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;


/**
 * A REST controller for Resource CRUD operations.
 *
 * @todo Paginate get all.
 * @todo Restrict resources to current user.
 */
@RestController
public class ResourceController {
  @Value("${app.locale}")
  private String locale;

  @Autowired
  private ResourceService resourceService;

  private final Logger logger = LoggerFactory.getLogger(ResourceController.class);

  /**
   * Create Resource endpoint.
   *
   * @param resourceModel Client-provided data with which the system will create the Resource.
   * @param result An input validation result.
   * @return A ResponseEntity containing the newly-created Resource.
   * @throws GouacheException Resource creation errors.
   */
  @PostMapping("/resource")
  public ResponseEntity<Resource> createResource(@Valid @RequestBody ResourceModel resourceModel,
      BindingResult result) throws GouacheException {
    if (result.hasErrors()) {
      throw new InvalidInputException(FormatterUtils.formatValidationErrors(result));
    }

    Resource newResource = resourceService.createResource(resourceModel);
    return new ResponseEntity<>(newResource, HttpStatus.CREATED);
  }

  /**
   * Get Resource by id endpoint.
   *
   * @param id Client provided id.
   * @return A ResponseEntity containing the found Resource, or null if not extant.
   */
  @GetMapping("/resource/{id}")
  public ResponseEntity<Resource> getResourceById(@PathVariable("id") String id) {

    Resource found = resourceService.getResourceById(id);
    if (found == null) {
      return new ResponseEntity<>(null, HttpStatus.NOT_FOUND);
    }

    return new ResponseEntity<>(found, HttpStatus.OK);
  }

  /**
   * Get all Resources.
   *
   * @return A list of Resources.
   */
  @GetMapping("/resource")
  public ResponseEntity<ArrayList<Resource>> getAllResources() {
    ArrayList<Resource> allResources = resourceService.getAllResources();

    return new ResponseEntity<>(allResources, HttpStatus.OK);
  }

  /**
   * Update a Resource by id endpoint.
   *
   * @param id Client provided id.
   * @param resourceModel Client provided ResourceModel for patching.
   * @param result An input validation result.
   * @return An empty ResponseEntity.
   * @throws GouacheException Operation error.
   */
  @PatchMapping("/resource/{id}")
  public ResponseEntity<Void> updateResourceById(@PathVariable("id") String id,
      @Valid @RequestBody ResourceModel resourceModel, BindingResult result)
      throws GouacheException {
    if (result.hasErrors()) {
      throw new InvalidInputException(FormatterUtils.formatValidationErrors(result));
    }

    try {
      System.out.println("AYY: " + resourceModel);
      resourceService.updateResourceById(id, resourceModel);

      return new ResponseEntity<>(null, HttpStatus.OK);
    } catch (Exception e) {
      throw new OperationFailedException(
          String.format("An exception occurred while updating the resource with id %s", id),
          e.getMessage(),
          e);
    }
  }

  /**
   * Delete a Resource by id endpoint.
   *
   * @param id Client provided id.
   * @return An empty ResponseEntity.
   * @throws GouacheException Operation error.
   */
  @DeleteMapping("/resource/{id}")
  public ResponseEntity<Void> deleteResourceById(@PathVariable("id") String id)
      throws GouacheException {
    try {
      resourceService.deleteResourceById(id);
      return new ResponseEntity<>(null, HttpStatus.OK);
    } catch (Exception e) {
      throw new OperationFailedException(
          String.format("An exception occurred while deleting the resource with id %s", id),
          e.getMessage(),
          e);
    }
  }
}
