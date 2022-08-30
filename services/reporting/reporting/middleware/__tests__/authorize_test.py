from datetime import datetime
import unittest
from unittest import mock

from reporting.main import create_app
from reporting.repositories.__mocks__.report_repository_mock import MockReportRepository
from reporting.repositories.session_repository import SessionRepository


class TestAuthorizationMiddleware(unittest.TestCase):
    def setUp(self):
        self.app = create_app()

    def tearDown(self):
        pass

    @mock.patch(
        'reporting.context.context.SessionRepository',
    )
    @mock.patch('reporting.context.context.ReportRepository', new=MockReportRepository)
    def test_session_expired(self, m: mock.Mock):

        m.return_value = SessionRepository()
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.get = mock.MagicMock(  # type: ignore
            return_value='{"username":"user","expiry":"1000-08-26T15:28:03.683Z"}'
        )

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get('/api/report/c22c1173-93be-4550-9200-afe7df28bf2f')

                self.assertEqual(res.status_code, 401)

    @mock.patch(
        'reporting.context.context.SessionRepository',
    )
    @mock.patch('reporting.context.context.ReportRepository', new=MockReportRepository)
    def test_session_not_extant(self, m: mock.Mock):

        m.return_value = SessionRepository()
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.get = mock.MagicMock(return_value=None)  # type: ignore

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get('/api/report/c22c1173-93be-4550-9200-afe7df28bf2f')

                self.assertEqual(res.status_code, 401)

    @mock.patch(
        'reporting.context.context.SessionRepository',
    )
    @mock.patch('reporting.context.context.ReportRepository', new=MockReportRepository)
    def test_cookie_not_extant(self, m: mock.Mock):

        m.return_value = SessionRepository()
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()  # type: ignore
        m.return_value.get = mock.MagicMock(  # type: ignore
            return_value='{"username":"user","expiry":"9999-08-26T15:28:03.683Z"}'
        )

        with self.app.app_context():
            with self.app.test_client() as c:
                res = c.get('/api/report/c22c1173-93be-4550-9200-afe7df28bf2f')

                self.assertEqual(res.status_code, 401)


if __name__ == '__main__':
    unittest.main()
