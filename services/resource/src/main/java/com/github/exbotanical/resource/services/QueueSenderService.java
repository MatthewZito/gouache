package com.github.exbotanical.resource.services;

import com.amazonaws.services.sqs.model.SendMessageResult;
import com.github.exbotanical.resource.models.ReportName;

/**
 * A message sender service for SQS.
 */
public interface QueueSenderService {
  SendMessageResult sendMessage(String message, ReportName name);

  SendMessageResult sendMessage(Object message, ReportName name);
}
