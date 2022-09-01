package com.github.exbotanical.resource.services;

import com.amazonaws.services.sqs.model.SendMessageResult;
import com.github.exbotanical.resource.models.reporting.GouacheReportName;

/**
 * A message sender service for SQS.
 */
public interface QueueSenderService {
  SendMessageResult sendMessage(Object message, GouacheReportName name);
}
