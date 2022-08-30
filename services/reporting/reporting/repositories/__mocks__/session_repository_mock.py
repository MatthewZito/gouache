"""Repository mocks for testing.
"""

from datetime import datetime


class MockSessionRepository:
    """A SessionRepository mock that returns an infinite session."""

    def get(self, key: str):

        return f'{{"Username":"user","Expiry":"{datetime.max}"}}'

    def put(self, **kwargs):
        pass


class MockSessionRepositoryUnauthorized:
    """A SessionRepository mock that returns a null session."""

    def get(self, key: str):

        return None

    def put(self, **kwargs):
        pass
