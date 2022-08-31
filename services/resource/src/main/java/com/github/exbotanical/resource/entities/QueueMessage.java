package com.github.exbotanical.resource.entities;

import lombok.Builder;
import lombok.Data;

/**
 * Required data for a queue message as utilized by the QueueSenderService.
 */
@Builder
@Data
public class QueueMessage {
  String sender;

  String body;
}
