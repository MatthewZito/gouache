"""Repository mocks for testing.
"""


class MockReportRepository:
    """A ReportRepository mock that no-ops."""

    def __init__(self, table_name: str) -> None:
        """Mock ReportRepository - no-op"""

    def get_all(self, last_page_key: str | None):
        """Mock get_all - no-op."""

    def get(self, key: str):
        """Mock get - no-op."""

    def put(self, **kwargs):
        """Mock put - no-op."""


class MockReportRepositoryConnectionRefused:
    """A ReportRepository mock that raises a connection exception."""

    def __init__(self, table_name: str) -> None:
        raise ConnectionRefusedError('test')
