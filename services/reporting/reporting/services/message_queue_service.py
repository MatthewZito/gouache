"""A message queue service and related functionality.
"""
import os
import threading
from typing import Dict, List
import uuid

import boto3
from botocore.config import Config
from boto3_type_annotations.sqs import Client  # type: ignore
from flask import Flask
from werkzeug.local import LocalProxy

from reporting.context.context import get_report_ctx


class MessageQueueService:
    """An SQS message listener service.
    Polls SQS for new report messages, processes them by generating
    a new persistent report, then deletes them."""

    def __init__(self, queue_name: str, app: Flask) -> None:
        host = os.getenv('SQS_HOST', 'http://localhost')
        port = os.getenv('SQS_PORT', '9324')
        region = os.getenv('SQS_REGION', 'us-east-1')
        accessKey = os.getenv('AWS_FAKE_ACCESS_KEY')
        secretKey = os.getenv('AWS_FAKE_SECRET_KEY')

        config = Config(
            region_name=region,
            signature_version='v4',
            retries={'max_attempts': 3, 'mode': 'standard'},
        )

        self.app = app
        self.endpoint = f"{host}:{port}/queue/{queue_name}"
        session = boto3.Session(
            aws_access_key_id=accessKey,
            aws_secret_access_key=secretKey,
        )

        self.client: Client = session.client(
            'sqs',
            endpoint_url=self.endpoint,
            config=config,
        )

    def init(self) -> None:
        """Initialize the message processing task in a background thread."""
        bg = threading.Thread(name='bg', target=self.listen, daemon=True)
        bg.start()

    def process_messages(self, messages: List[Dict]) -> None:
        """Process a batch of messages received from the SQS queue.

        Args:
            messages (List[Dict]): _description_
        """
        for message in messages:

            body = message.get('Body')
            if body is None:
                continue

            attrs = message.get('MessageAttributes')
            name = attrs.get('name')
            caller = attrs.get('caller')

            with self.app.app_context():
                db = LocalProxy(get_report_ctx)
                # @todo parse message, validate success
                db.put(  # type: ignore
                    caller=caller.get('StringValue') or 'unknown',
                    data=body,
                    report_id=str(uuid.uuid4()),
                    name=name.get('StringValue') or 'unknown',
                )

            self.delete_message(message.get('ReceiptHandle'))

    def delete_message(self, receipt_handle: str | None) -> None:
        """Delete a given message from the SQS queue.

        Args:
            receipt_handle (str): The message receipt handle,
            used to identify the message being deleted from the SQS queue.
        """

        if receipt_handle is not None:
            self.client.delete_message(
                QueueUrl=self.endpoint, ReceiptHandle=receipt_handle
            )

    def recv_messages(self):
        """Receive a batch of messages from the SQS queue."""
        try:
            response = self.client.receive_message(
                QueueUrl=self.endpoint,
                AttributeNames=['SentTimestamp'],
                MaxNumberOfMessages=10,
                MessageAttributeNames=['All'],
                VisibilityTimeout=20,
                WaitTimeSeconds=20,
            )
            print(response.get('Messages'))
            self.process_messages(response.get('Messages'))

        except Exception as ex:
            print(ex)

    def listen(self) -> None:
        """Listen for new SQS messages.
        This should be offloaded to a separate thread or process."""
        while True:
            self.recv_messages()
