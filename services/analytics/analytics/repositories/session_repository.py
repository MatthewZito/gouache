import os
import redis


class SessionRepository:
    def __init__(self) -> None:
        host = os.getenv('REDIS_HOST', 'localhost')
        port = os.getenv('REDIS_PORT', '6379')
        password = os.getenv('REDIS_PASSWORD', 'password')

        self.client = redis.Redis(host=host, port=port, db=0, password=password)

    def get(self, key: str):
        return self.client.get(key)
