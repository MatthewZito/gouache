import unittest

from ..deserialize import deserialize
from reporting.entities.report import Report, ReportMatcher
from flask import Flask, request


@deserialize(Report)
def automake_report(report: Report):
    return report


class TestRequestDeserialization(unittest.TestCase):
    def setUp(self):
        self.app = Flask(__name__)

    def tearDown(self):
        pass

    def test_deserialize_ok(self):

        raw_report = {
            'name': 'test report',
            'caller': 'gouache_test',
            'data': 'data',
        }

        expected = Report(
            name=raw_report['name'],
            caller=raw_report['caller'],
            data=raw_report['data'],
        )

        with self.app.test_request_context(json=raw_report):
            actual = automake_report(request.data)

            self.assertEqual(ReportMatcher(expected=expected), actual)

    def test_deserialize_empty(self):
        raw_report = {}

        with self.app.test_request_context(json=raw_report):
            actual = automake_report(request.data)

            self.assertIsNone(actual)


if __name__ == '__main__':
    unittest.main()
