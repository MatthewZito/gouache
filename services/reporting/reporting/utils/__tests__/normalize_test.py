import unittest

from ..normalize import normalize_dynamo_report


class TestNormalizationUtils(unittest.TestCase):
    def test_normalize_report_ok(self):

        raw_report = {
            'Data': 'some data',
            'Id': 'eddde4d5-90b7-4cc3-91bb-b4561b4137b4',
            'Caller': 'gouache-test',
            'Name': 'the report name',
            'TS': '1661786507886.366',
        }

        expected = {
            'id': raw_report['Id'],
            'data': raw_report['Data'],
            'caller': raw_report['Caller'],
            'name': raw_report['Name'],
            'ts': raw_report['TS'],
        }

        actual = normalize_dynamo_report(raw_report)
        self.assertEqual(expected, actual)

    def test_normalize_report_incomplete(self):

        raw_report = {
            'Data': 'some data',
            'Id': 'eddde4d5-90b7-4cc3-91bb-b4561b4137b4',
            'Caller': 'gouache-test',
            'TS': '1661786507886.366',
        }

        expected = {
            'id': raw_report['Id'],
            'data': raw_report['Data'],
            'caller': raw_report['Caller'],
            'name': None,
            'ts': raw_report['TS'],
        }

        actual = normalize_dynamo_report(raw_report)
        self.assertEqual(expected, actual)


if __name__ == '__main__':
    unittest.main()
