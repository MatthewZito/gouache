package com.github.exbotanical.resource;

import static org.junit.jupiter.api.Assertions.assertEquals;

import com.amazonaws.services.dynamodbv2.local.main.ServerRunner;
import com.amazonaws.services.dynamodbv2.local.server.DynamoDBProxyServer;
import java.io.IOException;
import java.net.ServerSocket;
import java.net.URI;
import org.springframework.beans.factory.annotation.Value;
import software.amazon.awssdk.auth.credentials.AwsBasicCredentials;
import software.amazon.awssdk.auth.credentials.StaticCredentialsProvider;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.dynamodb.DynamoDbAsyncClient;
import software.amazon.awssdk.services.dynamodb.model.AttributeDefinition;
import software.amazon.awssdk.services.dynamodb.model.CreateTableRequest;
import software.amazon.awssdk.services.dynamodb.model.KeySchemaElement;
import software.amazon.awssdk.services.dynamodb.model.KeyType;
import software.amazon.awssdk.services.dynamodb.model.ListTablesResponse;
import software.amazon.awssdk.services.dynamodb.model.ProvisionedThroughput;
import software.amazon.awssdk.services.dynamodb.model.ScalarAttributeType;

/**
 * Test utilities for DynamoDB local setup.
 */
public class DynamoTestUtils {

  @Value("${aws.dynamodb.host}")
  static String dynamoHost;

  @Value("${aws.dynamodb.region}")
  static String dynamoRegion;

  @Value("${aws.dynamodb.accessKey}")
  static String dynamoAccessKey;

  @Value("${aws.dynamodb.secretKey}")
  static String dynamoSecretKey;

  private static DynamoDBProxyServer dynamoProxy;

  private static int port;


  /**
   * Find an available port.
   *
   * @return Available port number.
   */
  private static int getAvailablePort() {
    try {
      ServerSocket socket = new ServerSocket(0);
      int port = socket.getLocalPort();

      socket.close();

      return port;
    } catch (IOException e) {
      throw new RuntimeException(e);
    }
  }

  public int getPort() {
    return port;
  }

  /**
   * Setup a DynamoDB local connection.
   */
  public static void setupDynamo() {
    port = getAvailablePort();

    try {
      dynamoProxy = ServerRunner.createServerFromCommandLineArgs(new String[] {
          "-inMemory",
          "-port",
          Integer.toString(port)
      });

      dynamoProxy.start();

      // setupDynamoTables();
    } catch (Exception e) {
      throw new RuntimeException();
    }
  }

  /**
   * Stop DynamoDB local proxy.
   */
  public static void teardownDynamo() {
    try {
      dynamoProxy.stop();
    } catch (Exception e) {
      throw new RuntimeException();
    }
  }

  /**
   * Setup tables for DynamoDB testing.
   *
   * @throws Exception Table creation errors.
   */
  public static void setupDynamoTables() throws Exception {
    try (DynamoDbAsyncClient client = DynamoDbAsyncClient.builder()
        .region(Region.US_EAST_1)
        .endpointOverride(URI.create(dynamoHost + port))
        .credentialsProvider(StaticCredentialsProvider.create(
            AwsBasicCredentials.create(
                dynamoAccessKey, dynamoSecretKey)))
        .build()) {
      ListTablesResponse listTablesResponse = client.listTables().get();

      assertEquals(listTablesResponse.tableNames().size(), 0);

      client.createTable(CreateTableRequest.builder()
          .keySchema(
              KeySchemaElement.builder().keyType(KeyType.HASH).attributeName("Id").build())
          .attributeDefinitions(
              AttributeDefinition.builder().attributeName("Id")
                  .attributeType(ScalarAttributeType.S).build())

          .provisionedThroughput(ProvisionedThroughput.builder().readCapacityUnits(100L)
              .writeCapacityUnits(100L).build())

          .tableName("resource")
          .build())
          .get();

      ListTablesResponse listTablesResponseAfterCreation = client.listTables().get();

      assertEquals(listTablesResponseAfterCreation.tableNames().size(), 1);
    }
  }
}
