package com.github.exbotanical.resource.controllers;

import com.github.exbotanical.resource.entities.Resource;
import com.github.exbotanical.resource.errors.GouacheException;
import com.github.exbotanical.resource.errors.InvalidInputException;
import com.github.exbotanical.resource.errors.OperationFailedException;
import com.github.exbotanical.resource.models.ResourceModel;
import com.github.exbotanical.resource.services.ResourceService;
import com.github.exbotanical.resource.utils.FormatterUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;

@RestController
public class ResourceController {
  @Value("${app.locale}")
  private String locale;

  @Autowired
  private ResourceService resourceService;

  private final Logger LOGGER = LoggerFactory.getLogger(ResourceController.class);

  @PostMapping("/resource")
  public ResponseEntity<Resource> createResource(@Valid @RequestBody ResourceModel resourceModel,
                                                 BindingResult result) throws GouacheException {
    if (result.hasErrors()) {
      throw new InvalidInputException(FormatterUtils.formatValidationErrors(result));
    }

    Resource newResource = resourceService.createResource(resourceModel);
    return new ResponseEntity<>(newResource, HttpStatus.CREATED);
  }

  @GetMapping("/resource/{id}")
  public ResponseEntity<Resource> getResourceById(@PathVariable("id") String id) {

    Resource found = resourceService.getResourceById(id);
    if (found == null) {
      return new ResponseEntity<>(null, HttpStatus.NO_CONTENT);
    }

    return new ResponseEntity<>(found, HttpStatus.OK);
  }

  @PatchMapping("/resource/{id}")
  public ResponseEntity<Void> updateResourceById(@PathVariable("id") String id,
                                                 @Valid @RequestBody ResourceModel resourceModel, BindingResult result) throws GouacheException {
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
        e
      );
    }
  }

  @DeleteMapping("/resource/{id}")
  public ResponseEntity<Void> deleteResourceById(@PathVariable("id") String id) throws GouacheException {
    try {
      resourceService.deleteResourceById(id);
      return new ResponseEntity<>(null, HttpStatus.OK);
    } catch (Exception e) {
      throw new OperationFailedException(
        String.format("An exception occurred while deleting the resource with id %s", id),
        e.getMessage(),
        e
      );
    }
  }
}
