package com.github.exbotanical.resource.entities;

import lombok.Builder;
import lombok.Data;

@Builder
@Data
public class QueueMessage {
  String sender;

  String body;
}
