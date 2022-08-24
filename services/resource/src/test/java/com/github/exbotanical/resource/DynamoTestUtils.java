package com.github.exbotanical.resource;

import com.amazonaws.services.dynamodbv2.local.main.ServerRunner;
import com.amazonaws.services.dynamodbv2.local.server.DynamoDBProxyServer;
import org.springframework.beans.factory.annotation.Value;
import software.amazon.awssdk.auth.credentials.AwsBasicCredentials;
import software.amazon.awssdk.auth.credentials.StaticCredentialsProvider;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.dynamodb.DynamoDbAsyncClient;
import software.amazon.awssdk.services.dynamodb.model.*;

import java.io.IOException;
import java.net.ServerSocket;
import java.net.URI;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class DynamoTestUtils {

  @Value("${aws.dynamodb.region}")
  public String dynamoRegion;

  private static DynamoDBProxyServer dynamoProxy;

  private static int port;

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

  public static void setupDynamo() {
    port = getAvailablePort();

    try {
      dynamoProxy = ServerRunner.createServerFromCommandLineArgs(new String[]{
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

  public static void teardownDynamo() {
    try {
      dynamoProxy.stop();
    } catch (Exception e) {
      throw new RuntimeException();
    }
  }

  public static void setupDynamoTables() throws Exception {

    try (DynamoDbAsyncClient client = DynamoDbAsyncClient.builder()
      .region(Region.US_EAST_1)
      .endpointOverride(URI.create("http://localhost:" + port))
      .credentialsProvider(StaticCredentialsProvider.create(
        AwsBasicCredentials.create("QXdzQWNjZXNzS2V5Cg==", "QXdzU2VjcmV0S2V5Cg==")))
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
