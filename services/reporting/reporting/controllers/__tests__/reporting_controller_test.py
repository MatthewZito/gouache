import json
from types import SimpleNamespace
import unittest
from unittest import mock

from reporting.entities.report import Report, ReportMatcher
from reporting.meta.const import (
    E_REPORT_CREATE,
    E_REPORT_CREATE_INVALID_INPUT,
    E_REPORT_GET,
    E_REPORT_GET_ALL,
    E_UNAUTHORIZED,
)

from reporting.repositories.__mocks__.session_repository_mock import (
    MockSessionRepository,
    MockSessionRepositoryUnauthorized,
)
from reporting.main import create_app
from reporting.repositories.report_repository import ReportRepository


class TestReportingController(unittest.TestCase):
    def setUp(self):
        self.app = create_app()
        self.test_key = 'c22c1173-93be-4550-9200-afe7df28bf2f'

    def tearDown(self):
        pass

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'reporting.context.context.ReportRepository',
    )
    def test_get_all_reports_ok(self, m: mock.Mock):
        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.get_all = mock.MagicMock(  # type: ignore
            return_value={
                "Count": 3,
                "Items": [
                    {
                        "caller": "x1",
                        "data": "y1",
                        "id": "73e0f8ad-4b01-44a2-a450-d62129c85675",
                        "name": "z1",
                        "ts": "1661839391418.198",
                    },
                    {
                        "caller": "x2",
                        "data": "y2",
                        "id": "3d84b7fc-fcf7-4cc0-a79e-7e1602893978",
                        "name": "z2",
                        "ts": "1661839402457.305",
                    },
                    {
                        "caller": "x",
                        "data": "y",
                        "id": "551f4ec2-3438-477e-a3db-d136a1269c3e",
                        "name": "z",
                        "ts": "1661839377285.796",
                    },
                ],
                "ResponseMetadata": {
                    "HTTPStatusCode": 200,
                },
                "ScannedCount": 3,
            }
        )

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get('/api/report')

                self.assertEqual(res.status_code, 200)

                res_payload_data = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                ).data

                self.assertEqual(len(res_payload_data.items), 3)
                self.assertEqual(res_payload_data.next, '')

                m.return_value.get_all.assert_called_once_with(None)

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'reporting.context.context.ReportRepository',
    )
    def test_get_all_reports_paginated_ok(self, m: mock.Mock):
        test_last_page_key = 'test'
        test_next_page_key = 'test2'

        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.get_all = mock.MagicMock(  # type: ignore
            return_value={
                "LastEvaluatedKey": {"id": test_next_page_key},
                "Count": 3,
                "Items": [
                    {
                        "caller": "x",
                        "data": "y",
                        "id": "551f4ec2-3438-477e-a3db-d136a1269c3e",
                        "name": "z",
                        "ts": "1661839377285.796",
                    },
                ],
                "ResponseMetadata": {
                    "HTTPStatusCode": 200,
                },
                "ScannedCount": 3,
            }
        )

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get(f'/api/report?last_page_key={test_last_page_key}')

                self.assertEqual(res.status_code, 200)

                res_payload_data = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                ).data

                self.assertEqual(len(res_payload_data.items), 1)
                self.assertEqual(res_payload_data.next, test_next_page_key)

                m.return_value.get_all.assert_called_once_with(test_last_page_key)

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'reporting.context.context.ReportRepository',
    )
    def test_get_all_reports_empty(self, m: mock.Mock):
        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.get_all = mock.MagicMock(  # type: ignore
            return_value={
                "Count": 0,
                "Items": [],
                "ResponseMetadata": {
                    "HTTPStatusCode": 200,
                },
                "ScannedCount": 3,
            }
        )

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get('/api/report')

                self.assertEqual(res.status_code, 200)

                res_payload_data = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                ).data

                self.assertEqual(res_payload_data.items, [])
                self.assertEqual(res_payload_data.next, '')

                m.return_value.get_all.assert_called_once_with(None)

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'reporting.context.context.ReportRepository',
    )
    def test_get_all_reports_bad_code(self, m: mock.Mock):
        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.get_all = mock.MagicMock(return_value={})  # type: ignore

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get('/api/report')

                self.assertEqual(res.status_code, 400)

                res_payload = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                )

                self.assertIsNone(res_payload.data)
                self.assertEqual(res_payload.friendly, E_REPORT_GET_ALL)
                self.assertEqual(res_payload.internal, SimpleNamespace())

                m.return_value.get_all.assert_called_once()

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'reporting.context.context.ReportRepository',
    )
    def test_get_all_reports_error(self, m: mock.Mock):
        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.get_all = mock.MagicMock(  # type: ignore
            return_value=str(Exception('test'))
        )

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get('/api/report')

                self.assertEqual(res.status_code, 400)

                res_payload = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                )

                self.assertIsNone(res_payload.data)
                self.assertEqual(res_payload.friendly, E_REPORT_GET_ALL)
                self.assertEqual(res_payload.internal, 'test')

                m.return_value.get_all.assert_called_once()

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepositoryUnauthorized,
    )
    @mock.patch(
        'reporting.context.context.ReportRepository',
    )
    def test_get_all_reports_unauthorized(self, m: mock.Mock):

        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.get_all = mock.MagicMock()  # type: ignore

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get('/api/report')

                self.assertEqual(res.status_code, 401)

                res_payload = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                )

                self.assertIsNone(res_payload.data)
                self.assertEqual(
                    E_UNAUTHORIZED,
                    res_payload.friendly,
                )

                m.return_value.get_all.assert_not_called()

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'reporting.context.context.ReportRepository',
    )
    def test_get_report_ok(self, m: mock.Mock):
        expected = Report(caller='t', data='t', name='t')

        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.get = mock.MagicMock(  # type: ignore
            return_value={
                'Item': {
                    'data': expected.data,
                    'id': self.test_key,
                    'caller': expected.caller,
                    'name': expected.name,
                    'ts': '1661786507886.366',
                },
                'ResponseMetadata': {
                    'HTTPStatusCode': 200,
                },
            }
        )

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get(f'/api/report/{self.test_key}')

                self.assertEqual(res.status_code, 200)

                res_payload_data = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                ).data

                actual = Report(
                    name=res_payload_data.name,
                    caller=res_payload_data.caller,
                    data=res_payload_data.data,
                )

                self.assertEqual(ReportMatcher(expected=expected), actual)

                m.return_value.get.assert_called_once_with(self.test_key)

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'reporting.context.context.ReportRepository',
    )
    def test_get_report_not_found(self, m: mock.Mock):

        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.get = mock.MagicMock(  # type: ignore
            return_value={
                'ResponseMetadata': {
                    'HTTPStatusCode': 200,
                }
            }
        )

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get(f'/api/report/{self.test_key}')

                self.assertEqual(res.status_code, 404)

                res_payload_data = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                ).data

                self.assertIsNone(res_payload_data)

                m.return_value.get.assert_called_once_with(self.test_key)

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'reporting.context.context.ReportRepository',
    )
    def test_get_report_error(self, m: mock.Mock):

        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.get = mock.MagicMock(  # type: ignore
            return_value=str(Exception('test get error'))
        )

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get(f'/api/report/{self.test_key}')

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
                    E_REPORT_GET,
                    res_payload.friendly,
                )

                m.return_value.get.assert_called_once_with(self.test_key)

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepositoryUnauthorized,
    )
    @mock.patch(
        'reporting.context.context.ReportRepository',
    )
    def test_get_report_unauthorized(self, m: mock.Mock):

        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.get = mock.MagicMock()  # type: ignore

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get(f'/api/report/{self.test_key}')

                self.assertEqual(res.status_code, 401)

                res_payload = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                )

                self.assertIsNone(res_payload.data)
                self.assertEqual(
                    E_UNAUTHORIZED,
                    res_payload.friendly,
                )

                m.return_value.get.assert_not_called()

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'reporting.context.context.ReportRepository',
    )
    def test_create_report_ok(self, m: mock.Mock):
        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.put = mock.MagicMock(  # type: ignore
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
                    report_id=mock.ANY,
                    name=raw_report['name'],
                )

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'reporting.context.context.ReportRepository',
    )
    def test_create_report_invalid_input(self, m: mock.Mock):
        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.put = mock.MagicMock(  # type: ignore
            return_value={'ResponseMetadata': {'HTTPStatusCode': 200}}
        )

        raw_report = {
            'namex': 'test report',
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
                self.assertEqual(res_payload.friendly, E_REPORT_CREATE_INVALID_INPUT)
                self.assertIsInstance(res_payload.internal, str)

                m.return_value.put.assert_not_called()

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepositoryUnauthorized,
    )
    @mock.patch(
        'reporting.context.context.ReportRepository',
    )
    def test_create_report_unauthorized(self, m: mock.Mock):
        raw_report = {
            'name': 'test report',
            'caller': 'gouache_test',
            'data': 'data',
        }

        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.put = mock.MagicMock()  # type: ignore

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
                    E_UNAUTHORIZED,
                    res_payload.friendly,
                )

                m.return_value.put.assert_not_called()

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch('reporting.context.context.ReportRepository')
    def test_create_report_error(self, m: mock.Mock):
        m.return_value = ReportRepository('test')
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.put = mock.MagicMock(  # type: ignore
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
                    E_REPORT_CREATE,
                    res_payload.friendly,
                )

                m.return_value.put.assert_called_once_with(
                    caller=raw_report['caller'],
                    data=raw_report['data'],
                    report_id=mock.ANY,
                    name=raw_report['name'],
                )


if __name__ == '__main__':
    unittest.main()
