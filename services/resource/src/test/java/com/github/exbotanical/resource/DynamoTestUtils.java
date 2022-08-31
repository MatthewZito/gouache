package com.github.exbotanical.resource;

import org.springframework.beans.factory.annotation.Value;

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

  private static int port;
}
