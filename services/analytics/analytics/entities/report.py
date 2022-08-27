from datetime import datetime, timezone

id = 0


class Report:
    def __init__(self, name: str, caller: str, data: str) -> None:
        global id
        id += 1

        self.name = name
        self.caller = caller
        self.data = data
        self.ts = datetime.now(timezone.utc).timestamp() * 1000
        self.id = str(id)

    def __str__(self) -> str:
        data = {
            'id': self.id,
            'name': self.name,
            'caller': self.caller,
            'data': self.data,
            'ts': self.ts,
        }

        return str(data)
