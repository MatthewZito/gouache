from datetime import datetime
from reporting.repositories.base_repository import BaseRepository


class MockSessionRepository(BaseRepository):
    def get(self, key: str):

        return f'{{"Username":"user","Expiry":"{datetime.max}"}}'

    def put(self, **kwargs):
        pass


class MockSessionRepositoryUnauthorized(BaseRepository):
    def get(self, key: str):

        return None

    def put(self, **kwargs):
        pass
