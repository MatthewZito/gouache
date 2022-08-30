"""Reporting data repositories.
"""
import os
from datetime import datetime, timezone

import boto3
from botocore.exceptions import ClientError, ParamValidationError
from botocore.config import Config


class ReportRepository:
    """A data repository for reporting data.
    Uses DynamoDB as a data source.

    Args:
        table_name (str): The dynamodb table name.
    """

    def __init__(self, table_name: str) -> None:
        host = os.getenv('DYNAMO_HOST', 'http://localhost')
        port = os.getenv('DYNAMO_PORT', '8000')
        region = os.getenv('DYNAMO_REGION', 'us-east-2')

        config = Config(
            region_name=region,
            signature_version='v4',
            retries={'max_attempts': 3, 'mode': 'standard'},
        )

        self.client = boto3.resource(
            'dynamodb', endpoint_url=f"{host}:{port}", config=config
        )
        # we can store this as a field given it is lazy-loaded
        self.table = self.client.Table(table_name)

    def get(self, key: str):
        try:
            response = self.table.get_item(Key={'Id': key})
            return response

        except (ClientError, ParamValidationError, Exception) as ex:
            return str(ex)

    def put(self, name: str, caller: str, data: str, report_id: str):
        try:
            response = self.table.put_item(
                Item={
                    'Name': name,
                    'Caller': caller,
                    'Data': data,
                    'TS': str(datetime.now(timezone.utc).timestamp() * 1000),
                    'Id': report_id,
                }
            )

            return response

        except (ClientError, ParamValidationError, Exception) as ex:
            return str(ex)
