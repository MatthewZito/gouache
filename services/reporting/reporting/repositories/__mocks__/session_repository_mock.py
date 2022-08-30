"""Repository mocks for testing.
"""

from datetime import datetime


class MockSessionRepository:
    """A SessionRepository mock that returns an infinite session."""

    def get(self, key: str):  # pylint: disable=W0613
        """Mock get - returns an infinite session.

        Args:
            key (str): Key = not used.

        Returns:
            dict: Infinite session. @todo type
        """
        return '{"username":"user","expiry":"9999-08-26T15:28:03.683Z"}'


class MockSessionRepositoryUnauthorized:
    """A SessionRepository mock that returns a null session."""

    def get(self, key: str):  # pylint: disable=W0613
        """Mock get - returns None.

        Args:
            key (str): Key = not used.

        Returns:
            None: No session.
        """
        return None
