"""
This module exposes datatypes for the system's omnipresent "reports".
"""
from datetime import datetime, timezone
import uuid


class Report:
    """A system analytics log or report.

    Args:
        name (str): The report name, for user-facing display and querying.
        caller (str): The reporting system or module.
        data (str): The serialized report data.
    """

    def __init__(self, name: str, caller: str, data: str) -> None:
        self.name = name
        self.caller = caller
        self.data = data
        self.ts = datetime.now(timezone.utc).timestamp() * 1000
        self.id = str(uuid.uuid4())

    def __str__(self) -> str:
        data = {
            'id': self.id,
            'name': self.name,
            'caller': self.caller,
            'data': self.data,
            'ts': self.ts,
        }

        return str(data)


class ReportMatcher:
    """Implements equality comparison behavior for a given Report.

    Args:
        expected (Report): The expected Report type.
    """

    expected: Report

    def __init__(self, expected):
        self.expected = expected

    def __eq__(self, other):
        return (
            isinstance(other.id, str)
            and isinstance(other.ts, float)
            and self.expected.name == other.name
            and self.expected.caller == other.caller
        )
