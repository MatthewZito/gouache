import os

import boto3
from botocore.exceptions import ClientError, ParamValidationError
from botocore.config import Config


class MessageQueueService:
    def __init__(self, table_name: str) -> None:
        host = os.getenv('SQS_HOST', 'http://localhost')
        port = os.getenv('SQS_PORT', '8000')
        region = os.getenv('SQS_REGION', 'us-east-2')

        config = Config(
            region_name=region,
            signature_version='v4',
            retries={'max_attempts': 3, 'mode': 'standard'},
        )

        self.client = boto3.resource(
            'sqs', endpoint_url=f"{host}:{port}", config=config
        )
        # we can store this as a field given it is lazy-loaded
        self.table = self.client.Table(table_name)
