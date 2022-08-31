package com.github.exbotanical.resource.services;

import com.amazonaws.services.sqs.AmazonSQSAsync;
import com.amazonaws.services.sqs.model.SendMessageResult;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.exbotanical.resource.models.GouacheReport;
import com.github.exbotanical.resource.models.ReportName;
import com.github.exbotanical.resource.utils.FormatterUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

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
  public SendMessageResult sendMessage(String message, ReportName name) {
    return sendNormalizedMessage(message, name);
  }

  @Override
  public SendMessageResult sendMessage(Object message, ReportName name) {
    try {
      return sendNormalizedMessage(objectMapper.writeValueAsString(message), name);
    } catch (JsonProcessingException ex) {
      return null;
    }
  }

  private SendMessageResult sendNormalizedMessage(String reportData, ReportName name) {
    try {
      GouacheReport report = GouacheReport.builder()
        .data(objectMapper.writeValueAsString(reportData))
        .caller("gouache/resource")
        .name(name.toString())
        .build();

      String serializedMessage = objectMapper.writeValueAsString(report);
      System.out.printf("Writing message %s to queue %s%n", serializedMessage, getReportQueueEndpoint());

      return amazonSQSAsync.sendMessage(getReportQueueEndpoint(), serializedMessage);
    } catch (JsonProcessingException ex) {
      // @todo
      return null;
    }
  }
}
