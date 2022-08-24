package com.github.exbotanical.resource.controllers;

import com.github.exbotanical.resource.entities.Resource;
import com.github.exbotanical.resource.errors.GouacheException;
import com.github.exbotanical.resource.errors.InvalidInputException;
import com.github.exbotanical.resource.models.ResourceModel;
import com.github.exbotanical.resource.services.ResourceService;
import com.github.exbotanical.resource.utils.FormatterUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
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

  @GetMapping("/resource/{id}")
  public Resource getResourceById(@PathVariable("id") String id) {
    return resourceService.getResourceById(id);
  }

  @DeleteMapping("/resource/{id}")
  public void deleteResourceById(@PathVariable("id") String id) {
    resourceService.deleteResourceById(id);
  }

  @PatchMapping("/resource/{id}")
  public void updateResourceById(@PathVariable("id") String id,
                                 @Valid @RequestBody ResourceModel resourceModel) {
    resourceService.updateResourceById(id, resourceModel);
  }

  @PostMapping("/resource")
  public Resource createResource(@Valid @RequestBody ResourceModel resourceModel,
                                 BindingResult result) throws GouacheException {
    if (result.hasErrors()) {
      throw new InvalidInputException(FormatterUtils.formatValidationErrors(result));
    }

    return resourceService.createResource(resourceModel);
  }
}
