package com.github.exbotanical.resource.config;

import com.amazonaws.auth.AWSStaticCredentialsProvider;
import com.amazonaws.auth.BasicAWSCredentials;
import com.amazonaws.client.builder.AwsClientBuilder;
import com.amazonaws.services.sqs.AmazonSQSAsync;
import com.amazonaws.services.sqs.AmazonSQSAsyncClientBuilder;
import com.github.exbotanical.resource.utils.FormatterUtils;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.aws.messaging.config.QueueMessageHandlerFactory;
import org.springframework.cloud.aws.messaging.listener.QueueMessageHandler;
import org.springframework.cloud.aws.messaging.listener.SimpleMessageListenerContainer;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.messaging.converter.MappingJackson2MessageConverter;
import org.springframework.scheduling.concurrent.ThreadPoolTaskExecutor;

import java.util.Collections;

@Configuration
public class SQSConfig {

  @Value("${aws.access_key")
  private String accessKey;

  @Value("${aws.secret_key")
  private String secretKey;

  @Value("${aws.sqs.host}")
  private String host;

  @Value("${aws.sqs.port}")
  private String port;

  @Value("${aws.sqs.region}")
  private String region;

  private static final int N_MESSAGES = 10;

  private static final int THREAD_POOL_SIZE = 10;

  @Bean
  public AmazonSQSAsync amazonSQS() {
    return AmazonSQSAsyncClientBuilder
        .standard()
        .withEndpointConfiguration(
            new AwsClientBuilder.EndpointConfiguration(
                FormatterUtils.toEndpoint(host, port),
                region))
        .withCredentials(
            new AWSStaticCredentialsProvider(
                new BasicAWSCredentials(
                    accessKey,
                    secretKey)))
        .build();
  }

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

  // pull messages from queues
  @Bean
  public SimpleMessageListenerContainer simpleMessageListenerContainer(
      QueueMessageHandler queueMessageHandler) {
    SimpleMessageListenerContainer simpleMessageListenerContainer =
        new SimpleMessageListenerContainer();

    simpleMessageListenerContainer.setAmazonSqs(amazonSQS());
    simpleMessageListenerContainer.setMessageHandler(queueMessageHandler);
    simpleMessageListenerContainer.setMaxNumberOfMessages(N_MESSAGES);
    simpleMessageListenerContainer.setTaskExecutor(threadPoolTaskExecutor());

    return simpleMessageListenerContainer;
  }

  public ThreadPoolTaskExecutor threadPoolTaskExecutor() {
    ThreadPoolTaskExecutor executor = new ThreadPoolTaskExecutor();

    executor.setCorePoolSize(THREAD_POOL_SIZE);
    executor.setMaxPoolSize(THREAD_POOL_SIZE * 2);
    executor.initialize();

    return executor;
  }
}
