import os
import threading
from typing import Dict, List

import boto3
from botocore.config import Config
from boto3_type_annotations.sqs import Client
from flask import Flask
from werkzeug.local import LocalProxy

from reporting.context.context import get_report_ctx


class MessageQueueService:
    """An SQS message listener service.
    Polls SQS for new report messages, processes them by generating a new persistent report,
    then deletes them."""

    def __init__(self, queue_name: str, app: Flask) -> None:
        host = os.getenv('SQS_HOST', 'http://localhost')
        port = os.getenv('SQS_PORT', '9324')
        region = os.getenv('SQS_REGION', 'us-east-1')

        config = Config(
            region_name=region,
            signature_version='v4',
            retries={'max_attempts': 3, 'mode': 'standard'},
        )

        self.app = app
        self.endpoint = f"{host}:{port}/queue/{queue_name}"
        self.client: Client = boto3.client(
            'sqs',
            endpoint_url=self.endpoint,
            config=config,
        )

    def init(self) -> None:
        bg = threading.Thread(name='bg', target=self.listen)
        bg.start()

    def process_messages(self, messages: List[Dict]) -> None:
        for message in messages:

            body = message['Body']
            if body == None:
                continue

            with self.app.app_context():
                db = LocalProxy(get_report_ctx)

                # @todo parse message, validate success
                db.put(  # type: ignore
                    caller="gouache/queue",
                    data=body,
                    report_id="auto_gen",
                    name="report_name",
                )

            self.delete_message(message['ReceiptHandle'])

    def delete_message(self, receipt_handle: str) -> None:
        self.client.delete_message(QueueUrl=self.endpoint, ReceiptHandle=receipt_handle)

    def recv_messages(self):
        try:
            response = self.client.receive_message(
                QueueUrl=self.endpoint,
                AttributeNames=['SentTimestamp'],
                MaxNumberOfMessages=10,
                MessageAttributeNames=['All'],
                VisibilityTimeout=0,
                WaitTimeSeconds=20,
            )

            self.process_messages(response['Messages'])

        except Exception as ex:
            print(ex)

    def listen(self) -> None:
        while True:
            self.recv_messages()
