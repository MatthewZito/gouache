import unittest

from ..serialize import serialize
from flask import Flask, request


class Obj:
    def __init__(self, x: str, y: str, z: int):
        self.x = x
        self.y = y
        self.z = z

    def __str__(self) -> str:
        data = {'x': self.x, 'y': self.y, 'z': self.z}

        return str(data)


@serialize()
def automake_json(obj: Obj):
    return obj


class TestResponseSerialization(unittest.TestCase):
    def setUp(self):
        self.app = Flask(__name__)

    def tearDown(self):
        pass

    def test_serialize_ok(self):
        obj = Obj(x='x', y='y', z=1)

        with self.app.test_request_context():
            actual = automake_json(obj).data.decode('utf-8')

            self.assertEqual('"{\'x\': \'x\', \'y\': \'y\', \'z\': 1}"\n', actual)

    def test_serialize_empty(self):
        obj = None
        with self.app.test_request_context():
            actual = automake_json(obj).data.decode('utf-8')
            self.assertEqual('"{}"\n', actual)


if __name__ == '__main__':
    unittest.main()
