from datetime import datetime, timezone
import os

from botocore.exceptions import ClientError, ParamValidationError
import boto3
from botocore.config import Config

from .base_repository import BaseRepository


class ReportRepository(BaseRepository):
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

    def put(self, name: str, caller: str, data: str, id: str):
        try:
            response = self.table.put_item(
                Item={
                    'Name': name,
                    'Caller': caller,
                    'Data': data,
                    'TS': str(datetime.now(timezone.utc).timestamp() * 1000),
                    'Id': id,
                }
            )

            return response

        except ClientError as error:
            return str(error)
        except ParamValidationError as error:
            return str(error)

    def get(self, id: str):

        try:
            response = self.table.get_item(Key={'Id': id})
            return response

        except ClientError as error:
            return str(error)
        except ParamValidationError as error:
            return str(error)
