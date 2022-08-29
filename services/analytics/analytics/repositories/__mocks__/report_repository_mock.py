from analytics.repositories.base_repository import BaseRepository
from analytics.entities.report import Report


class MockReportRepository(BaseRepository):
    def get(self, key: str):
        return {'Item': Report(caller='t', data='t', name='t')}

    def put(self, **kwargs):
        return {'ResponseMetadata': {'HTTPStatusCode': 200}}
