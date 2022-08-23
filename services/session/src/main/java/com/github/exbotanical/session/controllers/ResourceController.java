package com.github.exbotanical.session.controllers;

import com.github.exbotanical.session.entities.Resource;
import com.github.exbotanical.session.models.ResourceModel;
import com.github.exbotanical.session.services.ResourceService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
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
}
