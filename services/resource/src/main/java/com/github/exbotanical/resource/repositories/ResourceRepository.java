package com.github.exbotanical.resource.repositories;

import com.amazonaws.services.dynamodbv2.datamodeling.DynamoDBMapper;
import com.amazonaws.services.dynamodbv2.datamodeling.DynamoDBSaveExpression;
import com.amazonaws.services.dynamodbv2.model.AttributeValue;
import com.amazonaws.services.dynamodbv2.model.ExpectedAttributeValue;
import com.github.exbotanical.resource.entities.Resource;
import com.github.exbotanical.resource.models.ResourceModel;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

/**
 * A repository for managing Resource data via DynamoDB.
 */
@Repository
public class ResourceRepository {

  @Autowired
  private DynamoDBMapper dynamoDBMapper;

  /**
   * Persist a given Resource.
   *
   * @param newResource The new Resource to persist.
   * @return The new Resource, with updated Dynamo-generated fields.
   */
  public Resource save(Resource newResource) {
    dynamoDBMapper.save(newResource);

    return newResource;
  }

  /**
   * Retrieve a Resource by its id.
   *
   * @param id A unique Resource id identifying the Resource to retrieve.
   * @return A Resource, or null if not found.
   */
  public Resource getById(String id) {
    return dynamoDBMapper.load(Resource.class, id);
  }

  /**
   * Delete a Resource by its id.
   *
   * @param id A unique Resource id identifying the Resource to delete.
   */
  public void deleteById(String id) {
    Resource resource = dynamoDBMapper.load(Resource.class, id);

    dynamoDBMapper.delete(resource);
  }

  /**
   * Update a Resource by its id.
   *
   * @param id A unique Resource id identifying the Resource to update.
   * @param resourceModel A ResourceModel containing the data to patch into the id-resolved
   *        Resource.
   */
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
