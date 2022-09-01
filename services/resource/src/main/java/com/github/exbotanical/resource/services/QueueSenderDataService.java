package com.github.exbotanical.resource.services;

import com.amazonaws.services.sqs.AmazonSQSAsync;
import com.amazonaws.services.sqs.model.MessageAttributeValue;
import com.amazonaws.services.sqs.model.SendMessageRequest;
import com.amazonaws.services.sqs.model.SendMessageResult;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.exbotanical.resource.models.reporting.GouacheReport;
import com.github.exbotanical.resource.models.reporting.GouacheReportName;
import com.github.exbotanical.resource.utils.FormatterUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.Map;

/**
 * An implementation of QueueSenderService.
 */
@Service
public class QueueSenderDataService implements QueueSenderService {
  @Autowired
  AmazonSQSAsync amazonSQSAsync;

  @Autowired
  ObjectMapper objectMapper;

  @Value("${aws.sqs.host}")
  private String host;

  @Value("${aws.sqs.port}")
  private String port;

  @Value("${aws.sqs.report_queue}")
  private String reportQueueId;

  private String reportQueueEndpoint;

  private String getReportQueueEndpoint() {
    if (reportQueueEndpoint == null) {
      reportQueueEndpoint = FormatterUtils.toEndpoint(host, port) + "/queue/" + reportQueueId;
    }

    return reportQueueEndpoint;
  }


  @Override
  public SendMessageResult sendMessage(Object message, GouacheReportName name) {
    return sendNormalizedMessage(message, name);

  }

  private SendMessageResult sendNormalizedMessage(Object reportData, GouacheReportName name) {
    try {

      GouacheReport report = GouacheReport.builder()
          .name(name)
          .data(reportData)
          .build();


      String serializedData = objectMapper.writeValueAsString(report.getData());

      final SendMessageRequest sendMessageRequest = new SendMessageRequest();

      sendMessageRequest.withMessageBody(serializedData);
      sendMessageRequest.withQueueUrl(getReportQueueEndpoint());
      sendMessageRequest.withMessageAttributes(deriveMessageAttributes(report));

      return amazonSQSAsync.sendMessage(sendMessageRequest);
    } catch (JsonProcessingException ex) {
      // @todo
      System.out.println(ex);

      return null;
    }
  }

  private Map<String, MessageAttributeValue> deriveMessageAttributes(GouacheReport report) {
    final Map<String, MessageAttributeValue> messageAttributes = new HashMap<>();

    messageAttributes.put("name", new MessageAttributeValue()
        .withDataType("String")
        .withStringValue(report.getName().toString()));

    messageAttributes.put("caller", new MessageAttributeValue()
        .withDataType("String")
        .withStringValue(report.getCaller()));

    return messageAttributes;
  }

}
