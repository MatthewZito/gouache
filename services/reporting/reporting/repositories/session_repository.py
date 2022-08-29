import os
from redis import Redis

from .base_repository import BaseRepository


class SessionRepository(BaseRepository):
    def __init__(self) -> None:
        host = os.getenv('REDIS_HOST', 'localhost')
        port = os.getenv('REDIS_PORT', '6379')
        password = os.getenv('REDIS_PASSWORD', 'password')

        # self.client = Redis(host=host, port=port, db=0, password=password)
        self.client = Redis(host=host, port=port, db=0)

    def get(self, key: str):
        return self.client.get(key)

    def put(self, **kwargs):
        pass
