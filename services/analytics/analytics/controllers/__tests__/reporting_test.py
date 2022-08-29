import json
from types import SimpleNamespace
import unittest
from unittest import mock

from analytics.entities.report import Report, ReportMatcher

from analytics.repositories.__mocks__.session_repository_mock import (
    MockSessionRepository,
    MockSessionRepositoryUnauthorized,
)
from analytics.main import app
from analytics.repositories.report_repository import ReportRepository


class TestReportingController(unittest.TestCase):
    def setUp(self):
        self.app = app

    def tearDown(self):
        pass

    @mock.patch(
        'analytics.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'analytics.context.context.ReportRepository',
    )
    def test_get_report_ok(self, m: mock.Mock):
        test_key = 'c22c1173-93be-4550-9200-afe7df28bf2f'
        expected = Report(caller='t', data='t', name='t')

        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()
        m.return_value.get = mock.MagicMock(
            return_value={
                'Item': {
                    'Data': expected.data,
                    'Id': test_key,
                    'Caller': expected.caller,
                    'Name': expected.name,
                    'TS': '1661786507886.366',
                },
                'ResponseMetadata': {
                    'RequestId': 'f850f0b0-6b1f-460f-b48d-acd82054ef53',
                    'HTTPStatusCode': 200,
                    'HTTPHeaders': {
                        'date': 'Mon, 29 Aug 2022 15:24:18 GMT',
                        'content-type': 'application/x-amz-json-1.0',
                        'x-amz-crc32': '198904004',
                        'x-amzn-requestid': 'f850f0b0-6b1f-460f-b48d-acd82054ef53',
                        'content-length': '177',
                        'server': 'Jetty(9.4.43.v20210629)',
                    },
                    'RetryAttempts': 0,
                },
            }
        )

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get(f'/api/report/{test_key}')

                self.assertEqual(res.status_code, 200)

                res_payload_data = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                ).data

                actual = Report(
                    # @todo normalize
                    name=res_payload_data.Name,
                    caller=res_payload_data.Caller,
                    data=res_payload_data.Data,
                )

                self.assertEqual(ReportMatcher(expected=expected), actual)

                m.return_value.get.assert_called_once_with(test_key)

    @mock.patch(
        'analytics.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'analytics.context.context.ReportRepository',
    )
    def test_get_report_not_found(self, m: mock.Mock):
        test_key = 'c22c1173-93be-4550-9200-afe7df28bf2f'

        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()
        m.return_value.get = mock.MagicMock(
            return_value={
                'ResponseMetadata': {
                    'RequestId': '682a0c3b-0b08-48a6-85d4-7f8486123758',
                    'HTTPStatusCode': 200,
                    'HTTPHeaders': {
                        'date': 'Mon, 29 Aug 2022 15:36:04 GMT',
                        'content-type': 'application/x-amz-json-1.0',
                        'x-amz-crc32': '2745614147',
                        'x-amzn-requestid': '682a0c3b-0b08-48a6-85d4-7f8486123758',
                        'content-length': '2',
                        'server': 'Jetty(9.4.43.v20210629)',
                    },
                    'RetryAttempts': 0,
                }
            }
        )

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get(f'/api/report/{test_key}')

                self.assertEqual(res.status_code, 404)

                res_payload_data = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                ).data

                self.assertIsNone(res_payload_data)

                m.return_value.get.assert_called_once_with(test_key)

    @mock.patch(
        'analytics.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'analytics.context.context.ReportRepository',
    )
    def test_get_report_error(self, m: mock.Mock):
        test_key = 'c22c1173-93be-4550-9200-afe7df28bf2f'

        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()
        m.return_value.get = mock.MagicMock(
            return_value=str(Exception('test get error'))
        )

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get(f'/api/report/{test_key}')

                self.assertEqual(res.status_code, 400)

                res_payload = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                )

                self.assertIsNone(res_payload.data)
                self.assertEqual(
                    'test get error',
                    res_payload.internal,
                )
                self.assertEqual(
                    # @todo as const
                    'An exception occurred while retrieving the report',
                    res_payload.friendly,
                )

                m.return_value.get.assert_called_once_with(test_key)

    @mock.patch(
        'analytics.context.context.SessionRepository',
        new=MockSessionRepositoryUnauthorized,
    )
    @mock.patch(
        'analytics.context.context.ReportRepository',
    )
    def test_get_report_unauthorized(self, m: mock.Mock):
        test_key = 'c22c1173-93be-4550-9200-afe7df28bf2f'

        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()
        m.return_value.get = mock.MagicMock()

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get(f'/api/report/{test_key}')

                self.assertEqual(res.status_code, 401)

                res_payload = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                )

                self.assertIsNone(res_payload.data)
                self.assertEqual(
                    'You must be authorized to access this resource',
                    res_payload.friendly,
                )

                m.return_value.get.assert_not_called()

    @mock.patch(
        'analytics.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'analytics.context.context.ReportRepository',
    )
    def test_create_report_ok(self, m: mock.Mock):
        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()
        m.return_value.put = mock.MagicMock(
            return_value={'ResponseMetadata': {'HTTPStatusCode': 200}}
        )

        raw_report = {
            'name': 'test report',
            'caller': 'gouache_test',
            'data': 'data',
        }

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.post('/api/report', json=raw_report)

                self.assertEqual(res.status_code, 201)

                res_payload = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                )

                self.assertIsInstance(res_payload.data.id, str)
                self.assertEqual(res_payload.friendly, '')
                self.assertEqual(res_payload.internal, '')

                m.return_value.put.assert_called_once_with(
                    caller=raw_report['caller'],
                    data=raw_report['data'],
                    id=mock.ANY,
                    name=raw_report['name'],
                )

    @mock.patch(
        'analytics.context.context.SessionRepository',
        new=MockSessionRepositoryUnauthorized,
    )
    @mock.patch(
        'analytics.context.context.ReportRepository',
    )
    def test_create_report_unauthorized(self, m: mock.Mock):
        raw_report = {
            'name': 'test report',
            'caller': 'gouache_test',
            'data': 'data',
        }

        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()
        m.return_value.put = mock.MagicMock()

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.post('/api/report', json=raw_report)

                self.assertEqual(res.status_code, 401)

                res_payload = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                )

                self.assertIsNone(res_payload.data)
                self.assertEqual(
                    'You must be authorized to access this resource',
                    res_payload.friendly,
                )

                m.return_value.put.assert_not_called()

    @mock.patch(
        'analytics.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch('analytics.context.context.ReportRepository')
    def test_create_report_error(self, m: mock.Mock):
        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()
        m.return_value.put = mock.MagicMock(
            return_value=str(Exception('test put error'))
        )

        raw_report = {
            'name': 'test report',
            'caller': 'gouache_test',
            'data': 'data',
        }

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.post('/api/report', json=raw_report)

                self.assertEqual(res.status_code, 400)

                res_payload = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                )

                self.assertIsNone(res_payload.data)
                self.assertEqual(
                    'test put error',
                    res_payload.internal,
                )
                self.assertEqual(
                    # @todo as const
                    'An exception occurred while creating the report',
                    res_payload.friendly,
                )

                m.return_value.put.assert_called_once_with(
                    caller=raw_report['caller'],
                    data=raw_report['data'],
                    id=mock.ANY,
                    name=raw_report['name'],
                )


if __name__ == '__main__':
    unittest.main()
