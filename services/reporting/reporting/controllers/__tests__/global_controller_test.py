import json
from types import SimpleNamespace
import unittest
from unittest import mock

from reporting.meta.const import (
    E_INTERNAL_SERVER_ERROR,
    E_METHOD_NOT_ALLOWED,
    E_ROUTE_NOT_FOUND,
    E_UNAUTHORIZED,
)
from reporting.main import create_app
from reporting.repositories.__mocks__.report_repository_mock import (
    MockReportRepositoryConnectionRefused,
)
from reporting.repositories.__mocks__.session_repository_mock import (
    MockSessionRepository,
)


class TestGlobalController(unittest.TestCase):
    def setUp(self):
        self.app = create_app()

    def tearDown(self):
        pass

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    def test_not_found_handler(self):
        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get('/api/mn')

                self.assertEqual(res.status_code, 404)

                res_payload = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                )

                self.assertIsNone(res_payload.data)
                self.assertEqual(res_payload.friendly, E_ROUTE_NOT_FOUND)
                self.assertIsInstance(res_payload.internal, str)

    def test_unauthorized_handler(self):
        with self.app.app_context():
            with self.app.test_client() as c:
                res = c.get('/api/report')

                self.assertEqual(res.status_code, 401)

                res_payload = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                )

                self.assertIsNone(res_payload.data)
                self.assertEqual(res_payload.friendly, E_UNAUTHORIZED)
                self.assertIsInstance(res_payload.internal, str)

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    def test_not_allowed_handler(self):
        with self.app.app_context():
            with self.app.test_client() as c:
                res = c.delete('/api/report')

                self.assertEqual(res.status_code, 405)

                res_payload = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                )

                self.assertIsNone(res_payload.data)
                self.assertEqual(res_payload.friendly, E_METHOD_NOT_ALLOWED)
                self.assertIsInstance(res_payload.internal, str)

    @mock.patch(
        'reporting.context.context.SessionRepository',
        new=MockSessionRepository,
    )
    @mock.patch(
        'reporting.context.context.ReportRepository',
        new=MockReportRepositoryConnectionRefused,
    )
    def test_server_error_handler(self):
        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get('/api/report')

                self.assertEqual(res.status_code, 500)

                res_payload = json.loads(
                    res.data,
                    object_hook=lambda d: SimpleNamespace(**d),
                )

                self.assertIsNone(res_payload.data)
                self.assertEqual(res_payload.friendly, E_INTERNAL_SERVER_ERROR)
                self.assertIsInstance(res_payload.internal, str)


if __name__ == '__main__':
    unittest.main()
