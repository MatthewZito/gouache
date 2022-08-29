from analytics.repositories.base_repository import BaseRepository


class MockSessionRepository(BaseRepository):
    def get(self, key: str):
        return f'{{"Username":"user","Expiry":"{datetime.datetime.max}"}}'
