package com.github.exbotanical.session.config;

import com.amazonaws.auth.AWSStaticCredentialsProvider;
import com.amazonaws.auth.BasicAWSCredentials;
import com.amazonaws.client.builder.AwsClientBuilder;
import com.amazonaws.services.dynamodbv2.AmazonDynamoDB;
import com.amazonaws.services.dynamodbv2.AmazonDynamoDBClientBuilder;
import com.amazonaws.services.dynamodbv2.datamodeling.DynamoDBMapper;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class DynamoDBConfig {

  @Bean
  public DynamoDBMapper dynamoDBMapper() {
    return new DynamoDBMapper(buildDb());
  }

  @Value("${aws.accessKey")
  private String accessKey;

  @Value("${aws.secretKey")
  private String secretKey;

  @Value("${aws.dynamodb.host}")
  private String host;

  @Value("${aws.dynamodb.port}")
  private String port;

  @Value("${aws.dynamodb.region}")
  private String region;

  private AmazonDynamoDB buildDb() {
    return AmazonDynamoDBClientBuilder
        .standard()
        .withEndpointConfiguration(
            new AwsClientBuilder.EndpointConfiguration(
                String.format("%s:%s", host, port),
                region))
        .withCredentials(
            new AWSStaticCredentialsProvider(
                new BasicAWSCredentials(
                    accessKey,
                    secretKey)))
        .build();
  }
}