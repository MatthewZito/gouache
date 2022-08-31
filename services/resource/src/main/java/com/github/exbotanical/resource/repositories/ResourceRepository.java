package com.github.exbotanical.resource.repositories;

import com.github.exbotanical.resource.entities.Resource;
import com.github.exbotanical.resource.models.ResourceModel;
import java.time.Instant;
import java.util.ArrayList;
import java.util.UUID;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbEnhancedClient;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbTable;
import software.amazon.awssdk.enhanced.dynamodb.Key;
import software.amazon.awssdk.enhanced.dynamodb.TableSchema;

/**
 * A repository for managing Resource data via DynamoDB.
 */
@Repository
public class ResourceRepository {

  private final DynamoDbTable<Resource> resourceTable;

  public ResourceRepository(@Autowired DynamoDbEnhancedClient dynamo) {
    this.resourceTable = dynamo.table(Resource.TABLE_NAME, TableSchema.fromBean(Resource.class));
  }

  /**
   * Persist a given Resource.
   *
   * @param newResource The new Resource to persist.
   *
   * @return The new Resource, with updated Dynamo-generated fields.
   */
  public Resource save(Resource newResource) {
    newResource.setId(UUID.randomUUID().toString());
    setTimeStampMutable(newResource, true);

    System.out.println(newResource.getId());
    System.out.println(newResource.getTitle());
    System.out.println(newResource.getTags());
    System.out.println(newResource.getCreatedAt());
    System.out.println(newResource.getUpdatedAt());

    resourceTable.putItem(newResource);

    return newResource;
  }

  /**
   * Retrieve a Resource by its id.
   *
   * @param id A unique Resource id identifying the Resource to retrieve.
   *
   * @return A Resource, or null if not found.
   */
  public Resource getById(String id) {
    return resourceTable.getItem(Key.builder().partitionValue(id).build());
  }

  /**
   * Retrieve all Resources.
   *
   * @return List of Resources.
   *
   * @todo Paginate
   */
  public ArrayList<Resource> getAll() {
    // @todo improve
    ArrayList<Resource> list = new ArrayList<>();

    resourceTable.scan().stream()
        .forEach(page -> page.items().stream()
            .filter(item -> item != null)
            .forEach(item -> list.add(item)));


    return list;
  }

  /**
   * Delete a Resource by its id.
   *
   * @param id A unique Resource id identifying the Resource to delete.
   */
  public void deleteById(String id) {
    resourceTable.deleteItem(Key.builder().partitionValue(id).build());
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

    setTimeStampMutable(updatedResource, false);

    resourceTable.updateItem(updatedResource);
  }

  /**
   * Set timestamps on the Resource, mutating the original object.
   *
   * @param resource The Resource on which to write the timestamp(s).
   * @param isCreate Whether the `createdAt` timestamp value should be set.
   *
   * @apiNote Dynamodb currently does not execute the autogenerate timestamps annotations, and their
   *          documentation thereof is very sparse.
   * @todo Leverage auto generation.
   */
  private void setTimeStampMutable(Resource resource, boolean isCreate) {
    Instant now = Instant.now();

    if (isCreate) {
      resource.setCreatedAt(now);
    }

    resource.setUpdatedAt(now);
  }
}
