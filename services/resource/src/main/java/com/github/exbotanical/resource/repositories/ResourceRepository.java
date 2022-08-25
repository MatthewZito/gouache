package com.github.exbotanical.resource.repositories;

import com.amazonaws.services.dynamodbv2.datamodeling.DynamoDBMapper;
import com.amazonaws.services.dynamodbv2.datamodeling.DynamoDBSaveExpression;
import com.amazonaws.services.dynamodbv2.model.AttributeValue;
import com.amazonaws.services.dynamodbv2.model.ExpectedAttributeValue;
import com.github.exbotanical.resource.entities.Resource;
import com.github.exbotanical.resource.models.ResourceModel;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

@Repository
public class ResourceRepository {
  @Autowired
  private DynamoDBMapper dynamoDBMapper;

  public Resource save(Resource newResource) {
    dynamoDBMapper.save(newResource);

    return newResource;
  }

  public Resource getById(String id) {
    return dynamoDBMapper.load(Resource.class, id);
  }

  public void deleteById(String id) {
    Resource resource = dynamoDBMapper.load(Resource.class, id);

    dynamoDBMapper.delete(resource);
  }

  public void updateById(String id, ResourceModel resourceModel) {
    Resource updatedResource = Resource.builder()
      .id(id)
      .title(resourceModel.getTitle())
      .tags(resourceModel.getTags())
      .build();

    dynamoDBMapper.save(
      updatedResource,
      new DynamoDBSaveExpression().withExpectedEntry(
        "Id",
        new ExpectedAttributeValue(
          new AttributeValue().withS(id))));
  }
}
