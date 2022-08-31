package com.github.exbotanical.resource.config;

import com.amazonaws.auth.AWSStaticCredentialsProvider;
import com.amazonaws.auth.AnonymousAWSCredentials;
import com.amazonaws.client.builder.AwsClientBuilder;
import com.amazonaws.services.sqs.AmazonSQSAsync;
import com.amazonaws.services.sqs.AmazonSQSAsyncClientBuilder;
import com.github.exbotanical.resource.utils.FormatterUtils;
import java.util.Collections;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.aws.messaging.config.QueueMessageHandlerFactory;
import org.springframework.cloud.aws.messaging.listener.QueueMessageHandler;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.messaging.converter.MappingJackson2MessageConverter;

/**
 * AWS SQS configurations.
 */
@Configuration
public class SQSConfig {

  @Value("${aws.sqs.host}")
  private String host;

  @Value("${aws.sqs.port}")
  private String port;

  @Value("${aws.sqs.region}")
  private String region;

  private static final int N_MESSAGES = 10;

  private static final int THREAD_POOL_SIZE = 10;

  /**
   * AWS SQS client builder.
   *
   * @return AWS SQS client.
   */
  @Bean
  public AmazonSQSAsync amazonSQS() {
    return AmazonSQSAsyncClientBuilder
        .standard()
        .withEndpointConfiguration(
            new AwsClientBuilder.EndpointConfiguration(
                FormatterUtils.toEndpoint(host, port), region))
        .withCredentials(
            new AWSStaticCredentialsProvider(new AnonymousAWSCredentials()))
        .build();
  }

  /**
   * An AWS SQS message handler and deserializer.
   *
   * @return An initialized QueueMessageHandler.
   */
  @Bean
  public QueueMessageHandler queueMessageHandler() {
    QueueMessageHandlerFactory queueMessageHandlerFactory = new QueueMessageHandlerFactory();
    queueMessageHandlerFactory.setAmazonSqs(amazonSQS());

    // Deserialize message strings into objects
    MappingJackson2MessageConverter messageConverter = new MappingJackson2MessageConverter();

    // Avoid failures given SQS messages don't include a MIME type
    messageConverter.setStrictContentTypeMatch(false);

    // @see
    // https://stackoverflow.com/questions/57613779/aws-sqslistener-deserialize-custom-object-with-jackson
    queueMessageHandlerFactory.setMessageConverters(Collections.singletonList(messageConverter));

    return queueMessageHandlerFactory.createQueueMessageHandler();
  }
}
