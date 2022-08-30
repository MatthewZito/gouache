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

    def get_all(self, last_page_key: str | None):
        # {"data":{"Count":3,"Items":[{"Caller":"x1","Data":"y1","Id":"73e0f8ad-4b01-44a2-a450-d62129c85675","Name":"z1","TS":"1661839391418.198"},{"Caller":"x2","Data":"y2","Id":"3d84b7fc-fcf7-4cc0-a79e-7e1602893978","Name":"z2","TS":"1661839402457.305"},{"Caller":"x","Data":"y","Id":"551f4ec2-3438-477e-a3db-d136a1269c3e","Name":"z","TS":"1661839377285.796"}],"ResponseMetadata":{"HTTPHeaders":{"content-length":"452","content-type":"application/x-amz-json-1.0","date":"Tue, 30 Aug 2022 06:04:10 GMT","server":"Jetty(9.4.43.v20210629)","x-amz-crc32":"1826559263","x-amzn-requestid":"d805d21a-0602-413f-bbe8-3e36348ec3aa"},"HTTPStatusCode":200,"RequestId":"d805d21a-0602-413f-bbe8-3e36348ec3aa","RetryAttempts":0},"ScannedCount":3},"flags":0,"friendly":"","internal":""}

        if last_page_key is None:
            return self.table.scan()
        return self.table.scan(ExclusiveStartKey=last_page_key)

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
