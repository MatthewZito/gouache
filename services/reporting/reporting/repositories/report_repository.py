"""Reporting data repositories.
"""
import os
from datetime import datetime, timezone

import boto3
from botocore.exceptions import ClientError, ParamValidationError
from botocore.config import Config
from boto3_type_annotations.dynamodb import Client


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

        self.client: Client = boto3.resource(
            'dynamodb', endpoint_url=f"{host}:{port}", config=config
        )
        # we can store this as a field given it is lazy-loaded
        self.table = self.client.Table(table_name)

    def get_all(self, last_page_key: str | None):
        """Get all Reports.

        Args:
            last_page_key (str | None): Optional pagination key.

        Returns:
            _type_: @todo
        """
        if last_page_key is None:
            return self.table.scan()
        return self.table.scan(ExclusiveStartKey={'id': last_page_key})

    def get(self, key: str):
        """Get a Report by id.

        Args:
            key (str): Report key (id).

        Returns:
            _type_: @todo
        """
        try:
            response = self.table.get_item(Key={'id': key})
            return response

        except (ClientError, ParamValidationError, Exception) as ex:
            return str(ex)

    def put(self, name: str, caller: str, data: str, report_id: str):
        """Put a Report into the database.

        Args:
            name (str): The Report name.
            caller (str): The Report caller.
            data (str): The Report data.
            report_id (str): The Report id.

        Returns:
            _type_: @todo
        """
        try:
            response = self.table.put_item(
                Item={
                    'name': name,
                    'caller': caller,
                    'data': data,
                    'ts': str(datetime.now(timezone.utc).timestamp() * 1000),
                    'id': report_id,
                }
            )

            return response

        except (ClientError, ParamValidationError, Exception) as ex:
            return str(ex)
