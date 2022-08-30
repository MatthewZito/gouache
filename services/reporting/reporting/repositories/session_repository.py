"""This module exposes APIs for interacting with session data sources.
"""
import os
from redis import Redis


class SessionRepository:
    """A repository for session data."""

    def __init__(self) -> None:
        host = os.getenv('REDIS_HOST', 'localhost')
        port = os.getenv('REDIS_PORT', '6379')
        password = os.getenv('REDIS_PASSWORD', 'password')

        self.client = Redis(host=host, port=int(port), db=0, password=password)

    def get(self, key: str):
        """Get a session by its session id.

        Args:
            key (str): The session id.

        Returns:
            Session: @todo
        """
        return self.client.get(key)
