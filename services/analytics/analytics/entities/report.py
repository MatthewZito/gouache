from datetime import datetime, timezone
import uuid


class Report:
    def __init__(self, name: str, caller: str, data: str) -> None:
        self.name = name
        self.caller = caller
        self.data = data
        self.ts = datetime.now(timezone.utc).timestamp() * 1000
        self.id = uuid.uuid4()

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
    expected: Report

    def __init__(self, expected):
        self.expected = expected

    def __eq__(self, other):
        return (
            type(other.id) is uuid.UUID
            and type(other.ts) is float
            and self.expected.name == other.name
            and self.expected.caller == other.caller
        )
